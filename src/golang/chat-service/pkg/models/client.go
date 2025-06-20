package models

type ClientInfo struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Online    bool   `json:"online"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
