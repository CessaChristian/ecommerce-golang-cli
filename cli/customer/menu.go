package customer

import (
	"fmt"
	"strings"

	"dealer_golang/utils"
	"dealer_golang/service"
)

func ShowCustomerMenu(userID int) {
	for {
		fmt.Println("\n===== CUSTOMER MENU =====")
		fmt.Println("1. Lihat Kendaraan")
		fmt.Println("2. Beli Kendaraan")
		fmt.Println("3. Wishlist")
		fmt.Println("4. Lihat Riwayat Transaksi")
		fmt.Println("5. Logout")

		fmt.Print("Choose an option: ")
		choice := utils.ReadInput()

		switch choice {
		case "1":
			ShowVehicleListForUser(VehicleService)

		case "2":
			StartTransactionFlow(userID)

		case "3":
			ShowWishlistMenu(userID)

		case "4":
			ShowTransactionHistory(userID)

		case "5":
			fmt.Println("Logging out...")
			return

		default:
			fmt.Println("Invalid option!")
		}
	}
}

func ShowVehicleListForUser(svc *service.VehicleService) {
	vehicles, err := svc.GetAllVehicles()
	if err != nil {
		fmt.Println("Failed to get vehicles:", err)
		return
	}

	fmt.Println("\n===== LIST SEMUA KENDARAAN =====")

	if len(vehicles) == 0 {
		fmt.Println("No vehicles found.")
		return
	}

	// Header tabel
	fmt.Printf("%-4s %-12s %-30s %-15s %-18s %-5s\n",
		"ID", "Brand", "Name", "Type", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 95))

	for _, v := range vehicles {
		fmt.Printf("%-4d %-12s %-30s %-15s %-18s %-5d\n",
			v.ID,
			v.Brand,
			v.Name,
			v.VehicleType,
			utils.FormatRupiah(v.Price),
			v.Stock,
		)
	}
}
