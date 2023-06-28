package mifer_test

import (
	"context"
	"testing"

	"github.com/lll-lll-lll-lll/mifer"
	"gotest.tools/assert"
)

func Test_Do(t *testing.T) {
	t.Parallel()
	gen := &mifer.DefaultMiferGenerator{}
	wantNum := 100
	data, err := gen.Do(context.Background(), int64(wantNum), func() interface{} { return mifer.DefaultInt64PrepareDataCallBack() })
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(data), wantNum)

}
