package review

type Review struct {
	Id          int    `json:"id"`
	Cid         int    `json:"cid"`
	Did         int    `json:"did"`
	Sid         int    `json:"sid"`
	Score       int    `json:"score"`
	Description string `json:"description"`
}

type CreateReviewRequest struct {
	Cid         int    `json:"cid"`
	Did         int    `json:"did"`
	Sid         int    `json:"sid"`
	Score       int    `json:"score"`
	Description string `json:"description"`
}

// check README: empty json in response
type CreateReviewResponse struct {
	Err string `json:"err"`
}
