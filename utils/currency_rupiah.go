package utils

import (
	"fmt"
	"strings"
)

// FormatRupiah mengubah float menjadi format Rupiah (Rp 100.000.000)
func FormatRupiah(value float64) string {
	// ubah ke integer (karena harga mobil/motor umumnya tanpa sen)
	n := int64(value)

	// format dasar tanpa titik
	s := fmt.Sprintf("%d", n)

	// tambahkan titik pemisah ribuan manual
	var result strings.Builder
	length := len(s)
	counter := 0

	for i := length - 1; i >= 0; i-- {
		result.WriteByte(s[i])
		counter++
		if counter%3 == 0 && i != 0 {
			result.WriteByte('.')
		}
	}

	// reverse hasilnya
	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return "Rp " + string(runes)
}
