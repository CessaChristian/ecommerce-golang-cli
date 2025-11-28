package service

import (
	"dealer_golang/models"
	"dealer_golang/repositories"
	"errors"
)

type VehicleService struct {
	RepoVehicle *repositories.VehicleRepository
	RepoType    *repositories.TypeRepository
	RepoBrand   *repositories.BrandRepository
}

func NewVehicleService(
	repoVehicle *repositories.VehicleRepository,
	repoType *repositories.TypeRepository,
	repoBrand *repositories.BrandRepository,
) *VehicleService {

	return &VehicleService{
		RepoVehicle: repoVehicle,
		RepoType:    repoType,
		RepoBrand:   repoBrand,
	}
}

// Create Vehicle
func (s *VehicleService) CreateVehicle(
	typeID int,
	brandID int,
	name string,
	fuelType string,
	transmission string,
	price float64,
	stock int,
) error {

	if name == "" {
		return errors.New("name tidak boleh kosong")
	}

	v := models.Vehicle{
		TypeID:       typeID,
		BrandID:      brandID,
		Name:         name,
		FuelType:     fuelType,
		Transmission: transmission,
		Price:        price,
		Stock:        stock,
	}

	return s.RepoVehicle.Create(v)
}

// List All
func (s *VehicleService) GetAllVehicles() ([]models.Vehicle, error) {
	return s.RepoVehicle.GetAll()
}

// Get By ID
func (s *VehicleService) GetVehicleByID(id int) (models.Vehicle, error) {
	return s.RepoVehicle.GetByID(id)
}

// Update stock
func (s *VehicleService) UpdateStock(vehicleID int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock tidak boleh negatif")
	}
	return s.RepoVehicle.UpdateStock(vehicleID, newStock)
}

// Low stock
func (s *VehicleService) LowStockReport() ([]models.Vehicle, error) {
	return s.RepoVehicle.GetLowStock()
}

// GET TYPES
func (s *VehicleService) GetAllVehicleTypes() ([]models.VehicleType, error) {
	return s.RepoType.GetAllTypes()
}

// GET BRANDS
func (s *VehicleService) GetAllBrands() ([]models.Brand, error) {
	return s.RepoBrand.GetAllBrands()
}

func (s *VehicleService) FindTypeByName(name string) (models.VehicleType, error) {
    return s.RepoType.GetTypeByName(name)
}

func (s *VehicleService) FindBrandByName(name string) (models.Brand, error) {
    return s.RepoBrand.GetBrandByName(name)
}

func (s *VehicleService) GetMostFavoritedVehicles() ([]models.VehicleFavoriteReport, error) {
    return s.RepoVehicle.GetMostFavorited()
}

