package tables

var createUserTable = `--sql
	CREATE TABLE "users" (
		"id" INT PRIMARY KEY,
		"login" VARCHAR(42) NOT NULL,
		"password" VARCHAR(100) NOT NULL,
		"role" INT NOT NULL
	);
`
