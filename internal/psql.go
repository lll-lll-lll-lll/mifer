package internal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lll-lll-lll-lll/mifer"
)

func injectToOption(ctx context.Context, option *mifer.MiferOption, numData int64, fn mifer.PrepareDataCallBack) error {
	defaultGen := mifer.DefaultMiferGenerator{}
	datum, err := defaultGen.Do(ctx, numData, fn)
	if err != nil {
		return err
	}
	option.Datum = datum
	return nil
}

func inferColumnType(ctx context.Context, db *sql.DB, tableName string, columnName string) (mifer.PrepareDataCallBack, error) {
	rows, err := db.QueryContext(ctx, fmt.Sprintf(mifer.ScanQuery, tableName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var clmType string
		if err := rows.Scan(name, clmType); err != nil {
			return nil, err
		}
		if columnName == name {
			fn := switchInferType(clmType)
			return fn, nil
		}
	}

	return nil, fmt.Errorf("unable to infer type of column")
}

func switchInferType(clmType string) mifer.PrepareDataCallBack {
	switch clmType {
	case "nvarchar":
		return mifer.DefaultStringPrepareDataCallBack
	case "varchar":
		return mifer.DefaultStringPrepareDataCallBack
	case "text":
		return mifer.DefaultStringPrepareDataCallBack
	case "uuid":
		return mifer.DefaultUUIDPrepareDataCallBack
	case "int":
		return mifer.DefaultInt64PrepareDataCallBack
	default:
		return mifer.NilPrepareDataCallBack
	}
}
