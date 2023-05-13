package tables

var createServiceTable = `
	CREATE TABLE "services" (
		"id" SERIAL PRIMARY KEY,
		"name" VARCHAR(100) NOT NULL,
		"description" VARCHAR(300) NOT NULL DEFAULT '',
		"cost" INT NOT NULL,
		"duration" INT NOT NULL
	);
`

var createServiceLinkTable = `
	CREATE TABLE "service_links" (
		"id" SERIAL PRIMARY KEY,
		"did" INT NOT NULL,
		"sid" INT NOT NULL
	);
`
