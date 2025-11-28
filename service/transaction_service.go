package service

import (
    "errors"
    "fmt"
    "time"

    "dealer_golang/models"
    "dealer_golang/repositories"
    "dealer_golang/utils"
)

type TransactionService struct {
    VehicleRepo     *repositories.VehicleRepository
    TransactionRepo *repositories.TransactionRepository
    PaymentRepo     *repositories.PaymentRepository
    PaymentService  *PaymentService
}

func NewTransactionService(
    vehicleRepo *repositories.VehicleRepository,
    transactionRepo *repositories.TransactionRepository,
    paymentRepo *repositories.PaymentRepository,
    paymentService *PaymentService,
) *TransactionService {

    return &TransactionService{
        VehicleRepo:     vehicleRepo,
        TransactionRepo: transactionRepo,
        PaymentRepo:     paymentRepo,
        PaymentService:  paymentService,
    }
}

// START TRANSACTION PROCESS
func (s *TransactionService) StartTransaction(userID int, vehicleID int, quantity int, paymentMethod string) error {

    // 1. Ambil kendaraan
    vehicle, err := s.VehicleRepo.GetByID(vehicleID)
    if err != nil {
        return err
    }

    // 2. Cek stok
    if vehicle.Stock < quantity {
        return errors.New("stok kendaraan tidak cukup")
    }

    // 3. Hitung total
    total := vehicle.Price * float64(quantity)

    // 4. Proses pembayaran sesuai metode
    paid, note, err := s.PaymentService.ProcessPayment(paymentMethod, total)
    if err != nil {
        return err
    }

    if !paid {
        return errors.New("pembayaran gagal atau dibatalkan")
    }

    // 5. Buat transaksi header
    t := models.Transaction{
        UserID:      userID,
        TotalAmount: total,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    transactionID, err := s.TransactionRepo.Create(t)
    if err != nil {
        return err
    }

    // 6. Tambah item transaksi
    item := models.TransactionItem{
        TransactionID: transactionID,
        VehicleID:     vehicleID,
        Quantity:      quantity,
        Price:         vehicle.Price,
    }

    err = s.TransactionRepo.AddItem(item)
    if err != nil {
        return err
    }

    // 7. Simpan payment_detail
	now := time.Now().UTC()
    payment := models.PaymentDetail{
        TransactionID: transactionID,
        PaymentMethod: paymentMethod,
        Status:        "paid",
        PaidAt:        &now,
        Note:          &note, // catatan tambahan seperti kode transfer atau DP
    }

    err = s.PaymentRepo.Create(payment)
    if err != nil {
        return err
    }

    // 8. Update stok kendaraan
    newStock := vehicle.Stock - quantity
    err = s.VehicleRepo.UpdateStock(vehicleID, newStock)
    if err != nil {
        return err
    }

    fmt.Println("\n=== TRANSACTION SUCCESS ===")
    fmt.Println("Vehicle:", vehicle.Name)
    fmt.Println("Quantity:", quantity)
    fmt.Printf("Total: Rp %s\n", utils.FormatRupiah(total))
    fmt.Println("Payment Method:", paymentMethod)
    fmt.Println("Note:", note)
    fmt.Println("===========================")

    return nil
}

// TRANSACTION HISTORY
func (s *TransactionService) GetTransactionHistory(userID int) ([]models.Transaction, error) {
    return s.TransactionRepo.GetByUser(userID)
}

func (s *TransactionService) GetTransactionDetail(transactionID int) (
    models.Transaction,
    []models.TransactionItemWithVehicle,
    models.PaymentDetail,
    error,
) {
    // ambil header
    header, err := s.TransactionRepo.GetTransactionByID(transactionID)
    if err != nil {
        return header, nil, models.PaymentDetail{}, err
    }

    // ambil items
    items, err := s.TransactionRepo.GetTransactionItems(transactionID)
    if err != nil {
        return header, nil, models.PaymentDetail{}, err
    }

    // ambil payment detail
    payment, err := s.PaymentRepo.GetPaymentDetail(transactionID)
    if err != nil {
        return header, items, models.PaymentDetail{}, err
    }

    return header, items, payment, nil
}

func (s *TransactionService) GetAllTransactions() ([]models.TransactionWithPayment, error) {
    return s.TransactionRepo.GetAllTransactions()
}


