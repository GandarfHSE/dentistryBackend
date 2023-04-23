package tables

// TODO: FOREIGN KEY ("uid") REFERENCES "users" ("id")
var createPatientInfoTable = `
	CREATE TABLE "patients" (
		"id" SERIAL PRIMARY KEY,
		"uid" INT NOT NULL,
		"name" VARCHAR(200) NOT NULL DEFAULT 'anon',
		"pass" VARCHAR(200) NOT NULL DEFAULT ''
	);
`
