package repository

import (
	"go-mini-project/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	Get(title, author string, category int) ([]model.Book, error)
	GetById(id int) (model.Book, error)
	Create(data model.Book) (model.Book, error)
	Update(data model.Book, ID int) error
	Delete(id int) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) Get(title, author string, category int) ([]model.Book, error) {
	var bookData []model.Book

	tx := r.db
	
	if title != "" {
        tx = tx.Where("title LIKE ?", "%"+title+"%")
    }

    if author != "" {
        tx = tx.Where("author LIKE ?", "%"+author+"%")
    }

    if category != 0 {
        tx = tx.Where("category_id = ?", category)
    }

	tx.Find(&bookData)
	if tx.Error != nil {
		return []model.Book{}, tx.Error
	}
	return bookData, nil
}

func (r *bookRepository) GetById(id int) (model.Book, error) {
	var bookData model.Book

	tx := r.db.First(&bookData, id)

	if tx.Error != nil {
		return model.Book{}, tx.Error
	}
	return bookData, nil
}

func (r *bookRepository) Create(data model.Book) (model.Book, error) {
	err := model.ValidateBookRequest(&data)
	if err != nil {
		return model.Book{}, err
	}

	tx := r.db.Save(&data)
	if tx.Error != nil {
		return model.Book{}, tx.Error
	}
	return data, nil
}

func (r *bookRepository) Update(data model.Book, id int) error {
	tx := r.db.Model(&data).Where("id = ?", id).Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *bookRepository) Delete(id int) error {
	tx := r.db.Delete(&model.Book{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}