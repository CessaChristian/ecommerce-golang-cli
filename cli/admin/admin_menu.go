package admin

import (
	"fmt"
	"dealer_golang/models"
	"dealer_golang/service"
	"dealer_golang/utils"
)

var VehicleService *service.VehicleService
var TransactionService *service.TransactionService
var UserService *service.UserService

func ShowAdminMenu(user models.User) {
	for {
		fmt.Println("\n========== ADMIN MENU ==========")
		fmt.Printf("Logged in as: %s (Admin)\n", user.Name)
		fmt.Println("1. Manage Vehicles")
		fmt.Println("2. Reports")
		fmt.Println("3. Logout")
		fmt.Print("Choose an option: ")

		choice := utils.ReadInput()

		switch choice {
		case "1":
			ShowVehicleMenu(VehicleService)

		case "2":
			ShowReportMenu()

		case "3":
			fmt.Println("Logging out...")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
