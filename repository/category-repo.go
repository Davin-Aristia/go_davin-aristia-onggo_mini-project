package repository

import (
	"go-mini-project/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Get(name string) ([]model.Category, error)
	GetById(id int) (model.Category, error)
	Create(data model.Category) (model.Category, error)
	Update(data model.Category, ID int) error
	Delete(id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Get(name string) ([]model.Category, error) {
	var categoryData []model.Category

	tx := r.db

    if name != "" {
        tx = tx.Where("name LIKE ?", "%"+name+"%")
    }

	tx.Preload("Books").Find(&categoryData)
	if tx.Error != nil {
		return []model.Category{}, tx.Error
	}
	return categoryData, nil
}

func (r *categoryRepository) GetById(id int) (model.Category, error) {
	var categoryData model.Category

	tx := r.db.Preload("Books").First(&categoryData, id)

	if tx.Error != nil {
		return model.Category{}, tx.Error
	}
	return categoryData, nil
}

func (r *categoryRepository) Create(data model.Category) (model.Category, error) {
	err := model.ValidateCategoryRequest(&data)
	if err != nil {
		return model.Category{}, err
	}

	tx := r.db.Save(&data)
	if tx.Error != nil {
		return model.Category{}, tx.Error
	}
	return data, nil
}

func (r *categoryRepository) Update(data model.Category, id int) error {
	tx := r.db.Model(&data).Where("id = ?", id).Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *categoryRepository) Delete(id int) error {
	tx := r.db.Delete(&model.Category{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}