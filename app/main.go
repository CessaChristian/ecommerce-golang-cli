package main

import (
	"log"

	"dealer_golang/config"
	"dealer_golang/cli"
	"dealer_golang/repositories"
	"dealer_golang/service"
	"dealer_golang/utils"

	"github.com/joho/godotenv"
)

func main() {

	// handle CTRL + C
	utils.HandleCtrlC()

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// connect to database
	config.Connect()


	// INIT REPOSITORY
	userRepo := repositories.NewUserRepository(config.DB)
	vehicleRepo := repositories.NewVehicleRepository(config.DB)
	transactionRepo := repositories.NewTransactionRepository(config.DB)
	paymentRepo := repositories.NewPaymentRepository(config.DB)
	favoriteRepo := repositories.NewFavoriteRepository(config.DB)
	typeRepo := repositories.NewTypeRepository(config.DB)
	brandRepo := repositories.NewBrandRepository(config.DB)


	// INIT SERVICE
	userService := service.NewUserService(userRepo)
	vehicleService := service.NewVehicleService(vehicleRepo, typeRepo, brandRepo)
	paymentService := service.NewPaymentService()
	favoriteService := service.NewFavoriteService(favoriteRepo, vehicleRepo)

	transactionService := service.NewTransactionService(
		vehicleRepo,
		transactionRepo,
		paymentRepo,
		paymentService,
	)

	// RUN CLI ENTRY MENU
	cli.ShowMainMenu(
		userService,
		vehicleService,
		transactionService,
		favoriteService,
	)



}
