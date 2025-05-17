package response

// LoginResponse represents the login response payload
type LoginResponse struct {
	Result bool   `json:"result" example:"true" description:"Indicates if the login was successful"`
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." description:"JWT token for authentication"`
	Code   string `json:"code" example:"" description:"Error code if login failed"`
}