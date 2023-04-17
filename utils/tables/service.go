package tables

var createServiceTable = `
	CREATE TABLE "services" (
		"id" SERIAL PRIMARY KEY,
		"name" VARCHAR(100) NOT NULL,
		"desc" VARCHAR(300) NOT NULL DEFAULT '',
		"cost" INT NOT NULL,
		"duration" INT NOT NULL
	);
`
