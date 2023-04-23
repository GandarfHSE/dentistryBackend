package user

// golang doesn't have enum ¯\_(ツ)_/¯
const (
	PatientRole int = 1
	DoctorRole  int = 2
	AdminRole   int = 4
	DevRole     int = 8
)

func IsRoleValid(role int) bool {
	return role == PatientRole || role == DoctorRole || role == AdminRole || role == DevRole
}

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type CreateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type CreateUserResponse struct {
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JWT string `json:"jwt"`
}

type GetUserListRequest struct {
}

type GetUserListResponse struct {
	UserList []User `json:"userlist"`
}

type WhoAmIRequest struct {
}

type WhoAmIResponse struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Role  int    `json:"role"`
}
