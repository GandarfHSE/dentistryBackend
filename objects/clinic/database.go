package clinic

import "github.com/GandarfHSE/dentistryBackend/utils/database"

func createClinic(s *database.Session, req CreateClinicRequest) error {
	q := `
		INSERT INTO "clinics" (name, address, phone)
		VALUES ($1, $2, $3);
	`

	err := database.Modify(s, q, req.Name, req.Address, req.Phone)
	return err
}

func getClinicList(s *database.Session) ([]Clinic, error) {
	q := `
		SELECT * FROM "clinics";
	`
	return database.Get[Clinic](s, q)
}
