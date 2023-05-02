package tables

var createClinicTable = `
	CREATE TABLE "clinics" (
		"id" SERIAL PRIMARY KEY,
		"name" VARCHAR(300) NOT NULL,
		"address" VARCHAR(300) NOT NULL,
		"phone" VARCHAR(20) NOT NULL DEFAULT ''
	);
`
