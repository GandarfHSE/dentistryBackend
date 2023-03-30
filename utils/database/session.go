package database

import (
	"context"

	"github.com/ansel1/merry"
	pgx "github.com/jackc/pgx/v5"
)

type Session struct {
	connector *pgx.Conn
	Config    *pgx.ConnConfig
	Ctx       context.Context
	tx        pgx.Tx
}

func (s *Session) Open() (err error) {
	s.connector, err = pgx.ConnectConfig(s.Ctx, s.Config)
	return err
}

func (s *Session) OpenTx(txOptions pgx.TxOptions) (err error) {
	if s.tx != nil {
		return merry.Errorf("Session::OpenTx - tx != nil")
	}

	s.tx, err = s.connector.BeginTx(s.Ctx, txOptions)
	return err
}

func (s *Session) RollbackTx() (err error) {
	if s.tx == nil {
		return merry.Errorf("Session::Rollback - tx == nil")
	}

	err = s.tx.Rollback(s.Ctx)
	s.tx = nil
	return err
}

func (s *Session) CloseTx() (err error) {
	if s.tx == nil {
		return merry.Errorf("Session::Close - tx == nil")
	}

	err = s.tx.Commit(s.Ctx)
	s.tx = nil
	return err
}

func (s *Session) Close() (err error) {
	if s.tx != nil {
		err = s.CloseTx()

		if err != nil {
			return err
		}
	}

	err = s.connector.Close(s.Ctx)
	return err
}
