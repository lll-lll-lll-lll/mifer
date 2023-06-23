package mifer

type ErrType string

const (
	MigrationErr ErrType = "migration error"
	SqlErr               = "sql error"
	DBErr                = "database error"
	InjectErr            = "inject error"
	MapKeyErr            = "map key error"
	NillErr              = "nill error"
	GenErr               = "random data generation error"
)

type MiferError struct {
	ErrType ErrType
	Msg     string
	err     error
}

func (me *MiferError) Error() string {
	return string(me.ErrType) + " " + me.Msg
}

func (me *MiferError) Unwrap() error {
	return me.err
}

func Error(errType ErrType, msg string) *MiferError {
	return &MiferError{
		Msg:     msg,
		ErrType: errType,
	}
}
