package database

import (
	"errors"

	"github.com/rs/zerolog/log"
)

func Modify(s *Session, q string, args ...any) error {
	if s == nil || s.tx == nil {
		return errors.New("Modify: session transaction does not exist!")
	}

	tag, err := s.tx.Exec(s.Ctx, q, args...)
	log.Info().Msg(tag.String())
	return err
}
