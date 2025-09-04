package routes

import (
	"Myportfolio/handlers"
	"net/http"
)

// ルーティングを定義する層
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.LoginHandler)
	return mux
}
