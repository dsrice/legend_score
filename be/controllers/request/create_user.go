package request

// CreateUserRequest represents the create user request payload
type CreateUserRequest struct {
	Name     string `json:"name" example:"John Doe" description:"User's full name"`
	Password string `json:"password" example:"password123" description:"User's password"`
	LoginID  string `json:"login_id" example:"john.doe" description:"User's login ID"`
}