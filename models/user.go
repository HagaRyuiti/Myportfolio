package models

// ユーザー情報を定義する構造体
type User struct {
	Email    string
	Password string
}

// ログインリクエストのJSONを扱うための構造体
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ログインレスポンスのJSONを扱うための構造体
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
}
