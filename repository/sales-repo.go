package repository

import (
	// "go-mini-project/dto"
	"go-mini-project/model"

	"gorm.io/gorm"
	"time"
	"fmt"
    "errors"
    "strconv"
)

type SalesRepository interface {
	Get(invoice string, user int) ([]model.Sales, error)
	GetById(id int) (model.Sales, error)
	
	GenerateNextInvoice() (string, error)

    BeginTransaction() *gorm.DB
    CreateWithTransaction(sales model.Sales, tx *gorm.DB) (model.Sales, error)
    CreateSalesDetailWithTransaction(salesDetail model.SalesDetail, tx *gorm.DB) error
    DeductBookStockWithTransaction(bookID uint, quantity int, tx *gorm.DB) error
	UpdateTotalWithTransaction(ID uint, total float64, tx *gorm.DB) error
    CommitTransaction(tx *gorm.DB) error
}

type salesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) *salesRepository {
	return &salesRepository{db}
}

func (r *salesRepository) Get(invoice string, user int) ([]model.Sales, error) {
	var salesData []model.Sales

	tx := r.db
	
	if invoice != "" {
        tx = tx.Where("invoice LIKE ?", "%"+invoice+"%")
    }

    if user != 0 {
        tx = tx.Where("user_id = ?", user)
    }

	tx.Preload("SalesDetails").Find(&salesData)
	if tx.Error != nil {
		return []model.Sales{}, tx.Error
	}
	return salesData, nil
}

func (r *salesRepository) GetById(id int) (model.Sales, error) {
	var salesData model.Sales

	tx := r.db.Preload("SalesDetails").First(&salesData, id)

	if tx.Error != nil {
		return model.Sales{}, tx.Error
	}
	return salesData, nil
}

func (r *salesRepository) GenerateNextInvoice() (string, error) {
    now := time.Now()
    year, month, day := now.Year(), now.Month(), now.Day()

    formattedInvoice := fmt.Sprintf("INV/%d/%02d/%02d/", year, month, day)

    var lastInvoice model.Sales
    if err := r.db.Where("invoice LIKE ?", formattedInvoice+"%").Order("invoice DESC").First(&lastInvoice).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            formattedInvoice += "0001"
        } else {
            return "", err
        }
    } else {
        var sequence int
        _, err := fmt.Sscanf(lastInvoice.Invoice, formattedInvoice+"%04d", &sequence)
        if err != nil {
            return "", err
        }

        formattedInvoice += fmt.Sprintf("%04d", sequence+1)
    }

    return formattedInvoice, nil
}

func (r *salesRepository) BeginTransaction() *gorm.DB {
    return r.db.Begin()
}

func (r *salesRepository) CreateWithTransaction(sales model.Sales, tx *gorm.DB) (model.Sales, error) {
	err := model.ValidateSalesRequest(&sales)
	if err != nil {
		return model.Sales{}, err
	}
	
    if err := tx.Create(&sales).Error; err != nil {
        return model.Sales{}, err
    }
    return sales, nil
}

func (r *salesRepository) CreateSalesDetailWithTransaction(salesDetail model.SalesDetail, tx *gorm.DB) error {
	err := model.ValidateSalesDetailRequest(&salesDetail)
	if err != nil {
		return err
	}

    if err := tx.Create(&salesDetail).Error; err != nil {
        return err
    }
    return nil
}

func (r *salesRepository) DeductBookStockWithTransaction(bookID uint, quantity int, tx *gorm.DB) error {
	var book model.Book
	if err := tx.First(&book, bookID).Error; err != nil {
		return err
	}

	if book.Stock < quantity {
		return errors.New("Insufficient stock for '" + book.Title + "'. Remaining Stock: " + strconv.Itoa(book.Stock) + ". Quantity requested: "+ strconv.Itoa(quantity))
	}

	book.Stock -= quantity

	if err := tx.Save(&book).Error; err != nil {
		return err
	}

	return nil
}

func (r *salesRepository) UpdateTotalWithTransaction(ID uint, total float64, tx *gorm.DB) error {
	var sales model.Sales

	if err:= tx.Model(&sales).Where("id = ?", ID).Updates(map[string]interface{}{"total":total}). Error; err != nil{
		return tx.Error
	}
	return nil
}

func (r *salesRepository) CommitTransaction(tx *gorm.DB) error {
    if err := tx.Commit().Error; err != nil {
        return err
    }
    return nil
}

