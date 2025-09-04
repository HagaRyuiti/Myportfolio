package services

import (
	"Myportfolio/repositories"
	"fmt"
)

// ログイン処理のビジネスロジックを担うサービス層
func LoginService(email, password string) (string, error) {
	// ユーザーリポジトリからユーザー情報を取得
	user, exists := repositories.Users[email]
	if !exists || user.Password != password {
		return "", fmt.Errorf("認証情報が正しくありません")
	}

	// 認証成功。ここでは簡単なトークンを返す
	// 本番環境ではJWTなどの認証トークンを生成する
	token := "dummy-auth-token"
	return token, nil
}
