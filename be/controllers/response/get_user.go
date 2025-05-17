package response

// GetUserResponse represents the get user response payload
type GetUserResponse struct {
	Result bool         `json:"result" example:"true" description:"Indicates if the operation was successful"`
	Code   string       `json:"code" example:"" description:"Error code if operation failed"`
	User   UserResponse `json:"user" description:"User details"`
}