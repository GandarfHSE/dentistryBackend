package tables

var createReviewTable = `
	CREATE TABLE "reviews" (
		"id" SERIAL PRIMARY KEY,
		"cid" INT NOT NULL DEFAULT 0,
		"did" INT NOT NULL DEFAULT 0,
		"sid" INT NOT NULL DEFAULT 0,
		"score" INT NOT NULL DEFAULT 0,
		"description" VARCHAR(1000) NOT NULL DEFAULT ''
	);
`
