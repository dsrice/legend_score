package response

type LoginResponse struct {
	Result bool   `json:"result"`
	Token  string `json:"token"`
}