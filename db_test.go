package mifer_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/lll-lll-lll-lll/mifer"
	"gotest.tools/assert"
)

func Test_Join(t *testing.T) {
	t.Parallel()
	ipt := []string{"id", "name", "email"}
	want := "id, name, email"
	got := strings.Join(ipt, ", ")
	assert.Equal(t, got, want)
}

func Text_JoinOptions(t *testing.T) {
	t.Parallel()
	ipt := []mifer.MiferOption{
		{ColumnKey: "id", Datum: []interface{}{[]int{1, 2}}},
		{ColumnKey: "name", Datum: []interface{}{[]string{"test1", "test2"}}}}
	want := "id, name"
	got := mifer.JoinOptions(ipt)
	assert.Equal(t, got, want)
}

func Test_Build(t *testing.T) {
	t.Parallel()
	psql := &mifer.PostreSQL{
		DBName:    "testDB",
		TableName: "users",
	}
	t.Run("single query", func(t *testing.T) {
		t.Parallel()
		randomID := mifer.DefaultInt64PrepareDataCallBack()
		randaomName := mifer.DefaultUserNamePrepareDataCallBack()
		clmns := mifer.Columns{"id": mifer.Column{ColumnType: "int"}, "name": mifer.Column{ColumnType: "nvarchar"}}
		queries, err := psql.BuildQueries(context.Background(), 10, clmns,
			mifer.MiferOption{"id", []interface{}{randomID}},
			mifer.MiferOption{"name", []interface{}{randaomName}})
		if err != nil {
			t.Log(err)
		}
		want := []string{fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, '%s');", randomID, randaomName)}
		assert.Equal(t, queries[0], want[0])
	})

	t.Run("multi queries", func(t *testing.T) {
		t.Parallel()
		randomID := mifer.DefaultInt64PrepareDataCallBack()
		randomID2 := mifer.DefaultInt64PrepareDataCallBack()
		randomID3 := mifer.DefaultInt64PrepareDataCallBack()
		randaomName := mifer.DefaultUserNamePrepareDataCallBack()
		randaomName2 := mifer.DefaultUserNamePrepareDataCallBack()
		randaomName3 := mifer.DefaultUserNamePrepareDataCallBack()
		clmns := mifer.Columns{"id": mifer.Column{ColumnType: "int"}, "name": mifer.Column{ColumnType: "nvarchar"}}
		queries, err := psql.BuildQueries(context.Background(), 10, clmns,
			mifer.MiferOption{"id", []interface{}{randomID, randomID2, randomID3}},
			mifer.MiferOption{"name", []interface{}{randaomName, randaomName2, randaomName3}})
		if err != nil {
			t.Log(err)
		}
		want := []string{
			fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, '%s');", randomID, randaomName),
			fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, '%s');", randomID2, randaomName2),
			fmt.Sprintf("INSERT INTO users (id, name) VALUES (%d, '%s');", randomID3, randaomName3),
		}
		assert.Equal(t, queries[0], want[0])
		assert.Equal(t, queries[1], want[1])
		assert.Equal(t, queries[2], want[2])
	})

	t.Run("one column and multi queries", func(t *testing.T) {
		t.Parallel()
		randomID := mifer.DefaultInt64PrepareDataCallBack()
		column := mifer.Columns{"id": mifer.Column{ColumnType: "int"}}
		queries, err := psql.BuildQueries(context.Background(), 10, column,
			mifer.MiferOption{"id", []interface{}{randomID}})
		if err != nil {
			t.Log(err)
		}
		want := fmt.Sprintf("INSERT INTO users (id) VALUES (%d);", randomID)
		assert.Equal(t, queries[0], want)
	})

	t.Run(" no options error", func(t *testing.T) {
		t.Parallel()
		column := mifer.Columns{"id": mifer.Column{ColumnType: "int"}}
		_, err := psql.BuildQueries(context.Background(), 10, column)
		e, ok := err.(*mifer.MiferError)
		if !ok {
			t.Errorf("expected *mifer.MiferError, got %T", err)
		}
		wantT := mifer.NoOptionsErr
		if e.ErrType != mifer.NoOptionsErr {
			t.Errorf("expected error type %q, got %q", wantT, e.ErrType)
		}
		wantMsg := "Not a option was provided. At least one option must be provided"
		if e.Msg != wantMsg {
			t.Errorf("expected error message %q, got %q", wantMsg, e.Msg)
		}
	})
}
