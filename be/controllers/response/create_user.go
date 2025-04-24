package response

// CreateUserResponse represents the create user response payload
type CreateUserResponse struct {
	Result bool   `json:"result" example:"true" description:"Indicates if the user creation was successful"`
	Code   string `json:"code" example:"" description:"Error code if user creation failed"`
}