package response

// UserResponse represents a single user in the response
type UserResponse struct {
    ID      int    `json:"id" example:"1" description:"User ID"`
    LoginID string `json:"login_id" example:"john.doe" description:"User's login ID"`
    Name    string `json:"name" example:"John Doe" description:"User's full name"`
}

// GetUsersResponse represents the get users response payload
type GetUsersResponse struct {
    Result bool           `json:"result" example:"true" description:"Indicates if the operation was successful"`
    Code   string         `json:"code" example:"" description:"Error code if operation failed"`
    Users  []UserResponse `json:"users" description:"List of users matching the filter criteria"`
}