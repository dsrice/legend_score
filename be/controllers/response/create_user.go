package response

type CreateUserResponse struct {
	Result bool   `json:"result"`
	Code   string `json:"code"`
}