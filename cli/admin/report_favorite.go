package admin

import (
    "fmt"
    "dealer_golang/utils"
)

func ShowFavoriteReport() {
    fmt.Println("\n===== MOST FAVORITED VEHICLES =====")

    list, err := VehicleService.GetMostFavoritedVehicles()
    if err != nil {
        fmt.Println("failed to load report:", err)
        return
    }

    if len(list) == 0 {
        fmt.Println("no data found")
        return
    }

    // Header
    fmt.Println("+------+----------------------+----------------------+------------------------------------+------------+---------------+")
	fmt.Println("| ID   | Brand                | Type                 | Name                               | Favorites  | Price         |")
	fmt.Println("+------+----------------------+----------------------+------------------------------------+------------+---------------+")


    // Rows
    for _, v := range list {
		fmt.Printf(
			"| %-4d | %-20s | %-20s | %-34s | %-10d | %-13s |\n",
			v.ID,
			v.Brand,
			v.Type,
			v.Name,
			v.TotalFavorites,
			utils.FormatRupiah(v.Price),
		)
	}

	fmt.Println("+------+----------------------+----------------------+------------------------------------+------------+---------------+")
}
