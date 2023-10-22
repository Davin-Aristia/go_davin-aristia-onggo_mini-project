package usecase

import (
	"errors"
	"strconv"
	"time"

	"go-mini-project/dto" 
	"go-mini-project/model"
	"go-mini-project/repository"
)

type PurchaseUsecase interface {
	Get(purchaseOrder, vendor string) ([]model.Purchase, error)
	GetById(id int) (model.Purchase, error)
	Create(payloads dto.PurchaseRequest) (model.Purchase, error)
}

type purchaseUsecase struct {
	purchaseRepository repository.PurchaseRepository
	bookRepository repository.BookRepository
}

func NewPurchaseUsecase(purchaseRepo repository.PurchaseRepository, bookRepo repository.BookRepository) *purchaseUsecase {
	return &purchaseUsecase{
		purchaseRepository: purchaseRepo,
		bookRepository: bookRepo,
	}
}

func (s *purchaseUsecase) Get(purchaseOrder, vendor string) ([]model.Purchase, error) {
	vendorId := 0
	var err error
	if vendor != ""{
		vendorId, err = strconv.Atoi(vendor)
		if err != nil {
			return []model.Purchase{}, errors.New("vendor id must be an integer")
		}
	}

	purchaseData, err := s.purchaseRepository.Get(purchaseOrder, vendorId)
	if err != nil {
		return []model.Purchase{}, err
	}

	return purchaseData, nil
}

func (s *purchaseUsecase) GetById(id int) (model.Purchase, error) {
	purchaseData, err := s.purchaseRepository.GetById(id)
	if err != nil {
		return model.Purchase{}, err
	}

	return purchaseData, nil
}

func (s *purchaseUsecase) Create(payloads dto.PurchaseRequest) (model.Purchase, error) {
	tx := s.purchaseRepository.BeginTransaction()
	if tx.Error != nil {
		return model.Purchase{}, tx.Error
	}
	defer tx.Rollback() // Rollback the transaction if there's an error

	purchaseOrder, err := s.purchaseRepository.GenerateNextPurchaseOrder()
	if err != nil {
		return model.Purchase{}, err
	}

	purchaseData := model.Purchase{
		VendorId: payloads.VendorId,
		PurchaseOrder: purchaseOrder,
		Date:    time.Now(),
		Total:   0.0,
	}

	purchaseData, err = s.purchaseRepository.CreateWithTransaction(purchaseData, tx)
	if err != nil {
		return model.Purchase{}, err
	}

	total := 0.0
	for _, detail := range payloads.Details {
		_, err := s.bookRepository.GetById(int(detail.BookID))
		if err != nil {
			return model.Purchase{}, err
		}

		subtotal := detail.Price * float64(detail.Quantity)

		purchaseDetail := model.PurchaseDetail{
			PurchaseId:  purchaseData.ID,
			BookId:   detail.BookID,
			Price:    detail.Price,
			Quantity: detail.Quantity,
			Subtotal: subtotal,
		}

		// Create the purchase detail record within the transaction
		if err := s.purchaseRepository.CreatePurchaseDetailWithTransaction(purchaseDetail, tx); err != nil {
			return model.Purchase{}, err
		}

		// Deduct the stock of the book based on the quantity sold
		if err := s.purchaseRepository.AddBookStockWithTransaction(detail.BookID, detail.Quantity, tx); err != nil {
			return model.Purchase{}, err
		}

		total += subtotal
	}

	err = s.purchaseRepository.UpdateTotalWithTransaction(purchaseData.ID, total, tx)
	if err != nil {
		return model.Purchase{}, err
	}

	// Commit the transaction if everything was successful
	if err := s.purchaseRepository.CommitTransaction(tx); err != nil {
		return model.Purchase{}, err
	}

	purchaseData, err = s.purchaseRepository.GetById(int(purchaseData.ID))
	if err != nil {
		return model.Purchase{}, err
	}

	return purchaseData, nil
}