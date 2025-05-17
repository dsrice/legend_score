package request

// GetUsersRequest represents the get users request payload with filtering options
type GetUsersRequest struct {
    UserID  *int    `query:"user_id" description:"Filter by user ID"`
    LoginID *string `query:"login_id" description:"Filter by login ID"`
    Name    *string `query:"name" description:"Filter by user name"`
}