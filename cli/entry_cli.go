package cli

import (
	"fmt"
	"dealer_golang/service"
	"dealer_golang/utils"

	"dealer_golang/cli/admin"
	"dealer_golang/cli/customer"
)

func ShowMainMenu(
	userService *service.UserService,
	vehicleService *service.VehicleService,
	transactionService *service.TransactionService,
	favoriteService *service.FavoriteService,
) {

	// Inject service ke package customer
	customer.VehicleService = vehicleService
	customer.TransactionService = transactionService
	customer.FavoriteService = favoriteService

	// Inject service ke package admin
	admin.VehicleService = vehicleService
	admin.TransactionService = transactionService
	admin.UserService = userService

	for {
		fmt.Println("===================================")
		fmt.Println("   DEALER OTOMOTIF CLI SYSTEM      ")
		fmt.Println("===================================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		choice := utils.ReadInput()

		switch choice {
		case "1":
			LoginMenu(userService)
		case "2":
			RegisterMenu(userService)
		case "3":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func RegisterMenu(userService *service.UserService) {
	fmt.Println("\n=== REGISTER ===")
	fmt.Print("Enter name: ")
	name := utils.ReadInput()

	fmt.Print("Enter email: ")
	email := utils.ReadInput()

	fmt.Print("Enter password: ")
	password := utils.ReadPassword()

	err := userService.Register(name, email, password)
	if err != nil {
		fmt.Println("Registration failed:", err)
		return
	}

	fmt.Println("Registration successful!")
}

func LoginMenu(userService *service.UserService) {
	fmt.Println("\n=== LOGIN ===")
	fmt.Print("Enter email: ")
	email := utils.ReadInput()

	fmt.Print("Enter password: ")
	password := utils.ReadPassword()

	user, err := userService.Login(email, password)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	fmt.Println("\nLogin successful! Welcome,", user.Name)

	if user.Role == "admin" {
		admin.ShowAdminMenu(user)
	} else {
		customer.ShowCustomerMenu(user.ID)
	}
}
