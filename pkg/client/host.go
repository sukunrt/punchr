package client

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	basichost "github.com/libp2p/go-libp2p/p2p/host/basic"
	"github.com/libp2p/go-libp2p/p2p/net/nat"
	"github.com/libp2p/go-libp2p/p2p/protocol/holepunch"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/dennis-tra/punchr/pkg/pb"
	"github.com/dennis-tra/punchr/pkg/util"
)

var (
	CommunicationTimeout = 30 * time.Second
	RetryCount           = 3
	PingDuration         = 10 * time.Second
	MaxPingCount         = 10
)

// Host holds information of the honeypot libp2p host.
type Host struct {
	host.Host

	holePunchEventsPeers sync.Map
	bpAddrInfos          []peer.AddrInfo
	rcmgr                *ResourceManager
	maddrs               map[string]struct{}

	protocolFiltersLk sync.RWMutex
	protocolFilters   []int32
	natmngr           basichost.NATManager
}

var (
	_ holepunch.EventTracer = (*Host)(nil)
	_ holepunch.AddrFilter  = (*Host)(nil)
)

func InitHost(c *cli.Context, privKey crypto.PrivKey) (*Host, error) {
	log.Info("Starting libp2p host...")

	bpAddrInfos := kaddht.GetDefaultBootstrapPeerAddrInfos()
	if c.IsSet("bootstrap-peers") {
		addrInfos, err := parseBootstrapPeers(c.StringSlice("bootstrap-peers"))
		if err != nil {
			return nil, err
		}
		bpAddrInfos = addrInfos
	}

	rcmgr, err := NewResourceManager()
	if err != nil {
		return nil, errors.Wrap(err, "new resource manager")
	}

	h := &Host{
		holePunchEventsPeers: sync.Map{},
		bpAddrInfos:          bpAddrInfos,
		rcmgr:                rcmgr,
		maddrs:               map[string]struct{}{},
	}
	var nm basichost.NATManager
	// Configure new libp2p host
	libp2pHost, err := libp2p.New(
		libp2p.Identity(privKey),
		libp2p.UserAgent("punchr/go-client/"+c.App.Version),
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/udp/0/quic"),
		libp2p.ListenAddrStrings("/ip6/::/tcp/0"),
		libp2p.ListenAddrStrings("/ip6/::/udp/0/quic"),
		libp2p.EnableHolePunching(holepunch.WithTracer(h), holepunch.WithAddrFilter(h)),
		libp2p.ResourceManager(rcmgr),
		libp2p.NATManager(func(network network.Network) basichost.NATManager {
			nm = basichost.NewNATManager(network)
			return nm
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "new libp2p host")
	}

	h.Host = libp2pHost
	h.natmngr = nm

	return h, nil
}

func parseBootstrapPeers(maddrStrs []string) ([]peer.AddrInfo, error) {
	addrInfos := make([]peer.AddrInfo, len(maddrStrs))
	for i, maddrStr := range maddrStrs {
		maddr, err := multiaddr.NewMultiaddr(maddrStr)
		if err != nil {
			return nil, err
		}

		pi, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			return nil, err
		}

		addrInfos[i] = *pi
	}

	return addrInfos, nil
}

func (h *Host) logEntry(remoteID peer.ID) *log.Entry {
	return log.WithFields(log.Fields{
		"remoteID": util.FmtPeerID(remoteID),
		"hostID":   util.FmtPeerID(h.ID()),
	})
}

// Bootstrap connects this host to bootstrap peers.
func (h *Host) Bootstrap(ctx context.Context) error {
	errCount := 0
	var lastErr error
	for _, bp := range h.bpAddrInfos {
		h.logEntry(bp.ID).Info("Connecting to bootstrap peer...")
		if err := h.Connect(ctx, bp); err != nil {
			h.logEntry(bp.ID).WithError(err).WithField("maddrs", bp.Addrs).Warnln("Error connecting to bootstrap peer")
			errCount++
			lastErr = errors.Wrap(err, "connecting to bootstrap peer")
			continue
		}
		h.logEntry(bp.ID).Info("Connected to bootstrap peer!")
	}

	if errCount == len(h.bpAddrInfos) {
		return lastErr
	}

	return nil
}

// WaitForPublicAddr blocks execution until the host has identified its public address.
// As we currently don't have an event like this, just check our observed addresses
// regularly (exponential backoff starting at 250 ms, capped at 5s).
// TODO: There should be an event here that fires when identify discovers a new address
func (h *Host) WaitForPublicAddr(ctx context.Context) error {
	logEntry := log.WithField("hostID", util.FmtPeerID(h.ID()))
	logEntry.Infoln("Waiting for public address...")

	timeout := time.NewTimer(time.Minute)

	duration := 250 * time.Millisecond
	const maxDuration = 5 * time.Second
	t := time.NewTimer(duration)
	defer t.Stop()
	for {
		if util.ContainsPublicAddr(h.Host.Addrs()) {
			logEntry.Debug("Found >= 1 public addresses!")
			for _, maddr := range h.Host.Addrs() {
				if manet.IsPublicAddr(maddr) {
					h.maddrs[maddr.String()] = struct{}{}
				}
			}

			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			duration *= 2
			if duration > maxDuration {
				duration = maxDuration
			}
			t.Reset(duration)
		case <-timeout.C:
			return fmt.Errorf("timeout wait for public addrs")
		}
	}
}

// GetAgentVersion pulls the agent version from the peer store. Returns nil if no information is available.
func (h *Host) GetAgentVersion(pid peer.ID) *string {
	if value, err := h.Peerstore().Get(pid, "AgentVersion"); err == nil {
		av := value.(string)
		return &av
	} else {
		return nil
	}
}

// GetProtocols pulls the supported protocols of a peer from the peer store. Returns nil if no information is available.
func (h *Host) GetProtocols(pid peer.ID) []string {
	protocols, err := h.Peerstore().GetProtocols(pid)
	if err != nil {
		log.WithError(err).Warnln("Could not get protocols from peerstore")
		return nil
	}
	sort.Strings(protocols)
	return protocols
}

// ProtocolFilters copies the slice
func (h *Host) ProtocolFilters() []int32 {
	h.protocolFiltersLk.RLock()
	defer h.protocolFiltersLk.RUnlock()
	filters := make([]int32, len(h.protocolFilters))
	for i, pf := range h.protocolFilters {
		filters[i] = pf
	}
	return filters
}

func (h *Host) Close() error {
	return h.Host.Close()
}

func (h *Host) MeasurePing(ctx context.Context, pid peer.ID, mType pb.LatencyMeasurementType) <-chan LatencyMeasurement {
	logEntry := log.WithField("remoteID", util.FmtPeerID(pid)).WithField("type", mType.String())

	resultsChan := make(chan LatencyMeasurement)

	go func() {
		defer close(resultsChan)

		tctx, cancel := context.WithTimeout(ctx, PingDuration)
		defer cancel()

		logEntry.Infoln("Measuring ping", mType.String())
		rChan, stream := Ping(tctx, h.Host, pid)
		if stream == nil {
			result := <-rChan
			logEntry.WithError(result.Error).Debugln("Could not ping peer")
			return
		}
		defer logEntry.Debugln("Measuring ping", mType.String(), "done!")

		lm := LatencyMeasurement{
			remoteID: pid,
			mType:    mType,
			conn:     stream.Conn().RemoteMultiaddr(),
			rtts:     []time.Duration{},
			rttErrs:  []error{},
		}

		lm.agentVersion = h.GetAgentVersion(pid)
		lm.protocols = h.GetProtocols(pid)

		for i := 0; i < MaxPingCount; i++ {
			result, ok := <-rChan
			if !ok {
				// channel closed due to e.g. a closed context
				break
			}

			lm.rtts = append(lm.rtts, result.RTT)
			lm.rttErrs = append(lm.rttErrs, result.Error)
		}
		cancel()

		if len(lm.rtts) == 0 {
			// If there are no latency measurements just return.
			// resultsChan is closed via defer above
			return
		}

		resultsChan <- lm
	}()

	return resultsChan
}

func (h *Host) HolePunch(ctx context.Context, addrInfo peer.AddrInfo) (*HolePunchState, <-chan LatencyMeasurement) {
	// sanity operation -> clean up all resources before
	h.prunePeer(addrInfo.ID)

	// register for tracer events for this particular peer
	evtChan := h.RegisterPeerToTrace(addrInfo.ID)
	defer h.UnregisterPeerToTrace(addrInfo.ID)

	// initialize a new hole punch state
	mappings := []nat.Mapping{}
	if h.natmngr.NAT() != nil {
		mappings = h.natmngr.NAT().Mappings()
	}

	hpState := NewHolePunchState(h.ID(), addrInfo.ID, addrInfo.Addrs, h.Addrs(), h.ProtocolFilters(), mappings)
	defer func() { hpState.EndedAt = time.Now() }()

	// Track open connections after the hole punch
	defer func() {
		for _, conn := range h.Network().ConnsToPeer(addrInfo.ID) {
			if !util.IsRelayedMaddr(conn.RemoteMultiaddr()) {
				hpState.HasDirectConns = true
			}
			hpState.OpenMaddrsAfter = append(hpState.OpenMaddrsAfter, conn.RemoteMultiaddr())
		}
	}()

	// connect to the remote peer via relay
	hpState.ConnectStartedAt = time.Now()
	if err := h.Connect(ctx, addrInfo); err != nil {
		h.logEntry(addrInfo.ID).WithError(err).Infoln("Error connecting to remote peer")
		hpState.ConnectEndedAt = time.Now()
		hpState.Error = err.Error()
		hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_NO_CONNECTION
		return hpState, nil
	}
	hpState.ConnectEndedAt = time.Now()
	for _, conn := range h.Network().ConnsToPeer(addrInfo.ID) {
		hpState.OpenMaddrsBefore = append(hpState.OpenMaddrsBefore, conn.RemoteMultiaddr())
	}
	relayedPingChan := h.MeasurePing(ctx, addrInfo.ID, pb.LatencyMeasurementType_TO_REMOTE_THROUGH_RELAY)

	// we were able to connect to the remote peer.
	for i := 0; i < RetryCount; i++ {
		// wait for the DCUtR stream to be opened
		select {
		case _, ok := <-h.WaitForDCUtRStream(addrInfo.ID):
			if !ok {
				// Stream was not opened in time by the remote.
				hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_NO_STREAM
				hpState.Error = "/libp2p/dcutr stream was not opened after " + CommunicationTimeout.String()
				return hpState, relayedPingChan
			}
		case <-ctx.Done():
			hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_CANCELLED
			hpState.Error = ctx.Err().Error()
			return hpState, relayedPingChan
		}
		// stream was opened! Now, wait for the first hole punch event.

		hpa := hpState.TrackHolePunch(ctx, addrInfo.ID, evtChan)
		hpState.HolePunchAttempts = append(hpState.HolePunchAttempts, &hpa)

		switch hpa.Outcome {
		case pb.HolePunchAttemptOutcome_HOLE_PUNCH_ATTEMPT_OUTCOME_PROTOCOL_ERROR:
			hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_FAILED
			hpState.Error = hpa.Error
			return hpState, relayedPingChan
		case pb.HolePunchAttemptOutcome_HOLE_PUNCH_ATTEMPT_OUTCOME_DIRECT_DIAL:
			hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_SUCCESS
			return hpState, relayedPingChan
		case pb.HolePunchAttemptOutcome_HOLE_PUNCH_ATTEMPT_OUTCOME_UNKNOWN:
			hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_UNKNOWN
			hpState.Error = "unknown hole punch attempt outcome"
			return hpState, relayedPingChan
		case pb.HolePunchAttemptOutcome_HOLE_PUNCH_ATTEMPT_OUTCOME_CANCELLED:
			hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_CANCELLED
			hpState.Error = hpa.Error
			return hpState, relayedPingChan
		case pb.HolePunchAttemptOutcome_HOLE_PUNCH_ATTEMPT_OUTCOME_TIMEOUT:
			hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_FAILED
			hpState.Error = hpa.Error
			return hpState, relayedPingChan
		case pb.HolePunchAttemptOutcome_HOLE_PUNCH_ATTEMPT_OUTCOME_SUCCESS:
			hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_SUCCESS
			return hpState, relayedPingChan
		}
	}

	hpState.Outcome = pb.HolePunchOutcome_HOLE_PUNCH_OUTCOME_FAILED
	hpState.Error = fmt.Sprintf("none of the %d attempts succeeded", RetryCount)

	return hpState, relayedPingChan
}

type LatencyMeasurement struct {
	remoteID     peer.ID
	agentVersion *string
	protocols    []string
	mType        pb.LatencyMeasurementType
	conn         multiaddr.Multiaddr
	rtts         []time.Duration
	rttErrs      []error
}

func (h *Host) PingRelays(ctx context.Context, addrInfo map[peer.ID]*peer.AddrInfo) []LatencyMeasurement {
	var lats []LatencyMeasurement

	for relayID, ri := range addrInfo {
		if err := h.Connect(ctx, *ri); err != nil {
			h.logEntry(relayID).WithError(err).WithField("maddrs", ri.Addrs).Info("Could not connect to relay peer")
			continue
		}

		if lm, ok := <-h.MeasurePing(ctx, relayID, pb.LatencyMeasurementType_TO_RELAY); ok {
			lats = append(lats, lm)
		}
	}

	return lats
}

func (hps HolePunchState) TrackHolePunch(ctx context.Context, remoteID peer.ID, evtChan <-chan *holepunch.Event) HolePunchAttempt {
	hps.logEntry(remoteID).Infoln("Waiting for hole punch events...")
	hpa := &HolePunchAttempt{
		HostID:   hps.HostID,
		RemoteID: remoteID,
		OpenedAt: time.Now(),
	}

	for {
		select {
		case evt := <-evtChan:
			switch event := evt.Evt.(type) {
			case *holepunch.StartHolePunchEvt:
				hpa.handleStartHolePunchEvt(evt, event)
			case *holepunch.EndHolePunchEvt:
				hpa.handleEndHolePunchEvt(evt, event)
				return *hpa
			case *holepunch.HolePunchAttemptEvt:
				hpa.handleHolePunchAttemptEvt(event)
			case *holepunch.ProtocolErrorEvt:
				hpa.handleProtocolErrorEvt(evt, event)
				return *hpa
			case *holepunch.DirectDialEvt:
				hpa.handleDirectDialEvt(evt, event)
				if event.Success {
					return *hpa
				}
			default:
				panic(fmt.Sprintf("unexpected event %T", evt.Evt))
			}
		case <-time.After(CommunicationTimeout):
			hpa.handleHolePunchTimeout()
			return *hpa
		case <-ctx.Done():
			hpa.handleHolePunchCancelled(ctx.Err())
			return *hpa
		}
	}
}

func (h *Host) WaitForDCUtRStream(pid peer.ID) <-chan struct{} {
	dcutrOpenedChan := make(chan struct{})

	go func() {
		openedStream := h.rcmgr.Register(pid)
		defer h.rcmgr.Unregister(pid)

		for _, conn := range h.Network().ConnsToPeer(pid) {
			for _, stream := range conn.GetStreams() {
				if stream.Protocol() == holepunch.Protocol {
					dcutrOpenedChan <- struct{}{}
					close(dcutrOpenedChan)
					return
				}
			}
		}

		select {
		case <-time.After(CommunicationTimeout):
			h.logEntry(pid).Infoln("/libp2p/dcutr stream was not opened after " + CommunicationTimeout.String())
		case <-openedStream:
			h.logEntry(pid).Infoln("/libp2p/dcutr stream opened!")
			dcutrOpenedChan <- struct{}{}
		}
		close(dcutrOpenedChan)
	}()

	h.logEntry(pid).Infoln("Waiting for /libp2p/dcutr stream...")
	return dcutrOpenedChan
}

func (h *Host) RegisterPeerToTrace(pid peer.ID) <-chan *holepunch.Event {
	evtChan := make(chan *holepunch.Event, 3) // Start, attempt, end -> so that we get the events in this order!
	h.holePunchEventsPeers.Store(pid, evtChan)
	return evtChan
}

func (h *Host) UnregisterPeerToTrace(pid peer.ID) {
	val, found := h.holePunchEventsPeers.LoadAndDelete(pid)
	if !found {
		return
	}
	evtChan := val.(chan *holepunch.Event)

	// Drain channel
	for {
		select {
		case evt := <-evtChan:
			h.logEntry(pid).WithField("evtType", evt.Type).Warnln("Draining event channel")
		default:
			close(evtChan)
			return
		}
	}
}

// logAddrInfo logs address information about the given peer
func (h *Host) logAddrInfo(addrInfo peer.AddrInfo) {
	h.logEntry(addrInfo.ID).WithField("openConns", len(h.Network().ConnsToPeer(addrInfo.ID))).Infoln("Connecting to remote peer...")
	for i, maddr := range addrInfo.Addrs {
		log.Infoln("  ["+strconv.Itoa(i)+"]", maddr.String())
	}
}

// prunePeer closes all connections to the given peer and removes all information about it from the peer store.
func (h *Host) prunePeer(pid peer.ID) {
	if err := h.Network().ClosePeer(pid); err != nil {
		h.logEntry(pid).WithError(err).Warnln("Error closing connection")
	}
	h.Peerstore().RemovePeer(pid)
	h.Peerstore().ClearAddrs(pid)
}

// networkInformation checks if the multi addresses have changed - if that's the case we have likely switched networks
func (h *Host) networkInformation(ctx context.Context) *pb.NetworkInformation {
	for _, maddr := range h.Addrs() {
		if _, found := h.maddrs[maddr.String()]; found {
			continue
		}

		ni := &pb.NetworkInformation{}

		log.Infoln("Found new multi addresses - fetching Router Login")
		html, err := util.DefaultGatewayHTML(ctx)
		if err != nil {
			errStr := err.Error()
			ni.RouterLoginHtmlError = &errStr
		}
		ni.RouterLoginHtml = &html

		// TODO: check IPv6 support

		// Update list of multi addresses
		h.maddrs = map[string]struct{}{}
		for _, newMaddr := range h.Addrs() {
			h.maddrs[newMaddr.String()] = struct{}{}
		}

		return ni
	}

	return nil
}

// Trace is called during the hole punching process
func (h *Host) Trace(evt *holepunch.Event) {
	val, found := h.holePunchEventsPeers.Load(evt.Remote)
	if !found {
		h.logEntry(evt.Remote).Infoln("Tracer event for untracked peer")
		return
	}

	val.(chan *holepunch.Event) <- evt
}

func (h *Host) FilterLocal(remoteID peer.ID, maddrs []multiaddr.Multiaddr) []multiaddr.Multiaddr {
	return h.filter(remoteID, maddrs)
}

func (h *Host) FilterRemote(remoteID peer.ID, maddrs []multiaddr.Multiaddr) []multiaddr.Multiaddr {
	return h.filter(remoteID, maddrs)
}

func (h *Host) filter(remoteID peer.ID, maddrs []multiaddr.Multiaddr) []multiaddr.Multiaddr {
	h.protocolFiltersLk.RLock()
	defer h.protocolFiltersLk.RUnlock()

	result := make([]multiaddr.Multiaddr, 0, len(maddrs))
	onlyRelayFiltered := make([]multiaddr.Multiaddr, 0, len(maddrs))
	for _, maddr := range maddrs {
		if util.IsRelayedMaddr(maddr) {
			continue
		}

		onlyRelayFiltered = append(onlyRelayFiltered, maddr)

		// If there's no filter -> add all maddrs
		if len(h.protocolFilters) == 0 {
			result = append(result, maddr)
			continue
		}

		// If there's a filter -> only add those that match all protocols
		matchesAllFilters := true
		for _, p := range h.protocolFilters {
			if _, err := maddr.ValueForProtocol(int(p)); err != nil {
				matchesAllFilters = false
				break
			}
		}

		if matchesAllFilters {
			result = append(result, maddr)
		}
	}

	if len(result) == 0 {
		return onlyRelayFiltered
	}

	return result
}
