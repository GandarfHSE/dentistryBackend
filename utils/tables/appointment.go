package tables

// TODO: foreign keys
var createAppointmentTable = `
	CREATE TABLE "appointments" (
		"id" SERIAL PRIMARY KEY,
		"pid" INT NOT NULL,
		"did" INT NOT NULL,
		"sid" INT NOT NULL,
		"timebegin" TIMESTAMP NOT NULL,
		"timeend" TIMESTAMP NOT NULL
	);
`
