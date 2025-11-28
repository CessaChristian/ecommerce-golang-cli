package admin

import (
    "fmt"
    "dealer_golang/utils"
)

// MASTER DATA MENU
func ShowMasterDataMenu() {
    for {
        fmt.Println("\n===== MASTER DATA MANAGEMENT =====")
        fmt.Println("1. Manage Vehicle Types")
        fmt.Println("2. Manage Brands")
        fmt.Println("3. Back")
        fmt.Print("Choose an option: ")
        choice := utils.ReadInput()

        switch choice {
        case "1":
            manageTypes()
        case "2":
            manageBrands()
        case "3":
            return
        default:
            fmt.Println("Invalid option, try again.")
        }
    }
}

// VEHICLE TYPE MENU
func manageTypes() {
    for {
        fmt.Println("\n===== VEHICLE TYPE MANAGEMENT =====")
        fmt.Println("1. Add Type")
        fmt.Println("2. List Types")
        fmt.Println("3. Back")
        fmt.Print("Choose an option: ")

        choice := utils.ReadInput()

        switch choice {
        case "1":
            addType()
        case "2":
            listTypes()
        case "3":
            return
        default:
            fmt.Println("Invalid option.")
        }
    }
}

func addType() {
    fmt.Print("Enter new vehicle type: ")
    name := utils.ReadInput()

    if name == "" {
        fmt.Println("Type name tidak boleh kosong.")
        return
    }

    err := VehicleService.RepoType.CreateType(name)
    if err != nil {
        fmt.Println("Failed:", err)
        return
    }

    fmt.Println("Vehicle type added successfully!")
}


func listTypes() {
    types, err := VehicleService.GetAllVehicleTypes()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(types) == 0 {
        fmt.Println("No vehicle types found.")
        return
    }

    fmt.Println("\n+------+----------------------+")
    fmt.Println("| ID   | Vehicle Type         |")
    fmt.Println("+------+----------------------+")
    for _, t := range types {
        fmt.Printf("| %-4d | %-20s |\n", t.ID, t.TypeName)
    }
    fmt.Println("+------+----------------------+")
}

// BRAND MENU
func manageBrands() {
    for {
        fmt.Println("\n===== BRAND MANAGEMENT =====")
        fmt.Println("1. Add Brand")
        fmt.Println("2. List Brands")
        fmt.Println("3. Back")
        fmt.Print("Choose an option: ")

        choice := utils.ReadInput()

        switch choice {
        case "1":
            addBrand()
        case "2":
            listBrands()
        case "3":
            return
        default:
            fmt.Println("Invalid option.")
        }
    }
}

func addBrand() {
    fmt.Print("Enter new brand name: ")
    name := utils.ReadInput()

    if name == "" {
        fmt.Println("Brand name tidak boleh kosong.")
        return
    }

    err := VehicleService.RepoBrand.CreateBrand(name)
    if err != nil {
        fmt.Println("Failed:", err)
        return
    }

    fmt.Println("Brand added successfully!")
}


func listBrands() {
    brands, err := VehicleService.GetAllBrands()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(brands) == 0 {
        fmt.Println("No brands found.")
        return
    }

    fmt.Println("\n+------+----------------------+")
    fmt.Println("| ID   | Brand Name           |")
    fmt.Println("+------+----------------------+")
    for _, b := range brands {
        fmt.Printf("| %-4d | %-20s |\n", b.ID, b.BrandName)
    }
    fmt.Println("+------+----------------------+")
}

