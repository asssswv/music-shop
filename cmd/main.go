package main

import "github.com/asadbek280604/server_on_golang_example/pkg/handler"

func main() {
	RunServer()
}

func RunServer() {
	router := handler.GetRouter()
	_ = router.Run("localhost:8080")
}