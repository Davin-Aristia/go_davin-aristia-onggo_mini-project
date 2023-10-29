package usecase

import (
	"go-mini-project/model"
	"go-mini-project/dto"
	mocks "go-mini-project/mock"
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateBook(t *testing.T) {
	repo := new(mocks.BookData)
	insertData := dto.BookRequest{
		Title:    "Kisah Hidupku",
		Author:   "Davin",
		Price:   50000,
		Stock:   37,
		CategoryId: 1,
	}

	returnData := model.Book{
		Model:gorm.Model{ID:1},
		Title:    insertData.Title,
		Author:   insertData.Author,
		Price:   insertData.Price,
		Stock:   insertData.Stock,
		CategoryId: uint(insertData.CategoryId),
	}

	mockResponse := dto.ConvertToBookResponse(returnData)

	t.Run("Success Create", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(returnData, nil).Once()
		service := NewBookUsecase(repo)

		res, err := service.Create(insertData)
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Create", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(model.Book{}, assert.AnError).Once()
		service := NewBookUsecase(repo)

		res, err := service.Create(insertData)
		assert.Error(t, err)
		assert.Equal(t, dto.BookResponse{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetBooks(t *testing.T) {
	repo := new(mocks.BookData)
	var filter = struct{
		Title       string
		Author      string
		Category    int
	}{
		Title:    "Kisah",
		Author:   "Davin",
		Category: 1,
	}

	bookList := []model.Book{
		{
			Model:gorm.Model{ID:1},
			Title:    "Kisah Hidupku",
			Author:   "Davin",
			Price:   50000,
			Stock:   37,
			CategoryId: 1,
		},
		{
			Model:gorm.Model{ID:2},
			Title:    "Kisah Perjalananku",
			Author:   "Davin",
			Price:   60000,
			Stock:   12,
			CategoryId: 1,
		},
	}

	mockResponse := make([]dto.BookResponse, 0)

	for _, book := range bookList {
		mockResponse = append(mockResponse, dto.ConvertToBookResponse(book))
	}

	t.Run("Success GetBooks", func(t *testing.T) {
		repo.On("Get", filter.Title, filter.Author, filter.Category).Return(bookList, nil).Once()
		service := NewBookUsecase(repo)

		res, err := service.Get(filter.Title, filter.Author, strconv.Itoa(filter.Category))
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error GetBooks", func(t *testing.T) {
		repo.On("Get", filter.Title, filter.Author, filter.Category).Return([]model.Book{}, assert.AnError).Once()
		service := NewBookUsecase(repo)

		res, err := service.Get(filter.Title, filter.Author, strconv.Itoa(filter.Category))
		assert.Error(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetBookByID(t *testing.T) {
    repo := new(mocks.BookData)
    bookID := 1

    returnData := model.Book{
        Model: gorm.Model{ID: uint(bookID)},
        Title: "Kisah Hidupku",
        Author: "Davin",
        Price: 50000,
        Stock: 37,
        CategoryId: 1,
    }

    mockResponse := dto.ConvertToBookResponse(returnData)

    t.Run("Success Get by ID", func(t *testing.T) {
        repo.On("GetById", bookID).Return(returnData, nil).Once()
        service := NewBookUsecase(repo)

        res, err := service.GetById(bookID)
        assert.NoError(t, err)
        assert.Equal(t, mockResponse, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error Get by ID", func(t *testing.T) {
        repo.On("GetById", bookID).Return(model.Book{}, assert.AnError).Once()
        service := NewBookUsecase(repo)

        res, err := service.GetById(bookID)
        assert.Error(t, err)
        assert.Equal(t, dto.BookResponse{}, res)
        repo.AssertExpectations(t)
    })
}

func TestUpdateBook(t *testing.T) {
    repo := new(mocks.BookData)
    updateData := dto.BookRequest{
        Title:       "Updated Title",
        Author:      "Updated Author",
        Price:       60000,
        Stock:       42,
        CategoryId:  2,
    }
    bookID := 1

	returnData := model.Book{
		Model:gorm.Model{ID:1},
		Title:      updateData.Title,
		Author:     updateData.Author,
		Price:   	updateData.Price,
		Stock: 		updateData.Stock,
		CategoryId: uint(updateData.CategoryId),
	}

    // Mock the Get and Update methods
    repo.On("GetById", bookID).Return(model.Book{Model:gorm.Model{ID:uint(bookID)}}, nil).Once()
    repo.On("Update", mock.Anything, bookID).Return(nil).Once()
    repo.On("GetById", bookID).Return(returnData, nil).Once()

    service := NewBookUsecase(repo)

    t.Run("Success Update", func(t *testing.T) {
        res, err := service.Update(updateData, bookID)
        assert.NoError(t, err)
        expectedResponse := dto.ConvertToBookResponse(returnData)
        assert.Equal(t, expectedResponse, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error GetByID", func(t *testing.T) {
        repo.On("GetById", bookID).Return(model.Book{}, assert.AnError).Once()
        res, err := service.Update(updateData, bookID)
        assert.Error(t, err)
        assert.Equal(t, dto.BookResponse{}, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error Update", func(t *testing.T) {
        repo.On("GetById", bookID).Return(model.Book{Model:gorm.Model{ID:uint(bookID)}}, nil).Once()
        repo.On("Update", mock.Anything, bookID).Return(assert.AnError).Once()
        res, err := service.Update(updateData, bookID)
        assert.Error(t, err)
        assert.Equal(t, dto.BookResponse{}, res)
        repo.AssertExpectations(t)
    })
}

func TestDeleteBook(t *testing.T) {
    repo := new(mocks.BookData)
    bookID := 1

    // Mock the GetById and Delete methods
    repo.On("GetById", bookID).Return(model.Book{Model:gorm.Model{ID:uint(bookID)}}, nil).Once()
    repo.On("Delete", bookID).Return(nil).Once()

    service := NewBookUsecase(repo)

    t.Run("Success Delete", func(t *testing.T) {
        err := service.Delete(bookID)
        assert.NoError(t, err)
        repo.AssertExpectations(t)
    })

    t.Run("Error GetByID", func(t *testing.T) {
        repo.On("GetById", bookID).Return(model.Book{}, assert.AnError).Once()
        err := service.Delete(bookID)
        assert.Error(t, err)
        repo.AssertExpectations(t)
    })

    t.Run("Error Delete", func(t *testing.T) {
        repo.On("GetById", bookID).Return(model.Book{Model:gorm.Model{ID:uint(bookID)}}, nil).Once()
        repo.On("Delete", bookID).Return(assert.AnError).Once()
        err := service.Delete(bookID)
        assert.Error(t, err)
        repo.AssertExpectations(t)
    })
}