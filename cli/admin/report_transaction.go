package admin

import (
    "fmt"
    "dealer_golang/service"
    "dealer_golang/utils"
)

func ShowAllTransactionsReport(transactionService *service.TransactionService) {
    fmt.Println("\n===== ALL TRANSACTIONS REPORT =====")

    list, err := transactionService.GetAllTransactions()
    if err != nil {
        fmt.Println("Gagal mengambil laporan transaksi:", err)
        return
    }

    if len(list) == 0 {
        fmt.Println("Belum ada transaksi.")
        return
    }

    // Header tabel (dengan spacing lebih besar & rapi)
    fmt.Printf("%-5s %-15s %-20s %-12s %-22s %-10s\n",
        "ID", "Customer", "Amount", "Method", "Date", "Status")
    fmt.Println("--------------------------------------------------------------------------------------")

    for _, t := range list {
        fmt.Printf("%-5d %-15s %-20s %-12s %-22s %-10s\n",
            t.ID,
            t.UserName,
            utils.FormatRupiah(t.TotalAmount),
            t.PaymentMethod,
            utils.FormatWIB(t.CreatedAt),
            t.Status,
        )
    }
}
