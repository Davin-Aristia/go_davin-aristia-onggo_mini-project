package usecase

import (
	"errors"
	"strconv"
	"time"

	"go-mini-project/dto" 
	"go-mini-project/model"
	"go-mini-project/repository"
)

type SalesUsecase interface {
	Get(invoice, user, role string) ([]model.Sales, error)
	GetById(id, user int, role string) (model.Sales, error)
	Create(payloads dto.SalesRequest, user int) (model.Sales, error)
}

type salesUsecase struct {
	salesRepository repository.SalesRepository
}

func NewSalesUsecase(salesRepo repository.SalesRepository) *salesUsecase {
	return &salesUsecase{salesRepository: salesRepo}
}

func (s *salesUsecase) Get(invoice, user, role string) ([]model.Sales, error) {
	userId := 0
	var err error
	if user != ""{
		userId, err = strconv.Atoi(user)
		if err != nil {
			return []model.Sales{}, errors.New("user id must be an integer")
		}
	}

	salesData, err := s.salesRepository.Get(invoice, userId)
	if err != nil {
		return []model.Sales{}, err
	}

	return salesData, nil
}

func (s *salesUsecase) GetById(id, user int, role string) (model.Sales, error) {
	salesData, err := s.salesRepository.GetById(id)
	if err != nil {
		return model.Sales{}, err
	}

	if (role != "admin" && salesData.UserId != uint(user)){
		return model.Sales{}, errors.New("unauthorized")
	}

	return salesData, nil
}

// func (s *salesUsecase) Create(payloads dto.SalesRequest, user int) (model.Sales, error) {
// 	invoice, err := s.salesRepository.GenerateNextInvoice()
// 	if err != nil {
// 		return model.Sales{}, err
// 	}

// 	salesData := model.Sales{
// 		UserId : uint(user),
// 		Invoice : invoice,
// 		Date : time.Now(),
// 		Total : payloads.Total,
// 	}

// 	salesData, err = s.salesRepository.Create(salesData)
// 	if err != nil {
// 		return model.Sales{}, err
// 	}
// 	return salesData, nil
// }

func (s *salesUsecase) Create(payloads dto.SalesRequest, user int) (model.Sales, error) {
	// Start a database transaction
	tx := s.salesRepository.BeginTransaction()
	if tx.Error != nil {
		return model.Sales{}, tx.Error
	}
	defer tx.Rollback() // Rollback the transaction if there's an error

	invoice, err := s.salesRepository.GenerateNextInvoice()
	if err != nil {
		return model.Sales{}, err
	}

	// Calculate the total from the sales details' subtotals
	total := 0.0
	for _, detail := range payloads.Details {
		total += detail.Subtotal
	}

	salesData := model.Sales{
		UserId:  uint(user),
		Invoice: invoice,
		Date:    time.Now(),
		Total:   total,
	}

	// Create the sales record within the transaction
	salesData, err = s.salesRepository.CreateWithTransaction(salesData, tx)
	if err != nil {
		return model.Sales{}, err
	}

	// Create the sales details within the transaction
	for _, detail := range payloads.Details {
		salesDetail := model.SalesDetail{
			SalesId:  salesData.ID,
			BookId:   detail.BookID,
			Price:    detail.Price,
			Quantity: detail.Quantity,
			Subtotal: detail.Subtotal,
		}

		// Create the sales detail record within the transaction
		if err := s.salesRepository.CreateSalesDetailWithTransaction(salesDetail, tx); err != nil {
			return model.Sales{}, err
		}
	}

	// Commit the transaction if everything was successful
	if err := s.salesRepository.CommitTransaction(tx); err != nil {
		return model.Sales{}, err
	}

	return salesData, nil
}