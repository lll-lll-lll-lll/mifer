package mifer

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
)

type RandomData interface{}

var _ MiferGenerator = (*DefaultMiferGenerator)(nil)

// return generate random type data
// the data generated can be customized.
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

func DefaultUserNamePrepareDataCallBack() RandomData   { return gofakeit.Username() }
func DefaultInt64PrepareDataCallBack() RandomData      { return gofakeit.Int64() }
func DefaultEmailPrepareDataCallBack() RandomData      { return gofakeit.Email() }
func DefaultUUIDPrepareDataCallBack() RandomData       { return gofakeit.UUID() }
func DefaultDATEStringPrepareDataCallBack() RandomData { return gofakeit.Date().String() }
func DefaultLanguagePrepareDataCallBack() RandomData   { return gofakeit.Language() }
