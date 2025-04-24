package request

// LoginRequest represents the login request payload
type LoginRequest struct {
	LoginID  string `json:"login_id" example:"user123" description:"User login ID"`
	Password string `json:"password" example:"password123" description:"User password"`
}