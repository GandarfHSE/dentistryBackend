package tables

var createUserTable = `
	CREATE TABLE "users" (
		"id" SERIAL PRIMARY KEY,
		"login" VARCHAR(42) NOT NULL,
		"password" VARCHAR(100) NOT NULL,
		"role" INT NOT NULL
	);
`
