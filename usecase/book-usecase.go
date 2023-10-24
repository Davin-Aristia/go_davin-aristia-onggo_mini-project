package usecase

import (
	"errors"
	"strconv"

	"go-mini-project/dto"
	"go-mini-project/model"
	"go-mini-project/repository"
)

type BookUsecase interface {
	Get(title, author, category string) ([]dto.BookResponse, error)
	GetById(id int) (dto.BookResponse, error)
	Create(payloads dto.BookRequest) (dto.BookResponse, error)
	Update(payloads dto.BookRequest, id int) (dto.BookResponse, error)
	Delete(id int) error
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) *bookUsecase {
	return &bookUsecase{bookRepository: bookRepo}
}

func (s *bookUsecase) Get(title, author, category string) ([]dto.BookResponse, error) {
	categ := 0
	var err error
	if category != ""{
		categ, err = strconv.Atoi(category)
		if err != nil {
			return []dto.BookResponse{}, errors.New("category must be an integer")
		}
	}

	bookData, err := s.bookRepository.Get(title, author, categ)
	if err != nil {
		return []dto.BookResponse{}, err
	}

	var bookResponse []dto.BookResponse
	for _, book := range bookData {
		bookResponse = append(bookResponse, dto.ConvertToBookResponse(book))
	}

	return bookResponse, nil
}

func (s *bookUsecase) GetById(id int) (dto.BookResponse, error) {
	bookData, err := s.bookRepository.GetById(id)
	if err != nil {
		return dto.BookResponse{}, err
	}

	bookResponse := dto.ConvertToBookResponse(bookData)

	return bookResponse, nil
}

func (s *bookUsecase) Create(payloads dto.BookRequest) (dto.BookResponse, error) {
	bookData := model.Book{
		Title : payloads.Title,
		Author : payloads.Author,
		Price : payloads.Price,
		Stock : payloads.Stock,
		CategoryId : uint(payloads.CategoryId),
	}

	bookData, err := s.bookRepository.Create(bookData)
	if err != nil {
		return dto.BookResponse{}, err
	}

	bookResponse := dto.ConvertToBookResponse(bookData)

	return bookResponse, nil
}

func (s *bookUsecase) Update(payloads dto.BookRequest, id int) (dto.BookResponse, error) {
	_, err := s.bookRepository.GetById(id)
	if err != nil {
		return dto.BookResponse{}, err
	}
	
	bookData := model.Book{
		Title : payloads.Title,
		Author : payloads.Author,
		Price : payloads.Price,
		Stock : payloads.Stock,
		CategoryId : uint(payloads.CategoryId),
	}

	err = s.bookRepository.Update(bookData, id)
	if err != nil {
		return dto.BookResponse{}, err
	}

	//recall the GetById repo because if I return it from update, it only fill the updated field and leaves everything else null or 0
	bookData , err = s.bookRepository.GetById(id)
	if err != nil {
		return dto.BookResponse{}, err
	}

	bookResponse := dto.ConvertToBookResponse(bookData)

	return bookResponse, nil
}

func (s *bookUsecase) Delete(id int) error {
	_, err := s.bookRepository.GetById(id)
	if err != nil {
		return err
	}

	err = s.bookRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
