package tables

var createDoctorInfoTable = `
	CREATE TABLE "doctor" (
		"id" SERIAL PRIMARY KEY,
		"uid" INT NOT NULL,
		"name" VARCHAR(200) NOT NULL DEFAULT 'anon',
		"post" VARCHAR(200) NOT NULL DEFAULT '',
		"exp" INT NOT NULL DEFAULT 0,
		FOREIGN KEY ("uid") REFERENCES "users" ("id")
	);
`
