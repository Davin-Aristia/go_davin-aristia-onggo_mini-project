package repository

import (
	"go-mini-project/model"

	"gorm.io/gorm"
)

type VendorRepository interface {
	Get(name string) ([]model.Vendor, error)
	GetById(id int) (model.Vendor, error)
	Create(data model.Vendor) (model.Vendor, error)
	Update(data model.Vendor, ID int) error
	Delete(id int) error
}

type vendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *vendorRepository {
	return &vendorRepository{db}
}

func (r *vendorRepository) Get(name string) ([]model.Vendor, error) {
	var vendorData []model.Vendor

	tx := r.db

    if name != "" {
        tx = tx.Where("name LIKE ?", "%"+name+"%")
    }

	tx.Find(&vendorData)
	if tx.Error != nil {
		return []model.Vendor{}, tx.Error
	}
	return vendorData, nil
}

func (r *vendorRepository) GetById(id int) (model.Vendor, error) {
	var vendorData model.Vendor

	tx := r.db.First(&vendorData, id)

	if tx.Error != nil {
		return model.Vendor{}, tx.Error
	}
	return vendorData, nil
}

func (r *vendorRepository) Create(data model.Vendor) (model.Vendor, error) {
	err := model.ValidateVendorRequest(&data)
	if err != nil {
		return model.Vendor{}, err
	}

	tx := r.db.Save(&data)
	if tx.Error != nil {
		return model.Vendor{}, tx.Error
	}
	return data, nil
}

func (r *vendorRepository) Update(data model.Vendor, id int) error {
	tx := r.db.Model(&data).Where("id = ?", id).Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *vendorRepository) Delete(id int) error {
	tx := r.db.Delete(&model.Vendor{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}