package mifer_test

import (
	"context"
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
		{ColumnKey: "id", Datum: []mifer.RandomData{[]int{1, 2}}},
		{ColumnKey: "name", Datum: []mifer.RandomData{[]string{"test1", "test2"}}}}
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
	clmns := mifer.Columns{"id": mifer.Column{ColumnType: "int"}, "name": mifer.Column{ColumnType: "nvarchar"}}
	queries := psql.BuildQueries(context.Background(), 10, clmns,
		mifer.MiferOption{"id", []mifer.RandomData{mifer.DefaultIntPrepareDataCallBack()}},
		mifer.MiferOption{"name", []mifer.RandomData{mifer.DefaultStrPrepareDataCallBack()}})
	want := []string{"INSERT INTO users (id, name) VALUES (0, 'default');"}
	t.Log(queries[0])
	assert.Equal(t, queries[0], want[0])
}
