package service

import (
	"dealer_golang/models"
	"dealer_golang/repositories"
	"errors"
)

type FavoriteService struct {
	favoriteRepo *repositories.FavoriteRepository
	vehicleRepo *repositories.VehicleRepository
}


func NewFavoriteService(favRepo *repositories.FavoriteRepository, vehicleRepo *repositories.VehicleRepository) *FavoriteService {
	return &FavoriteService{
		favoriteRepo: favRepo,
		vehicleRepo:  vehicleRepo,
	}
}

// tambahkan ke favorite
func(s *FavoriteService) AddFavorite(userID int, vehicleID int) error {

	// cek kendaraan
	_, err := s.vehicleRepo.GetByID(vehicleID)
	if err != nil {
		return errors.New("kendaraan tidak ditemukan")
	}

	// tambahkan ke favorite
	return s.favoriteRepo.Add(userID, vehicleID)
}

// lihat daftar favorite
func (s *FavoriteService) GetFavorites(userID int) ([]models.Vehicle, error) {
	return s.favoriteRepo.GetByUser(userID)
}

// hapus favorite
func (s *FavoriteService) RemoveFavorite(userID int, vehicleID int) error {
	return s.favoriteRepo.Remove(userID, vehicleID)
}