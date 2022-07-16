package main

import (
	"log"

	"github.com/asadbek280604/server_on_golang_example/pkg/handler"
	"github.com/spf13/viper"
)

func main() {
	RunServer()
}

func RunServer() {
	if err := InitConfig(); err != nil {
		log.Fatalf("ooops: %s", err.Error())
	}

	router := handler.GetRouter()
	_ = router.Run("localhost:" + viper.GetString("port"))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
