package controllers

import (
	"errors"

	"github.com/jgcaceres97/goly/app/database"
	"github.com/jgcaceres97/goly/app/model"
	"gorm.io/gorm"
)

func GetAllGolies() ([]model.Goly, error) {
	var golies []model.Goly

	tx := database.DB.Find(&golies)
	if tx.Error != nil {
		return []model.Goly{}, tx.Error
	}

	return golies, nil
}

func GetGoly(id *uint64) (model.Goly, error) {
	var goly model.Goly

	tx := database.DB.Where("id = ?", id).First(&goly)
	if tx.Error != nil {
		return model.Goly{}, tx.Error
	}

	return goly, nil
}

func CheckGoly(slug *string) error {
	var goly model.Goly

	tx := database.DB.Where("goly = ?", slug).First(&goly)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return tx.Error
	}

	if tx.RowsAffected > 0 {
		return errors.New("goly already taken")
	}

	return nil
}

func CreateGoly(goly *model.Goly) error {
	tx := database.DB.Create(&goly)
	return tx.Error
}

func UpdateGoly(goly *model.Goly) error {
	tx := database.DB.Save(&goly)
	return tx.Error
}

func DeleteGoly(id *uint64) error {
	tx := database.DB.Unscoped().Delete(&model.Goly{}, id)
	return tx.Error
}

func FindByGolyUrl(url *string) (model.Goly, error) {
	var goly model.Goly

	tx := database.DB.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}
