package main

import (
	"log"
	"net/http"

	"github.com/moriwakihikari/clean_architecture_with_todo.git/infrastructure"
)

// データの流れ
// DB操作
// Use case(Interactor)→interface(Repository)→infrastructure(Sqlhandler)

// リクエスト処理
// アクセス→infrastructure(handler)→interfaces(controller)→httpリスポンス

func main() {
	// 最外殻router.goへ
	mux := infrastructure.SetUpRouting()
    log.Fatal(http.ListenAndServe(":8080", mux))
}