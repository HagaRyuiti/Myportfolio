package main

import (
	"Myportfolio/routes"
	"log"
	"net/http"
)

// サーバーの起動と設定を行う
func main() {
	router := routes.SetupRoutes()

	// サーバーをポート8080で起動
	log.Println("Goサーバーをポート8080で起動しました...")
	log.Println("http://localhost:8080/login でログインAPIにアクセス可能です")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
