package request

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	LoginID  string `json:"login_id"`
}