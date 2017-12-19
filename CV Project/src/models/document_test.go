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

func testDocuments(t *testing.T) {
	t.Parallel()

	query := Documents(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDocumentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = document.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDocumentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Documents(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDocumentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DocumentSlice{document}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDocumentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DocumentExists(tx, document.DocumentID)
	if err != nil {
		t.Errorf("Unable to check if Document exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DocumentExistsG to return true, but got false.")
	}
}
func testDocumentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	documentFound, err := FindDocument(tx, document.DocumentID)
	if err != nil {
		t.Error(err)
	}

	if documentFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDocumentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Documents(tx).Bind(document); err != nil {
		t.Error(err)
	}
}

func testDocumentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Documents(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDocumentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	documentOne := &Document{}
	documentTwo := &Document{}
	if err = randomize.Struct(seed, documentOne, documentDBTypes, false, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}
	if err = randomize.Struct(seed, documentTwo, documentDBTypes, false, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = documentOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = documentTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Documents(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDocumentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	documentOne := &Document{}
	documentTwo := &Document{}
	if err = randomize.Struct(seed, documentOne, documentDBTypes, false, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}
	if err = randomize.Struct(seed, documentTwo, documentDBTypes, false, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = documentOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = documentTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func documentBeforeInsertHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func documentAfterInsertHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func documentAfterSelectHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func documentBeforeUpdateHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func documentAfterUpdateHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func documentBeforeDeleteHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func documentAfterDeleteHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func documentBeforeUpsertHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func documentAfterUpsertHook(e boil.Executor, o *Document) error {
	*o = Document{}
	return nil
}

func testDocumentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Document{}
	o := &Document{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, documentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Document object: %s", err)
	}

	AddDocumentHook(boil.BeforeInsertHook, documentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	documentBeforeInsertHooks = []DocumentHook{}

	AddDocumentHook(boil.AfterInsertHook, documentAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	documentAfterInsertHooks = []DocumentHook{}

	AddDocumentHook(boil.AfterSelectHook, documentAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	documentAfterSelectHooks = []DocumentHook{}

	AddDocumentHook(boil.BeforeUpdateHook, documentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	documentBeforeUpdateHooks = []DocumentHook{}

	AddDocumentHook(boil.AfterUpdateHook, documentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	documentAfterUpdateHooks = []DocumentHook{}

	AddDocumentHook(boil.BeforeDeleteHook, documentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	documentBeforeDeleteHooks = []DocumentHook{}

	AddDocumentHook(boil.AfterDeleteHook, documentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	documentAfterDeleteHooks = []DocumentHook{}

	AddDocumentHook(boil.BeforeUpsertHook, documentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	documentBeforeUpsertHooks = []DocumentHook{}

	AddDocumentHook(boil.AfterUpsertHook, documentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	documentAfterUpsertHooks = []DocumentHook{}
}
func testDocumentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDocumentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx, documentColumns...); err != nil {
		t.Error(err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDocumentToManyPages(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Document
	var b, c Page

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, pageDBTypes, false, pageColumnsWithDefault...)
	randomize.Struct(seed, &c, pageDBTypes, false, pageColumnsWithDefault...)

	b.DocumentID.Valid = true
	c.DocumentID.Valid = true
	b.DocumentID.String = a.DocumentID
	c.DocumentID.String = a.DocumentID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	page, err := a.Pages(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range page {
		if v.DocumentID.String == b.DocumentID.String {
			bFound = true
		}
		if v.DocumentID.String == c.DocumentID.String {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DocumentSlice{&a}
	if err = a.L.LoadPages(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Pages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Pages = nil
	if err = a.L.LoadPages(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Pages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", page)
	}
}

func testDocumentToManyAddOpPages(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Document
	var b, c, d, e Page

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, documentDBTypes, false, strmangle.SetComplement(documentPrimaryKeyColumns, documentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Page{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pageDBTypes, false, strmangle.SetComplement(pagePrimaryKeyColumns, pageColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Page{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPages(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.DocumentID != first.DocumentID.String {
			t.Error("foreign key was wrong value", a.DocumentID, first.DocumentID.String)
		}
		if a.DocumentID != second.DocumentID.String {
			t.Error("foreign key was wrong value", a.DocumentID, second.DocumentID.String)
		}

		if first.R.Document != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Document != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Pages[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Pages[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Pages(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDocumentToManySetOpPages(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Document
	var b, c, d, e Page

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, documentDBTypes, false, strmangle.SetComplement(documentPrimaryKeyColumns, documentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Page{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pageDBTypes, false, strmangle.SetComplement(pagePrimaryKeyColumns, pageColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetPages(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Pages(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetPages(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Pages(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.DocumentID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.DocumentID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.DocumentID != d.DocumentID.String {
		t.Error("foreign key was wrong value", a.DocumentID, d.DocumentID.String)
	}
	if a.DocumentID != e.DocumentID.String {
		t.Error("foreign key was wrong value", a.DocumentID, e.DocumentID.String)
	}

	if b.R.Document != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Document != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Document != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Document != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Pages[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Pages[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDocumentToManyRemoveOpPages(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Document
	var b, c, d, e Page

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, documentDBTypes, false, strmangle.SetComplement(documentPrimaryKeyColumns, documentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Page{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pageDBTypes, false, strmangle.SetComplement(pagePrimaryKeyColumns, pageColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddPages(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Pages(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemovePages(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Pages(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.DocumentID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.DocumentID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Document != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Document != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Document != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Document != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Pages) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Pages[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Pages[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDocumentToOneOwnerUsingOwner(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Document
	var foreign Owner

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	local.OwnerID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OwnerID.String = foreign.OwnerID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Owner(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.OwnerID != foreign.OwnerID {
		t.Errorf("want: %v, got %v", foreign.OwnerID, check.OwnerID)
	}

	slice := DocumentSlice{&local}
	if err = local.L.LoadOwner(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Owner == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Owner = nil
	if err = local.L.LoadOwner(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Owner == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDocumentToOneSetOpOwnerUsingOwner(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Document
	var b, c Owner

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, documentDBTypes, false, strmangle.SetComplement(documentPrimaryKeyColumns, documentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, ownerDBTypes, false, strmangle.SetComplement(ownerPrimaryKeyColumns, ownerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, ownerDBTypes, false, strmangle.SetComplement(ownerPrimaryKeyColumns, ownerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Owner{&b, &c} {
		err = a.SetOwner(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Owner != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Documents[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OwnerID.String != x.OwnerID {
			t.Error("foreign key was wrong value", a.OwnerID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OwnerID.String))
		reflect.Indirect(reflect.ValueOf(&a.OwnerID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OwnerID.String != x.OwnerID {
			t.Error("foreign key was wrong value", a.OwnerID.String, x.OwnerID)
		}
	}
}

func testDocumentToOneRemoveOpOwnerUsingOwner(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Document
	var b Owner

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, documentDBTypes, false, strmangle.SetComplement(documentPrimaryKeyColumns, documentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, ownerDBTypes, false, strmangle.SetComplement(ownerPrimaryKeyColumns, ownerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetOwner(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveOwner(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Owner(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Owner != nil {
		t.Error("R struct entry should be nil")
	}

	if a.OwnerID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Documents) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDocumentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = document.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDocumentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DocumentSlice{document}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDocumentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Documents(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	documentDBTypes = map[string]string{`DocHash`: `character`, `DocType`: `character`, `DocumentID`: `uuid`, `OwnerID`: `uuid`, `Status`: `character`}
	_               = bytes.MinRead
)

func testDocumentsUpdate(t *testing.T) {
	t.Parallel()

	if len(documentColumns) == len(documentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, document, documentDBTypes, true, documentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	if err = document.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDocumentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(documentColumns) == len(documentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	document := &Document{}
	if err = randomize.Struct(seed, document, documentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, document, documentDBTypes, true, documentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(documentColumns, documentPrimaryKeyColumns) {
		fields = documentColumns
	} else {
		fields = strmangle.SetComplement(
			documentColumns,
			documentPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(document))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DocumentSlice{document}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDocumentsUpsert(t *testing.T) {
	t.Parallel()

	if len(documentColumns) == len(documentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	document := Document{}
	if err = randomize.Struct(seed, &document, documentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = document.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Document: %s", err)
	}

	count, err := Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &document, documentDBTypes, false, documentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Document struct: %s", err)
	}

	if err = document.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Document: %s", err)
	}

	count, err = Documents(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
