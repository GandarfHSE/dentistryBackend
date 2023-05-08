package review

import (
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

// [TODO: #32] Check values validity
func CreateReviewHandler(req CreateReviewRequest, _ *cookie.Cookie) (*CreateReviewResponse, merry.Error) {
	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateReviewHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	err = addReview(s, req)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &CreateReviewResponse{Err: "-"}, nil
}

func GetReviewListHandler(req GetReviewListRequest, _ *cookie.Cookie) (*ReviewListResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetReviewListHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	reviews, err := getReviewList(s)
	if err != nil {
		log.Error().Err(err).Msg("Can't get review list!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &ReviewListResponse{ReviewList: reviews}, nil
}

func FindClinicReviewHandler(req FindClinicReviewRequest, _ *cookie.Cookie) (*ReviewListResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at FindClinicReviewHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	reviews, err := findReviewBy(s, "cid", req.Cid)
	if err != nil {
		log.Error().Err(err).Msg("Can't find reviews by clinic id!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &ReviewListResponse{ReviewList: reviews}, nil
}

func FindDoctorReviewHandler(req FindDoctorReviewRequest, _ *cookie.Cookie) (*ReviewListResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at FindDoctorReviewHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	reviews, err := findReviewBy(s, "did", req.Did)
	if err != nil {
		log.Error().Err(err).Msg("Can't find reviews by doctor id!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &ReviewListResponse{ReviewList: reviews}, nil
}

func FindServiceReviewHandler(req FindServiceReviewRequest, _ *cookie.Cookie) (*ReviewListResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at FindServiceReviewHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	reviews, err := findReviewBy(s, "sid", req.Sid)
	if err != nil {
		log.Error().Err(err).Msg("Can't find reviews by service id!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &ReviewListResponse{ReviewList: reviews}, nil
}
