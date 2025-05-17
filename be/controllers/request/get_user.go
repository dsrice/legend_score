package request

// GetUserRequest represents the get user request payload
type GetUserRequest struct {
	UserID int `param:"user_id" description:"User ID to retrieve"`
}