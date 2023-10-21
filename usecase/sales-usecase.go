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
	bookRepository repository.BookRepository
}

func NewSalesUsecase(salesRepo repository.SalesRepository, bookRepo repository.BookRepository) *salesUsecase {
	return &salesUsecase{
		salesRepository: salesRepo,
		bookRepository: bookRepo,
	}
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

func (s *salesUsecase) Create(payloads dto.SalesRequest, user int) (model.Sales, error) {
	tx := s.salesRepository.BeginTransaction()
	if tx.Error != nil {
		return model.Sales{}, tx.Error
	}
	defer tx.Rollback() // Rollback the transaction if there's an error

	invoice, err := s.salesRepository.GenerateNextInvoice()
	if err != nil {
		return model.Sales{}, err
	}

	salesData := model.Sales{
		UserId:  uint(user),
		Invoice: invoice,
		Date:    time.Now(),
		Total:   0.0,
	}

	salesData, err = s.salesRepository.CreateWithTransaction(salesData, tx)
	if err != nil {
		return model.Sales{}, err
	}

	total := 0.0
	for _, detail := range payloads.Details {
		bookData, err := s.bookRepository.GetById(int(detail.BookID))
		if err != nil {
			return model.Sales{}, err
		}

		subtotal := bookData.Price * float64(detail.Quantity)

		salesDetail := model.SalesDetail{
			SalesId:  salesData.ID,
			BookId:   detail.BookID,
			Price:    bookData.Price,
			Quantity: detail.Quantity,
			Subtotal: subtotal,
		}

		// Create the sales detail record within the transaction
		if err := s.salesRepository.CreateSalesDetailWithTransaction(salesDetail, tx); err != nil {
			return model.Sales{}, err
		}

		// Deduct the stock of the book based on the quantity sold
		if err := s.salesRepository.DeductBookStockWithTransaction(detail.BookID, detail.Quantity, tx); err != nil {
			return model.Sales{}, err
		}

		total += subtotal
	}

	err = s.salesRepository.UpdateTotalWithTransaction(salesData.ID, total, tx)
	if err != nil {
		return model.Sales{}, err
	}

	// Commit the transaction if everything was successful
	if err := s.salesRepository.CommitTransaction(tx); err != nil {
		return model.Sales{}, err
	}

	salesData, err = s.salesRepository.GetById(int(salesData.ID))
	if err != nil {
		return model.Sales{}, err
	}

	return salesData, nil
}