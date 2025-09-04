package handlers

import (
	"Myportfolio/models"
	"Myportfolio/services"
	"Myportfolio/utils"
	"encoding/json"
	"net/http"
)

// HTTPリクエストのハンドリングを行う層
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// CORS設定
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Preflightリクエストの処理
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// POSTリクエストのみ許可
	if r.Method != http.MethodPost {
		utils.WriteJSONResponse(w, http.StatusMethodNotAllowed, models.LoginResponse{Error: "許可されていないメソッドです"})
		return
	}

	var req models.LoginRequest
	// リクエストボディをJSONとしてデコード
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.LoginResponse{Error: "無効なリクエストボディです"})
		return
	}

	// サービス層に処理を委譲
	token, err := services.LoginService(req.Email, req.Password)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusUnauthorized, models.LoginResponse{Error: err.Error()})
		return
	}

	// 成功レスポンスを返す
	res := models.LoginResponse{
		Message: "ログイン成功",
		Token:   token,
	}
	utils.WriteJSONResponse(w, http.StatusOK, res)
}
