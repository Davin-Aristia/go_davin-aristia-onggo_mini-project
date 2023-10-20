package repository

import (
	// "go-mini-project/dto"
	"go-mini-project/model"

	"gorm.io/gorm"
	"time"
	"fmt"
)

type SalesRepository interface {
	Get(invoice string, user int) ([]model.Sales, error)
	GetById(id int) (model.Sales, error)
	Create(data model.Sales) (model.Sales, error)
	GenerateNextInvoice() (string, error)

	// BeginTransaction starts a database transaction and returns a reference to it.
    BeginTransaction() *gorm.DB

    // CreateWithTransaction creates a new sales record within the given transaction.
    CreateWithTransaction(sales model.Sales, tx *gorm.DB) (model.Sales, error)

    // CreateSalesDetailWithTransaction creates a new sales detail record within the given transaction.
    CreateSalesDetailWithTransaction(salesDetail model.SalesDetail, tx *gorm.DB) error

    // CommitTransaction commits the provided transaction.
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

	tx := r.db.First(&salesData, id)

	if tx.Error != nil {
		return model.Sales{}, tx.Error
	}
	return salesData, nil
}

func (r *salesRepository) Create(data model.Sales) (model.Sales, error) {
	err := model.ValidateSalesRequest(&data)
	if err != nil {
		return model.Sales{}, err
	}

	tx := r.db.Save(&data)
	if tx.Error != nil {
		return model.Sales{}, tx.Error
	}
	return data, nil
}

func (r *salesRepository) GenerateNextInvoice() (string, error) {
    // Get the current date components
    now := time.Now()
    year, month, day := now.Year(), now.Month(), now.Day()

    // Construct the desired format
    formattedInvoice := fmt.Sprintf("INV/%d/%02d/%02d/", year, month, day)

    // Find the last invoice with the same date prefix
    var lastInvoice model.Sales
    if err := r.db.Where("invoice LIKE ?", formattedInvoice+"%").Order("invoice DESC").First(&lastInvoice).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // If no invoice found, start with 0001
            formattedInvoice += "0001"
        } else {
            return "", err
        }
    } else {
        // Extract the sequence from the last invoice
        var sequence int
        _, err := fmt.Sscanf(lastInvoice.Invoice, formattedInvoice+"%04d", &sequence)
        if err != nil {
            return "", err
        }

        // Increment the sequence for the new invoice
        formattedInvoice += fmt.Sprintf("%04d", sequence+1)
    }

    return formattedInvoice, nil
}

func (r *salesRepository) BeginTransaction() *gorm.DB {
    return r.db.Begin()
}

// CreateWithTransaction creates a new sales record within the given transaction.
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

// CreateSalesDetailWithTransaction creates a new sales detail record within the given transaction.
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

// CommitTransaction commits the provided transaction.
func (r *salesRepository) CommitTransaction(tx *gorm.DB) error {
    if err := tx.Commit().Error; err != nil {
        return err
    }
    return nil
}
