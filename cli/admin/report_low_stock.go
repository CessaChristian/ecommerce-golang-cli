package admin

import (
    "fmt"
    "dealer_golang/service"
    "dealer_golang/utils"
)

func ShowLowStockReport(vehicleService *service.VehicleService) {
    fmt.Println("\n===== LOW STOCK VEHICLES REPORT =====")

    vehicles, err := vehicleService.LowStockReport()
    if err != nil {
        fmt.Println("Gagal mengambil laporan:", err)
        return
    }

    if len(vehicles) == 0 {
        fmt.Println("Tidak ada kendaraan dengan stok rendah.")
        return
    }

    fmt.Print("\nKendaraan dengan stok ≤ 3 unit:\n")

    // Header tabel
    fmt.Printf("%-5s %-15s %-20s %-10s %-15s %-6s\n",
        "ID", "Brand", "Name", "Type", "Price", "Stock")
    fmt.Println("--------------------------------------------------------------------------")

    // Data rows
    for _, v := range vehicles {
        fmt.Printf("%-5d %-15s %-20s %-10s %-15s %-6d\n",
            v.ID,
            v.Brand,
            v.Name,
            v.VehicleType,
            utils.FormatRupiah(v.Price),
            v.Stock,
        )
    }
}
