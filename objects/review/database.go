package review

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
)

func addReview(s *database.Session, req CreateReviewRequest) error {
	q := `
		INSERT INTO "reviews" (cid, did, sid, score, description)
		VALUES ($1, $2, $3, $4, $5);
	`

	err := database.Modify(s, q, req.Cid, req.Did, req.Sid, req.Score, req.Description)
	return err
}

func getReviewList(s *database.Session) ([]Review, error) {
	q := `
		SELECT * FROM "reviews";
	`

	rs, err := database.Get[Review](s, q)
	return rs, err
}

func findReviewBy(s *database.Session, column string, value int) ([]Review, error) {
	q := `
		SELECT * FROM "reviews"
		WHERE "%s" = %d;
	`

	rs, err := database.Get[Review](s, fmt.Sprintf(q, column, value))
	return rs, err
}
