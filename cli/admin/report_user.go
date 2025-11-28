package admin

import (
    "fmt"
    "dealer_golang/service"
    "dealer_golang/utils"
)

func ShowAllUsersReport(userService *service.UserService) {
    fmt.Println("\n===== ALL USERS REPORT =====")

    list, err := userService.GetAllUsers()
    if err != nil {
        fmt.Println("Gagal mengambil data user:", err)
        return
    }

    if len(list) == 0 {
        fmt.Println("Belum ada user terdaftar.")
        return
    }

    fmt.Printf("%-5s %-20s %-25s %-10s %-20s\n",
        "ID", "Name", "Email", "Role", "Created At")
    fmt.Println("-------------------------------------------------------------------------------")

    for _, u := range list {
        fmt.Printf("%-5d %-20s %-25s %-10s %-20s\n",
            u.ID,
            u.Name,
            u.Email,
            u.Role,
            utils.FormatWIB(u.CreatedAt),
        )
    }
}
