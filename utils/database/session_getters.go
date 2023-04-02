package database

import (
	"context"

	"github.com/GandarfHSE/dentistryBackend/utils/config"
	"github.com/jackc/pgx/v5"
)

func getSession(opt pgx.TxOptions) (*Session, error) {
	conf, err := config.GetConnConfig()
	if err != nil {
		return nil, err
	}

	s := Session{
		Ctx:    context.Background(),
		Config: conf,
	}

	err = s.Open()
	if err != nil {
		return nil, err
	}

	err = s.OpenTx(opt)
	return &s, err
}

func GetReadSession() (*Session, error) {
	return getSession(pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadOnly,
	})
}

func GetWriteSession() (*Session, error) {
	return getSession(pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	})
}

func GetScanSession() (*Session, error) {
	return getSession(pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadOnly,
	})
}
