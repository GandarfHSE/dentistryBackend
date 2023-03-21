package user

// golang doesn't have enum ¯\_(ツ)_/¯
const (
	ClientRole int = 1
	DoctorRole int = 2
	AdminRole  int = 4
	DevRole    int = 8
)

func IsRoleValid(role int) bool {
	return role == ClientRole || role == DoctorRole || role == AdminRole || role == DevRole
}

type User struct {
	Id       int
	Login    string
	Password string
	Role     int
}

type CreateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type CreateUserResponse struct {
	Id int `json:"id"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponce struct {
	JWT string `json:"jwt"`
}
