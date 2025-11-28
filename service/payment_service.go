package service

import (
	"dealer_golang/utils"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

// main entry
func (p *PaymentService) ProcessPayment(method string, total float64) (bool, string, error) {

	switch method {
	case "cash":
		return p.payCash(total)
	case "transfer":
		return p.payTransfer(total)
	case "credit":
		return p.payCredit(total)
	default:
		return false, "", errors.New("invalid payment method")
	}
}

// cash payment
func (p *PaymentService) payCash(total float64) (bool, string, error) {
	fmt.Printf("\nTotal harus dibayar: Rp %s\n", utils.FormatRupiah(total))

	for {
		fmt.Print("Masukkan jumlah uang: Rp ")
		uangStr := utils.ReadInput()

		var uang float64
		fmt.Sscanf(uangStr, "%f", &uang)

		if uang < total {
			fmt.Println("Uang tidak cukup. Harap masukkan jumlah uang yang sesuai.")
			continue
		}

		kembalian := uang - total
		fmt.Printf("Pembayaran Berhasil! Kembalian: %s\n", utils.FormatRupiah(kembalian))

		note := fmt.Sprintf("Cash: terima %s, kembalian %s", utils.FormatRupiah(uang), utils.FormatRupiah(kembalian))
		return true, note, nil
	}
}

// transfer payment
func (p *PaymentService) payTransfer(total float64) (bool, string, error) {
	kode := generateCode()
	fmt.Println("\n=== PEMBAYARAN TRANSFER ===")
	fmt.Println("Silakan transfer ke rekening BCA 1234567890")
	fmt.Println("Kode pembayaran Anda:", kode)

	fmt.Print("Masukkan kode pembayaran untuk verifikasi: ")
	input := utils.ReadInput()

	if input != kode {
		fmt.Println("Kode salah. Pembayaran gagal.")
		return false, "", nil
	}

	fmt.Println("Pembayaran berhasil diverifikasi!")
	note := fmt.Sprintf("Transfer: kode %s", kode)
	return true, note, nil
}

// credit payment
func (p *PaymentService) payCredit(total float64) (bool, string, error) {
	minDP := total * 0.20
	fmt.Printf("\n=== Pembayaran Kredit ===\n")
	fmt.Printf("Total: Rp %s\n", utils.FormatRupiah(total))
	fmt.Printf("Minimal DP: Rp %s\n", utils.FormatRupiah(minDP))

	var dp float64

	for {
		fmt.Print("Masukkan DP: Rp ")
		dpStr := utils.ReadInput()
		fmt.Sscanf(dpStr, "%f", &dp)

		if dp < minDP {
			fmt.Println("DP tidak mencukupi. Harap masukkan jumlah DP yang sesuai.")
			continue
		}
		break
	}

	fmt.Println("\nPilih Tenor Cicilan:")
	fmt.Println("1. 12 bulan")
	fmt.Println("2. 24 bulan")
	fmt.Println("3. 36 bulan")
	fmt.Print("Pilih tenor: ")

	tenor := utils.ReadInput()
	var bulan int
	switch tenor {
	case "1":
		bulan = 12
	case "2":
		bulan = 24
	case "3":
		bulan = 36
	default:
		return false, " ", errors.New("tenor tidak valid")
	}

	cicilanbulan := (total - dp) / float64(bulan)

	fmt.Printf("Cicilan per bulan: Rp %s selama %d bulan\n", utils.FormatRupiah(cicilanbulan), bulan)
	fmt.Println("Credit disetujui.")

	note := fmt.Sprintf("Credit: DP %s, tenor %d bulan, cicilan %s/bulan", utils.FormatRupiah(dp), bulan, utils.FormatRupiah(cicilanbulan))

	return true, note, nil
}

// generate random code
func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%05d", rand.Intn(99999))
}
