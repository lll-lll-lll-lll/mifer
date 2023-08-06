package mifer_test

import (
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
	got := mifer.JoinClmnKeys(ipt)
	assert.Equal(t, got, want)
}
