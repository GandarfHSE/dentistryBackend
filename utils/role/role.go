package role

// golang doesn't have enum ¯\_(ツ)_/¯
const (
	Invalid int = 0
	Patient int = 1
	Doctor  int = 2
	Admin   int = 4
	Dev     int = 8
)

func IsRoleValid(role int) bool {
	return role == Patient || role == Doctor || role == Admin || role == Dev
}

func IsRoleAtLeast(role1 int, role2 int) bool {
	return role1 >= role2
}
