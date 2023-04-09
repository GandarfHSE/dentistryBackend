package database

import (
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func parseModel[Model any](rows pgx.Rows) ([]Model, error) {
	defer rows.Close()

	res := make([]Model, 0)
	for rows.Next() {
		var model Model
		err := pgxscan.ScanRow(&model, rows)
		if err != nil {
			return nil, err
		}
		res = append(res, model)
	}

	return res, nil
}

func Get[Model any](s *Session, q string, args ...any) ([]Model, error) {
	if s == nil || s.tx == nil {
		return nil, errors.New("Get: session transaction does not exist!")
	}

	rows, err := s.tx.Query(s.Ctx, q, args...)
	if err != nil {
		return nil, err
	}

	return parseModel[Model](rows)
}
