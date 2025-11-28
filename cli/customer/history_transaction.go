package customer

import (
	"fmt"
	"strconv"

	"dealer_golang/service"
	"dealer_golang/utils"
)

// Menampilkan list transaksi terlebih dahulu
func ShowTransactionHistory(userID int) {
	fmt.Println("\n===== RIWAYAT TRANSAKSI =====")

	// ambil list transaksi milik user
	list, err := TransactionService.GetTransactionHistory(userID)
	if err != nil {
		fmt.Println("Gagal mengambil data transaksi:", err)
		return
	}

	if len(list) == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}

	// tampilkan daftar transaksi (WIB)
	fmt.Println("\nDaftar Transaksi:")
	for i, t := range list {
		fmt.Printf("%d. Transaction #%d | Total: %s | Tanggal: %s\n",
			i+1,
			t.ID,
			utils.FormatWIB(t.CreatedAt), 
			utils.FormatRupiah(t.TotalAmount),
		)
	}

	fmt.Print("\nPilih nomor transaksi untuk melihat detail: ")

	choice := utils.ReadInput()
	idx, err := strconv.Atoi(choice)

	if err != nil || idx < 1 || idx > len(list) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	selected := list[idx-1]

	// tampilkan detail
	showTransactionDetail(selected.ID, TransactionService)
}

// Menampilkan detail transaksi pilihan user
func showTransactionDetail(transactionID int, transactionService *service.TransactionService) {

	header, items, payment, err := transactionService.GetTransactionDetail(transactionID)
	if err != nil {
		fmt.Println("Gagal mengambil detail transaksi:", err)
		return
	}

	fmt.Println("\n===================================")
	fmt.Println("      DETAIL TRANSAKSI")
	fmt.Println("===================================")
	fmt.Printf("Transaction ID : %d\n", header.ID)
	fmt.Printf("Tanggal        : %s\n", utils.FormatWIB(header.CreatedAt)) // ← WIB
	fmt.Printf("Total          : %s\n", utils.FormatRupiah(header.TotalAmount))
	fmt.Println("-----------------------------------")

	// ITEMS
	fmt.Println("Items:")
	for _, it := range items {
		fmt.Printf("- %s %s (%s) x%d @ %s\n",
			it.Brand,
			it.Name,
			it.VehicleType,
			it.Quantity,
			utils.FormatRupiah(it.Price),
		)
	}

	fmt.Println("-----------------------------------")
	// PAYMENT
	fmt.Println("PAYMENT DETAIL:")
	fmt.Printf("Metode : %s\n", payment.PaymentMethod)
	fmt.Printf("Status : %s\n", payment.Status)

	// PaidAt pointer → WIB
	fmt.Printf("Paid At : %s\n", utils.FormatWIBPtr(payment.PaidAt))

	// Note
	if payment.Note != nil {
		fmt.Printf("Catatan : %s\n", *payment.Note)
	}

	fmt.Print("\n===================================\n")

}
