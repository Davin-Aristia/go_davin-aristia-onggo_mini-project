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
	Get(invoice, user string) ([]dto.SalesResponse, error)
	GetById(id, user int, role string) (dto.SalesResponse, error)
	Create(payloads dto.SalesRequest, user int) (dto.SalesResponse, error)
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

func (s *salesUsecase) Get(invoice, user string) ([]dto.SalesResponse, error) {
	userId := 0
	var err error
	if user != ""{
		userId, err = strconv.Atoi(user)
		if err != nil {
			return []dto.SalesResponse{}, errors.New("user id must be an integer")
		}
	}

	salesData, err := s.salesRepository.Get(invoice, userId)
	if err != nil {
		return []dto.SalesResponse{}, err
	}

	var salesResponse []dto.SalesResponse
	for _, sales := range salesData {
		salesResponse = append(salesResponse, dto.ConvertToSalesResponse(sales))
	}

	return salesResponse, nil
}

func (s *salesUsecase) GetById(id, user int, role string) (dto.SalesResponse, error) {
	salesData, err := s.salesRepository.GetById(id)
	if err != nil {
		return dto.SalesResponse{}, err
	}

	if (role != "admin" && salesData.UserId != uint(user)){
		return dto.SalesResponse{}, errors.New("unauthorized")
	}

	salesResponse := dto.ConvertToSalesResponse(salesData)

	return salesResponse, nil
}

func (s *salesUsecase) Create(payloads dto.SalesRequest, user int) (dto.SalesResponse, error) {
	tx := s.salesRepository.BeginTransaction()
	if tx.Error != nil {
		return dto.SalesResponse{}, tx.Error
	}
	defer tx.Rollback() // Rollback the transaction if there's an error

	invoice, err := s.salesRepository.GenerateNextInvoice()
	if err != nil {
		return dto.SalesResponse{}, err
	}

	salesData := model.Sales{
		UserId:  uint(user),
		Invoice: invoice,
		Date:    time.Now(),
		Total:   0.0,
	}

	salesData, err = s.salesRepository.CreateWithTransaction(salesData, tx)
	if err != nil {
		return dto.SalesResponse{}, err
	}

	total := 0.0
	for _, detail := range payloads.Details {
		bookData, err := s.bookRepository.GetById(int(detail.BookId))
		if err != nil {
			return dto.SalesResponse{}, err
		}

		subtotal := bookData.Price * float64(detail.Quantity)

		salesDetail := model.SalesDetail{
			SalesId:  salesData.ID,
			BookId:   detail.BookId,
			Price:    bookData.Price,
			Quantity: detail.Quantity,
			Subtotal: subtotal,
		}

		// Create the sales detail record within the transaction
		if err := s.salesRepository.CreateSalesDetailWithTransaction(salesDetail, tx); err != nil {
			return dto.SalesResponse{}, err
		}

		// Deduct the stock of the book based on the quantity sold
		if err := s.salesRepository.DeductBookStockWithTransaction(detail.BookId, detail.Quantity, tx); err != nil {
			return dto.SalesResponse{}, err
		}

		total += subtotal
	}

	err = s.salesRepository.UpdateTotalWithTransaction(salesData.ID, total, tx)
	if err != nil {
		return dto.SalesResponse{}, err
	}

	// Commit the transaction if everything was successful
	if err := s.salesRepository.CommitTransaction(tx); err != nil {
		return dto.SalesResponse{}, err
	}

	salesData, err = s.salesRepository.GetById(int(salesData.ID))
	if err != nil {
		return dto.SalesResponse{}, err
	}

	salesResponse := dto.ConvertToSalesResponse(salesData)

	return salesResponse, nil
}