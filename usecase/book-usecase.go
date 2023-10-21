package usecase

import (
	"errors"
	"strconv"

	"go-mini-project/dto"
	// "go-mini-project/middleware"
	"go-mini-project/model"
	"go-mini-project/repository"
)

type BookUsecase interface {
	Get(title, author, category string) ([]model.Book, error)
	GetById(id int) (model.Book, error)
	Create(payloads dto.BookRequest) (model.Book, error)
	Update(payloads dto.BookRequest, id int) (model.Book, error)
	Delete(id int) error
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) *bookUsecase {
	return &bookUsecase{bookRepository: bookRepo}
}

func (s *bookUsecase) Get(title, author, category string) ([]model.Book, error) {
	categ := 0
	var err error
	if category != ""{
		categ, err = strconv.Atoi(category)
		if err != nil {
			return []model.Book{}, errors.New("category must be an integer")
		}
	}

	bookData, err := s.bookRepository.Get(title, author, categ)
	if err != nil {
		return []model.Book{}, err
	}

	return bookData, nil
}

func (s *bookUsecase) GetById(id int) (model.Book, error) {
	bookData, err := s.bookRepository.GetById(id)
	if err != nil {
		return model.Book{}, err
	}

	return bookData, nil
}

func (s *bookUsecase) Create(payloads dto.BookRequest) (model.Book, error) {
	bookData := model.Book{
		Title : payloads.Title,
		Author : payloads.Author,
		Price : payloads.Price,
		Stock : payloads.Stock,
		CategoryId : uint(payloads.CategoryId),
	}

	bookData, err := s.bookRepository.Create(bookData)
	if err != nil {
		return model.Book{}, err
	}
	return bookData, nil
}

func (s *bookUsecase) Update(payloads dto.BookRequest, id int) (model.Book, error) {
	_, err := s.bookRepository.GetById(id)
	if err != nil {
		return model.Book{}, err
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
		return model.Book{}, err
	}

	//recall the GetById repo because if I return it from update, it only fill the updated field and leaves everything else null or 0
	bookData , err = s.bookRepository.GetById(id)
	if err != nil {
		return model.Book{}, err
	}

	return bookData, nil
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
