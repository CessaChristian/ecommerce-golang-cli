package customer

import (
	"fmt"
	"strings"
	"strconv"

	"dealer_golang/utils"
)

func ShowWishlistMenu(userID int) {
	fmt.Println("\n===== Wishlist Menu =====")

	list, err := FavoriteService.GetFavorites(userID)
	if err != nil {
		fmt.Println("Gagal mengambil data wishlist:", err)
		return
	}

	if len(list) == 0 {
		fmt.Println("Belum ada kendaraan di wishlist.")
	} else {
		fmt.Println("Daftar Wishlist Kamu:")
		fmt.Printf("%-4s %-12s %-30s %-15s %-18s\n",
			"ID", "Brand", "Name", "Type", "Price")
		fmt.Println(strings.Repeat("-", 90))

		for _, v := range list {
			fmt.Printf("%-4d %-12s %-30s %-15s %-18s\n",
				v.ID,
				v.Brand,
				v.Name,
				v.VehicleType,
				utils.FormatRupiah(v.Price),
			)
		}
	}

	fmt.Println("----------------------------------------------------------")
	fmt.Println("1. Tambah Wishlist")
	fmt.Println("2. Hapus Wishlist")
	fmt.Println("3. Kembali")
	fmt.Print("Pilih opsi: ")

	choice := utils.ReadInput()

	switch choice {
	case "1":
		addWishlist(userID)

	case "2":
		removeWishlist(userID)

	case "3":
		return

	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func addWishlist(userID int) {
	fmt.Println("\n===== Tambah Wishlist =====")

	vehicles, err := VehicleService.GetAllVehicles()
	if err != nil {
		fmt.Println("Gagal mengambil data kendaraan:", err)
		return
	}

	fmt.Printf("%-4s %-12s %-30s %-15s %-18s\n",
		"ID", "Brand", "Name", "Type", "Price")
	fmt.Println(strings.Repeat("-", 90))

	for _, v := range vehicles {
		fmt.Printf("%-4d %-12s %-30s %-15s %-18s\n",
			v.ID,
			v.Brand,
			v.Name,
			v.VehicleType,
			utils.FormatRupiah(v.Price),
		)
	}

	fmt.Print("Masukkan Vehicle ID untuk ditambahkan: ")
	idStr := utils.ReadInput()

	vehicleID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID kendaraan tidak valid.")
		return
	}

	err = FavoriteService.AddFavorite(userID, vehicleID)
	if err != nil {
		fmt.Println("Gagal menambahkan ke wishlist:", err)
		return
	}

	fmt.Println("Berhasil menambahkan ke wishlist!")
}

func removeWishlist(userID int) {
	fmt.Println("\n===== Hapus Wishlist =====")

	list, err := FavoriteService.GetFavorites(userID)
	if err != nil {
		fmt.Println("Gagal mengambil wishlist:", err)
		return
	}

	if len(list) == 0 {
		fmt.Println("Belum ada kendaraan di wishlist.")
		return
	}

	fmt.Printf("%-4s %-12s %-30s %-15s %-18s\n",
		"ID", "Brand", "Name", "Type", "Price")
	fmt.Println(strings.Repeat("-", 90))

	for _, v := range list {
		fmt.Printf("%-4d %-12s %-30s %-15s %-18s\n",
			v.ID,
			v.Brand,
			v.Name,
			v.VehicleType,
			utils.FormatRupiah(v.Price),
		)
	}

	fmt.Print("Masukkan ID kendaraan yang ingin dihapus: ")
	idStr := utils.ReadInput()

	vehicleID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID kendaraan tidak valid.")
		return
	}

	err = FavoriteService.RemoveFavorite(userID, vehicleID)
	if err != nil {
		fmt.Println("Gagal menghapus wishlist:", err)
		return
	}

	fmt.Println("Berhasil menghapus wishlist!")
}
