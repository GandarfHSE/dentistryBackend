package tables

// TODO: FOREIGN KEY ("uid") REFERENCES "users" ("id")
var createDoctorInfoTable = `
	CREATE TABLE "doctors" (
		"id" SERIAL PRIMARY KEY,
		"uid" INT NOT NULL,
		"name" VARCHAR(200) NOT NULL DEFAULT 'anon',
		"post" VARCHAR(200) NOT NULL DEFAULT '',
		"exp" INT NOT NULL DEFAULT 0
	);
`
