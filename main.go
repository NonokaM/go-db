package main

import (
	"log"
	"io"
	"net/http"

	"github.com/NonokaM/Go-API/handlers"
)

func main() {
	// Register to use handler
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.HelloHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	// Start Server(use ListenAndServe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
