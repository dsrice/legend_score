package request

type LoginRequest struct {
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}