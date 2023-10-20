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

	tx.Find(&salesData)
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