package mifer

import (
	"context"
)

type RandomData interface{}

var _ MiferGenerator = (*DefaultMiferGenerator)(nil)

// return generate random type data
type PrepareDataCallBack func() RandomData

type MiferGenerator interface {
	Do(ctx context.Context, targetClmName string, numData int64, fn PrepareDataCallBack) ([]RandomData, error)
}

type DefaultMiferGenerator struct{}

func (dmg *DefaultMiferGenerator) Do(ctx context.Context, targetClmName string, numData int64, fn PrepareDataCallBack) ([]RandomData, error) {
	datum := make([]RandomData, 0, numData)
	var i int64
	for i = 0; i < numData; i++ {
		randomData := fn()
		datum = append(datum, randomData)
	}
	return datum, nil
}

func DefaultStrPrepareDataCallBack() RandomData   { return "default" }
func DefaultIntPrepareDataCallBack() RandomData   { return 0 }
func DefaultEmailPrepareDataCallBack() RandomData { return "default" }
func DefaultUUIDPrepareDataCallBack() RandomData  { return "default" }
