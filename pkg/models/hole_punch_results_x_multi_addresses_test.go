// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testHolePunchResultsXMultiAddresses(t *testing.T) {
	t.Parallel()

	query := HolePunchResultsXMultiAddresses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHolePunchResultsXMultiAddressesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHolePunchResultsXMultiAddressesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := HolePunchResultsXMultiAddresses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHolePunchResultsXMultiAddressesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HolePunchResultsXMultiAddressSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHolePunchResultsXMultiAddressesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HolePunchResultsXMultiAddressExists(ctx, tx, o.MultiAddressID, o.HolePunchResultID, o.Relationship)
	if err != nil {
		t.Errorf("Unable to check if HolePunchResultsXMultiAddress exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HolePunchResultsXMultiAddressExists to return true, but got false.")
	}
}

func testHolePunchResultsXMultiAddressesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	holePunchResultsXMultiAddressFound, err := FindHolePunchResultsXMultiAddress(ctx, tx, o.MultiAddressID, o.HolePunchResultID, o.Relationship)
	if err != nil {
		t.Error(err)
	}

	if holePunchResultsXMultiAddressFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHolePunchResultsXMultiAddressesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = HolePunchResultsXMultiAddresses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHolePunchResultsXMultiAddressesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := HolePunchResultsXMultiAddresses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHolePunchResultsXMultiAddressesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	holePunchResultsXMultiAddressOne := &HolePunchResultsXMultiAddress{}
	holePunchResultsXMultiAddressTwo := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, holePunchResultsXMultiAddressOne, holePunchResultsXMultiAddressDBTypes, false, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, holePunchResultsXMultiAddressTwo, holePunchResultsXMultiAddressDBTypes, false, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = holePunchResultsXMultiAddressOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = holePunchResultsXMultiAddressTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := HolePunchResultsXMultiAddresses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHolePunchResultsXMultiAddressesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	holePunchResultsXMultiAddressOne := &HolePunchResultsXMultiAddress{}
	holePunchResultsXMultiAddressTwo := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, holePunchResultsXMultiAddressOne, holePunchResultsXMultiAddressDBTypes, false, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, holePunchResultsXMultiAddressTwo, holePunchResultsXMultiAddressDBTypes, false, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = holePunchResultsXMultiAddressOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = holePunchResultsXMultiAddressTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func holePunchResultsXMultiAddressBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func holePunchResultsXMultiAddressAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func holePunchResultsXMultiAddressAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func holePunchResultsXMultiAddressBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func holePunchResultsXMultiAddressAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func holePunchResultsXMultiAddressBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func holePunchResultsXMultiAddressAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func holePunchResultsXMultiAddressBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func holePunchResultsXMultiAddressAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *HolePunchResultsXMultiAddress) error {
	*o = HolePunchResultsXMultiAddress{}
	return nil
}

func testHolePunchResultsXMultiAddressesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &HolePunchResultsXMultiAddress{}
	o := &HolePunchResultsXMultiAddress{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, false); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress object: %s", err)
	}

	AddHolePunchResultsXMultiAddressHook(boil.BeforeInsertHook, holePunchResultsXMultiAddressBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressBeforeInsertHooks = []HolePunchResultsXMultiAddressHook{}

	AddHolePunchResultsXMultiAddressHook(boil.AfterInsertHook, holePunchResultsXMultiAddressAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressAfterInsertHooks = []HolePunchResultsXMultiAddressHook{}

	AddHolePunchResultsXMultiAddressHook(boil.AfterSelectHook, holePunchResultsXMultiAddressAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressAfterSelectHooks = []HolePunchResultsXMultiAddressHook{}

	AddHolePunchResultsXMultiAddressHook(boil.BeforeUpdateHook, holePunchResultsXMultiAddressBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressBeforeUpdateHooks = []HolePunchResultsXMultiAddressHook{}

	AddHolePunchResultsXMultiAddressHook(boil.AfterUpdateHook, holePunchResultsXMultiAddressAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressAfterUpdateHooks = []HolePunchResultsXMultiAddressHook{}

	AddHolePunchResultsXMultiAddressHook(boil.BeforeDeleteHook, holePunchResultsXMultiAddressBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressBeforeDeleteHooks = []HolePunchResultsXMultiAddressHook{}

	AddHolePunchResultsXMultiAddressHook(boil.AfterDeleteHook, holePunchResultsXMultiAddressAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressAfterDeleteHooks = []HolePunchResultsXMultiAddressHook{}

	AddHolePunchResultsXMultiAddressHook(boil.BeforeUpsertHook, holePunchResultsXMultiAddressBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressBeforeUpsertHooks = []HolePunchResultsXMultiAddressHook{}

	AddHolePunchResultsXMultiAddressHook(boil.AfterUpsertHook, holePunchResultsXMultiAddressAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	holePunchResultsXMultiAddressAfterUpsertHooks = []HolePunchResultsXMultiAddressHook{}
}

func testHolePunchResultsXMultiAddressesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHolePunchResultsXMultiAddressesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(holePunchResultsXMultiAddressColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHolePunchResultsXMultiAddressToOneHolePunchResultUsingHolePunchResult(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local HolePunchResultsXMultiAddress
	var foreign HolePunchResult

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, holePunchResultsXMultiAddressDBTypes, false, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResult struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.HolePunchResultID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.HolePunchResult().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := HolePunchResultsXMultiAddressSlice{&local}
	if err = local.L.LoadHolePunchResult(ctx, tx, false, (*[]*HolePunchResultsXMultiAddress)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.HolePunchResult == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.HolePunchResult = nil
	if err = local.L.LoadHolePunchResult(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.HolePunchResult == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testHolePunchResultsXMultiAddressToOneMultiAddressUsingMultiAddress(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local HolePunchResultsXMultiAddress
	var foreign MultiAddress

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, holePunchResultsXMultiAddressDBTypes, false, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, multiAddressDBTypes, false, multiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddress struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MultiAddressID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.MultiAddress().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := HolePunchResultsXMultiAddressSlice{&local}
	if err = local.L.LoadMultiAddress(ctx, tx, false, (*[]*HolePunchResultsXMultiAddress)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.MultiAddress == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.MultiAddress = nil
	if err = local.L.LoadMultiAddress(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.MultiAddress == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testHolePunchResultsXMultiAddressToOneSetOpHolePunchResultUsingHolePunchResult(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a HolePunchResultsXMultiAddress
	var b, c HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, holePunchResultsXMultiAddressDBTypes, false, strmangle.SetComplement(holePunchResultsXMultiAddressPrimaryKeyColumns, holePunchResultsXMultiAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*HolePunchResult{&b, &c} {
		err = a.SetHolePunchResult(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.HolePunchResult != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.HolePunchResultsXMultiAddresses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.HolePunchResultID != x.ID {
			t.Error("foreign key was wrong value", a.HolePunchResultID)
		}

		if exists, err := HolePunchResultsXMultiAddressExists(ctx, tx, a.MultiAddressID, a.HolePunchResultID, a.Relationship); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testHolePunchResultsXMultiAddressToOneSetOpMultiAddressUsingMultiAddress(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a HolePunchResultsXMultiAddress
	var b, c MultiAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, holePunchResultsXMultiAddressDBTypes, false, strmangle.SetComplement(holePunchResultsXMultiAddressPrimaryKeyColumns, holePunchResultsXMultiAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, multiAddressDBTypes, false, strmangle.SetComplement(multiAddressPrimaryKeyColumns, multiAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, multiAddressDBTypes, false, strmangle.SetComplement(multiAddressPrimaryKeyColumns, multiAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MultiAddress{&b, &c} {
		err = a.SetMultiAddress(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.MultiAddress != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.HolePunchResultsXMultiAddresses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MultiAddressID != x.ID {
			t.Error("foreign key was wrong value", a.MultiAddressID)
		}

		if exists, err := HolePunchResultsXMultiAddressExists(ctx, tx, a.MultiAddressID, a.HolePunchResultID, a.Relationship); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testHolePunchResultsXMultiAddressesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHolePunchResultsXMultiAddressesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HolePunchResultsXMultiAddressSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHolePunchResultsXMultiAddressesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := HolePunchResultsXMultiAddresses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	holePunchResultsXMultiAddressDBTypes = map[string]string{`HolePunchResultID`: `integer`, `MultiAddressID`: `bigint`, `Relationship`: `enum.hole_punch_multi_address_relationship('INITIAL','FINAL')`}
	_                                    = bytes.MinRead
)

func testHolePunchResultsXMultiAddressesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(holePunchResultsXMultiAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(holePunchResultsXMultiAddressAllColumns) == len(holePunchResultsXMultiAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHolePunchResultsXMultiAddressesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(holePunchResultsXMultiAddressAllColumns) == len(holePunchResultsXMultiAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, holePunchResultsXMultiAddressDBTypes, true, holePunchResultsXMultiAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(holePunchResultsXMultiAddressAllColumns, holePunchResultsXMultiAddressPrimaryKeyColumns) {
		fields = holePunchResultsXMultiAddressAllColumns
	} else {
		fields = strmangle.SetComplement(
			holePunchResultsXMultiAddressAllColumns,
			holePunchResultsXMultiAddressPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := HolePunchResultsXMultiAddressSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHolePunchResultsXMultiAddressesUpsert(t *testing.T) {
	t.Parallel()

	if len(holePunchResultsXMultiAddressAllColumns) == len(holePunchResultsXMultiAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := HolePunchResultsXMultiAddress{}
	if err = randomize.Struct(seed, &o, holePunchResultsXMultiAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert HolePunchResultsXMultiAddress: %s", err)
	}

	count, err := HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, holePunchResultsXMultiAddressDBTypes, false, holePunchResultsXMultiAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize HolePunchResultsXMultiAddress struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert HolePunchResultsXMultiAddress: %s", err)
	}

	count, err = HolePunchResultsXMultiAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}