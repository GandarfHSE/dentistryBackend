package database

import "errors"

func Modify(s *Session, q string) error {
	if s == nil || s.tx == nil {
		return errors.New("Modify: session transaction does not exist!")
	}

	_, err := s.tx.Exec(s.Ctx, q)
	return err
}
