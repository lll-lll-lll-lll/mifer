package mifer

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
)

// return generate random type data
// the data generated can be customized.
type PrepareDataCallBack func() interface{}

type DefaultMiferGenerator struct{}

func (dmg DefaultMiferGenerator) Do(ctx context.Context, numData int64, fn PrepareDataCallBack) ([]interface{}, error) {
	datum := make([]interface{}, 0, numData)
	var i int64
	for i = 0; i < numData; i++ {
		randomData := fn()
		datum = append(datum, randomData)
	}
	return datum, nil
}

func DefaultUserNamePrepareDataCallBack() interface{}   { return gofakeit.Username() }
func DefaultStringPrepareDataCallBack() interface{}     { return gofakeit.Letter() }
func DefaultInt64PrepareDataCallBack() interface{}      { return gofakeit.Int64() }
func DefaultEmailPrepareDataCallBack() interface{}      { return gofakeit.Email() }
func DefaultUUIDPrepareDataCallBack() interface{}       { return gofakeit.UUID() }
func DefaultDATEStringPrepareDataCallBack() interface{} { return gofakeit.Date().String() }
func DefaultLanguagePrepareDataCallBack() interface{}   { return gofakeit.Language() }
func NilPrepareDataCallBack() interface{}               { return nil }
