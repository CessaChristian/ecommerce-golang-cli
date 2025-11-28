package admin


import (
	"fmt"
	"dealer_golang/utils"
)

func ShowReportMenu() {
	for {
		fmt.Println("\n===== REPORT MENU =====")
		fmt.Println("1. Low Stock Vehicles")
		fmt.Println("2. Semua Transaksi")
		fmt.Println("3. Semua User")
		fmt.Println("4. Kendaraan Favorite")
		fmt.Println("5. Kembali")
		fmt.Print("Choose an option: ")

		choice := utils.ReadInput()

		switch choice {

		case "1":
			ShowLowStockReport(VehicleService)

		case "2":
			ShowAllTransactionsReport(TransactionService)

		case "3":
			ShowAllUsersReport(UserService)
		case "4":
			ShowFavoriteReport()
		case "5":
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}
