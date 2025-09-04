package repositories

import (
	"Myportfolio/models"
)

// ユーザー情報をインメモリで管理するデータベース層
var Users = map[string]models.User{
	"user@example.com": {
		Email:    "user@example.com",
		Password: "password123", // パスワードは本番環境ではハッシュ化すること
	},
}
