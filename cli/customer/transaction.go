package customer

import (
	"dealer_golang/utils"
	"fmt"
	"strconv"
	"strings"
)

func StartTransactionFlow(userID int) {
	fmt.Println("\n===== BELI KENDARAAN =====")

	vehicles, err := VehicleService.GetAllVehicles()
	if err != nil {
		fmt.Println("Gagal mengambil data kendaraan:", err)
		return
	}

	if len(vehicles) == 0 {
		fmt.Println("Tidak ada kendaraan tersedia.")
		return
	}

	fmt.Println("\n=== LIST KENDARAAN ===")
	fmt.Printf("%-4s %-12s %-30s %-15s %-18s %-6s\n",
		"ID", "Brand", "Name", "Type", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 95))

	for _, v := range vehicles {
		fmt.Printf("%-4d %-12s %-30s %-15s %-18s %-6d\n",
			v.ID,
			v.Brand,
			v.Name,
			v.VehicleType,
			utils.FormatRupiah(v.Price),
			v.Stock,
		)
	}

	fmt.Print("\nMasukkan Vehicle ID: ")
	idStr := utils.ReadInput()

	vehicleID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID kendaraan tidak valid.")
		return
	}

	fmt.Print("Masukkan jumlah kendaraan: ")
	qtyStr := utils.ReadInput()

	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		fmt.Println("Jumlah kendaraan tidak valid.")
		return
	}

	fmt.Println("\nPilih Metode Pembayaran:")
	fmt.Println("1. Cash")
	fmt.Println("2. Transfer")
	fmt.Println("3. Credit")
	fmt.Print("Pilih: ")

	method := utils.ReadInput()

	var paymentMethod string
	switch method {
	case "1":
		paymentMethod = "cash"
	case "2":
		paymentMethod = "transfer"
	case "3":
		paymentMethod = "credit"
	default:
		fmt.Println("Metode tidak valid.")
		return
	}

	err = TransactionService.StartTransaction(
		userID,
		vehicleID,
		qty,
		paymentMethod,
	)

	if err != nil {
		fmt.Println("\nTransaksi gagal:", err)
		return
	}

	fmt.Println("\nTransaksi berhasil!")
}
