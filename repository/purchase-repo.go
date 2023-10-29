package repository

import (
	"go-mini-project/model"

	"gorm.io/gorm"
	"time"
	"fmt"
)

type PurchaseRepository interface {
	Get(purchaseOrder string, vendor int) ([]model.Purchase, error)
	GetById(id int) (model.Purchase, error)
	
	GenerateNextPurchaseOrder() (string, error)

    BeginTransaction() *gorm.DB
    CreateWithTransaction(purchase model.Purchase, tx *gorm.DB) (model.Purchase, error)
    CreatePurchaseDetailWithTransaction(purchaseDetail model.PurchaseDetail, tx *gorm.DB) error
    AddBookStockWithTransaction(bookID uint, quantity int, tx *gorm.DB) error
	UpdateTotalWithTransaction(ID uint, total float64, tx *gorm.DB) error
    CommitTransaction(tx *gorm.DB) error
}

type purchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *purchaseRepository {
	return &purchaseRepository{db}
}

func (r *purchaseRepository) Get(purchaseOrder string, vendor int) ([]model.Purchase, error) {
	var purchaseData []model.Purchase

	tx := r.db
	
	if purchaseOrder != "" {
        tx = tx.Where("purchase_order LIKE ?", "%"+purchaseOrder+"%")
    }

    if vendor != 0 {
        tx = tx.Where("vendor_id = ?", vendor)
    }

	tx.Preload("PurchaseDetails").Find(&purchaseData)
	if tx.Error != nil {
		return []model.Purchase{}, tx.Error
	}
	return purchaseData, nil
}

func (r *purchaseRepository) GetById(id int) (model.Purchase, error) {
	var purchaseData model.Purchase

	tx := r.db.Preload("PurchaseDetails").First(&purchaseData, id)

	if tx.Error != nil {
		return model.Purchase{}, tx.Error
	}
	return purchaseData, nil
}

func (r *purchaseRepository) GenerateNextPurchaseOrder() (string, error) {
    now := time.Now()
    year, month, day := now.Year(), now.Month(), now.Day()

    formattedPurchaseOrder := fmt.Sprintf("PUR/%d/%02d/%02d/", year, month, day)

    var lastPurchaseOrder model.Purchase
    if err := r.db.Where("purchase_order LIKE ?", formattedPurchaseOrder+"%").Order("purchase_order DESC").First(&lastPurchaseOrder).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            formattedPurchaseOrder += "0001"
        } else {
            return "", err
        }
    } else {
        var sequence int
        _, err := fmt.Sscanf(lastPurchaseOrder.PurchaseOrder, formattedPurchaseOrder+"%04d", &sequence)
        if err != nil {
            return "", err
        }

        formattedPurchaseOrder += fmt.Sprintf("%04d", sequence+1)
    }

    return formattedPurchaseOrder, nil
}

func (r *purchaseRepository) BeginTransaction() *gorm.DB {
    return r.db.Begin()
}

func (r *purchaseRepository) CreateWithTransaction(purchase model.Purchase, tx *gorm.DB) (model.Purchase, error) {
	err := model.ValidatePurchaseRequest(&purchase)
	if err != nil {
		return model.Purchase{}, err
	}
	
    if err := tx.Create(&purchase).Error; err != nil {
        return model.Purchase{}, err
    }
    return purchase, nil
}

func (r *purchaseRepository) CreatePurchaseDetailWithTransaction(purchaseDetail model.PurchaseDetail, tx *gorm.DB) error {
	err := model.ValidatePurchaseDetailRequest(&purchaseDetail)
	if err != nil {
		return err
	}

    if err := tx.Create(&purchaseDetail).Error; err != nil {
        return err
    }
    return nil
}

func (r *purchaseRepository) AddBookStockWithTransaction(bookID uint, quantity int, tx *gorm.DB) error {
	var book model.Book
	if err := tx.First(&book, bookID).Error; err != nil {
		return err
	}

	book.Stock += quantity

	if err := tx.Save(&book).Error; err != nil {
		return err
	}

	return nil
}

func (r *purchaseRepository) UpdateTotalWithTransaction(ID uint, total float64, tx *gorm.DB) error {
	var purchase model.Purchase

	if err:= tx.Model(&purchase).Where("id = ?", ID).Updates(map[string]interface{}{"total":total}). Error; err != nil{
		return tx.Error
	}
	return nil
}

func (r *purchaseRepository) CommitTransaction(tx *gorm.DB) error {
    if err := tx.Commit().Error; err != nil {
        return err
    }
    return nil
}

