package types

type Session struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	TokenID     string `json:"token_id"`
	TokenSecret string `json:"token_secret"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
