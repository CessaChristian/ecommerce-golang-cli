package admin

import (
	"fmt"
	"dealer_golang/service"
	"dealer_golang/utils"
)

func ShowVehicleMenu(vehicleService *service.VehicleService) {
	for {
		fmt.Println("\n===== VEHICLE MANAGEMENT =====")
		fmt.Println("1. Add Vehicle")
		fmt.Println("2. Add Brand and Type")
		fmt.Println("3. List All Vehicles")
		fmt.Println("4. Get Vehicle by ID")
		fmt.Println("5. Update Vehicle Stock")
		fmt.Println("6. Back")
		fmt.Print("Choose an option: ")

		choice := utils.ReadInput()

		switch choice {
		case "1":
			addVehicle(vehicleService)
		case "2":
			ShowMasterDataMenu()
		case "3":
			listVehicles(vehicleService)
		case "4":
			getVehicleByID(vehicleService)
		case "5":
			updateStock(vehicleService)
		case "6":
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

// ADD VEHICLE
func addVehicle(vehicleService *service.VehicleService) {
	fmt.Println("\n===== add vehicle =====")

	// input type pakai nama
	fmt.Print("Enter vehicle Type Name: ")
	typeName := utils.ReadInput()

	vt, err := vehicleService.FindTypeByName(typeName)
	if err != nil {
		fmt.Println("Type tidak ditemukan:", typeName)
		return
	}

	// input brand pakai nama
	fmt.Print("Enter brand name: ")
	brandName := utils.ReadInput()

	br, err := vehicleService.FindBrandByName(brandName)
	if err != nil {
		fmt.Println("brand tidak ditemukan:", brandName)
		return
	}

	// input data lainnya
	fmt.Print("Enter vehicle name: ")
	name := utils.ReadInput()

	fmt.Print("Enter fuel type: ")
	fuel := utils.ReadInput()

	fmt.Print("Enter transmission: ")
	trans := utils.ReadInput()

	fmt.Print("Enter price: ")
	price := utils.ReadFloat()

	fmt.Print("Enter stock: ")
	stock := utils.ReadInt()

	// create
	err = vehicleService.CreateVehicle(vt.ID, br.ID, name, fuel, trans, price, stock)
	if err != nil {
		fmt.Println("Failed to add vehicle:", err)
		return
	}

	fmt.Print("\nVehicle added successfully!\n")
}

// LIST VEHICLES
func listVehicles(vehicleService *service.VehicleService) {
	vehicles, err := vehicleService.GetAllVehicles()
	if err != nil {
		fmt.Println("Failed to get vehicles:", err)
		return
	}

	fmt.Println("\n===== ALL VEHICLES =====")

	if len(vehicles) == 0 {
		fmt.Println("No vehicles found.")
		return
	}

	for _, v := range vehicles {
		fmt.Printf("[%d] %s %s | Type: %s | Fuel: %s | Trans: %s | Price: %s | Stock: %d\n",
			v.ID,
			v.Brand,
			v.Name,
			v.VehicleType,
			v.FuelType,
			v.Transmission,
			utils.FormatRupiah(v.Price),
			v.Stock,
		)
	}
}

// GET VEHICLE BY ID
func getVehicleByID(vehicleService *service.VehicleService) {
	fmt.Print("Enter Vehicle ID: ")
	id := utils.ReadInt() // ← FIX

	v, err := vehicleService.GetVehicleByID(id)
	if err != nil {
		fmt.Println("Vehicle not found:", err)
		return
	}

	fmt.Println("\n===== VEHICLE DETAIL =====")
	fmt.Printf("ID      : %d\n", v.ID)
	fmt.Printf("Type    : %s\n", v.VehicleType)
	fmt.Printf("Brand   : %s\n", v.Brand)
	fmt.Printf("Name    : %s\n", v.Name)
	fmt.Printf("Fuel    : %s\n", v.FuelType)
	fmt.Printf("Trans   : %s\n", v.Transmission)
	fmt.Printf("Price   : %s\n", utils.FormatRupiah(v.Price))
	fmt.Printf("Stock   : %d\n", v.Stock)
}

// UPDATE STOCK
func updateStock(vehicleService *service.VehicleService) {
	fmt.Print("Enter Vehicle ID: ")
	id := utils.ReadInt()

	v, err := vehicleService.GetVehicleByID(id)
	if err != nil {
		fmt.Println("Vehicle not found:", err)
		return
	}

	fmt.Printf("Current stock of %s: %d\n", v.Name, v.Stock)

	fmt.Print("New Stock: ")
	newStock := utils.ReadInt()

	err = vehicleService.UpdateStock(id, newStock)
	if err != nil {
		fmt.Println("Failed to update stock:", err)
		return
	}

	fmt.Println("Stock updated successfully!")
}
