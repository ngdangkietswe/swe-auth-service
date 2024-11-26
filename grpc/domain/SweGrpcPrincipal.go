package domain

type SweGrpcPrincipal struct {
	UserId   string   `json:"user_id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Token    string   `json:"token"`
	Roles    []string `json:"roles"`
}
