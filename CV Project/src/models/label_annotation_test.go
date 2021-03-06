// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testLabelAnnotations(t *testing.T) {
	t.Parallel()

	query := LabelAnnotations(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testLabelAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = labelAnnotation.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLabelAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = LabelAnnotations(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLabelAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := LabelAnnotationSlice{labelAnnotation}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testLabelAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := LabelAnnotationExists(tx, labelAnnotation.LabelID)
	if err != nil {
		t.Errorf("Unable to check if LabelAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LabelAnnotationExistsG to return true, but got false.")
	}
}
func testLabelAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	labelAnnotationFound, err := FindLabelAnnotation(tx, labelAnnotation.LabelID)
	if err != nil {
		t.Error(err)
	}

	if labelAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}
func testLabelAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = LabelAnnotations(tx).Bind(labelAnnotation); err != nil {
		t.Error(err)
	}
}

func testLabelAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := LabelAnnotations(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLabelAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotationOne := &LabelAnnotation{}
	labelAnnotationTwo := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotationOne, labelAnnotationDBTypes, false, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, labelAnnotationTwo, labelAnnotationDBTypes, false, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = labelAnnotationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := LabelAnnotations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLabelAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	labelAnnotationOne := &LabelAnnotation{}
	labelAnnotationTwo := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotationOne, labelAnnotationDBTypes, false, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, labelAnnotationTwo, labelAnnotationDBTypes, false, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = labelAnnotationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func labelAnnotationBeforeInsertHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func labelAnnotationAfterInsertHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func labelAnnotationAfterSelectHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func labelAnnotationBeforeUpdateHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func labelAnnotationAfterUpdateHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func labelAnnotationBeforeDeleteHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func labelAnnotationAfterDeleteHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func labelAnnotationBeforeUpsertHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func labelAnnotationAfterUpsertHook(e boil.Executor, o *LabelAnnotation) error {
	*o = LabelAnnotation{}
	return nil
}

func testLabelAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &LabelAnnotation{}
	o := &LabelAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, labelAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation object: %s", err)
	}

	AddLabelAnnotationHook(boil.BeforeInsertHook, labelAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	labelAnnotationBeforeInsertHooks = []LabelAnnotationHook{}

	AddLabelAnnotationHook(boil.AfterInsertHook, labelAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	labelAnnotationAfterInsertHooks = []LabelAnnotationHook{}

	AddLabelAnnotationHook(boil.AfterSelectHook, labelAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	labelAnnotationAfterSelectHooks = []LabelAnnotationHook{}

	AddLabelAnnotationHook(boil.BeforeUpdateHook, labelAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	labelAnnotationBeforeUpdateHooks = []LabelAnnotationHook{}

	AddLabelAnnotationHook(boil.AfterUpdateHook, labelAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	labelAnnotationAfterUpdateHooks = []LabelAnnotationHook{}

	AddLabelAnnotationHook(boil.BeforeDeleteHook, labelAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	labelAnnotationBeforeDeleteHooks = []LabelAnnotationHook{}

	AddLabelAnnotationHook(boil.AfterDeleteHook, labelAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	labelAnnotationAfterDeleteHooks = []LabelAnnotationHook{}

	AddLabelAnnotationHook(boil.BeforeUpsertHook, labelAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	labelAnnotationBeforeUpsertHooks = []LabelAnnotationHook{}

	AddLabelAnnotationHook(boil.AfterUpsertHook, labelAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	labelAnnotationAfterUpsertHooks = []LabelAnnotationHook{}
}
func testLabelAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLabelAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx, labelAnnotationColumns...); err != nil {
		t.Error(err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLabelAnnotationToOnePageUsingPage(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local LabelAnnotation
	var foreign Page

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pageDBTypes, true, pageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Page struct: %s", err)
	}

	local.PageID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PageID.String = foreign.PageID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Page(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.PageID != foreign.PageID {
		t.Errorf("want: %v, got %v", foreign.PageID, check.PageID)
	}

	slice := LabelAnnotationSlice{&local}
	if err = local.L.LoadPage(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Page == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Page = nil
	if err = local.L.LoadPage(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Page == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLabelAnnotationToOneSetOpPageUsingPage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a LabelAnnotation
	var b, c Page

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, labelAnnotationDBTypes, false, strmangle.SetComplement(labelAnnotationPrimaryKeyColumns, labelAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pageDBTypes, false, strmangle.SetComplement(pagePrimaryKeyColumns, pageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pageDBTypes, false, strmangle.SetComplement(pagePrimaryKeyColumns, pageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Page{&b, &c} {
		err = a.SetPage(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Page != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.LabelAnnotations[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PageID.String != x.PageID {
			t.Error("foreign key was wrong value", a.PageID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PageID.String))
		reflect.Indirect(reflect.ValueOf(&a.PageID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PageID.String != x.PageID {
			t.Error("foreign key was wrong value", a.PageID.String, x.PageID)
		}
	}
}

func testLabelAnnotationToOneRemoveOpPageUsingPage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a LabelAnnotation
	var b Page

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, labelAnnotationDBTypes, false, strmangle.SetComplement(labelAnnotationPrimaryKeyColumns, labelAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pageDBTypes, false, strmangle.SetComplement(pagePrimaryKeyColumns, pageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPage(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePage(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Page(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Page != nil {
		t.Error("R struct entry should be nil")
	}

	if a.PageID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.LabelAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testLabelAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = labelAnnotation.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testLabelAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := LabelAnnotationSlice{labelAnnotation}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testLabelAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := LabelAnnotations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	labelAnnotationDBTypes = map[string]string{`Description`: `character`, `LabelID`: `uuid`, `PageID`: `uuid`, `Score`: `double precision`}
	_                      = bytes.MinRead
)

func testLabelAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if len(labelAnnotationColumns) == len(labelAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	if err = labelAnnotation.Update(tx); err != nil {
		t.Error(err)
	}
}

func testLabelAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(labelAnnotationColumns) == len(labelAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	labelAnnotation := &LabelAnnotation{}
	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, labelAnnotation, labelAnnotationDBTypes, true, labelAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(labelAnnotationColumns, labelAnnotationPrimaryKeyColumns) {
		fields = labelAnnotationColumns
	} else {
		fields = strmangle.SetComplement(
			labelAnnotationColumns,
			labelAnnotationPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(labelAnnotation))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := LabelAnnotationSlice{labelAnnotation}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testLabelAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(labelAnnotationColumns) == len(labelAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	labelAnnotation := LabelAnnotation{}
	if err = randomize.Struct(seed, &labelAnnotation, labelAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = labelAnnotation.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert LabelAnnotation: %s", err)
	}

	count, err := LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &labelAnnotation, labelAnnotationDBTypes, false, labelAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LabelAnnotation struct: %s", err)
	}

	if err = labelAnnotation.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert LabelAnnotation: %s", err)
	}

	count, err = LabelAnnotations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
