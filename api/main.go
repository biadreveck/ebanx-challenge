package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"ebanx/challenge/api/middleware"
	_apiRouter "ebanx/challenge/api/router"
	"ebanx/challenge/data/repository"
	"ebanx/challenge/domain/usecase"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())

	// Repositories
	accountRepo := repository.NewAccountInMemoryRepository()

	// Usecases
	eventUcase := usecase.NewEventUsecase(accountRepo)
	balanceUcase := usecase.NewBalanceUsecase(accountRepo)
	resetUcase := usecase.NewResetUsecase(accountRepo)

	// Routers
	_apiRouter.NewEventRouter(router, eventUcase)
	_apiRouter.NewBalanceRouter(router, balanceUcase)
	_apiRouter.NewResetRouter(router, resetUcase)

	router.Run(viper.GetString("server.address"))
}
