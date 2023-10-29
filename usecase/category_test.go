package usecase

import (
	"go-mini-project/model"
	"go-mini-project/dto"
	mocks "go-mini-project/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateCategory(t *testing.T) {
	repo := new(mocks.CategoryData)
	insertData := dto.CategoryRequest{
		Name: 		"Science Fiction",
	}

	returnData := model.Category{
		Model:gorm.Model{ID:1},
		Name:    insertData.Name,
	}

	mockResponse := dto.ConvertToCategoryResponse(returnData)

	t.Run("Success Create", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(returnData, nil).Once()
		service := NewCategoryUsecase(repo)

		res, err := service.Create(insertData)
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Create", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(model.Category{}, assert.AnError).Once()
		service := NewCategoryUsecase(repo)

		res, err := service.Create(insertData)
		assert.Error(t, err)
		assert.Equal(t, dto.CategoryResponse{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetCategorys(t *testing.T) {
	repo := new(mocks.CategoryData)
	filtername := "Science"

	categoryList := []model.Category{
		{
			Model:gorm.Model{ID:1},
			Name: 		"Science Fiction",
		},
		{
			Model:gorm.Model{ID:2},
			Name: 		"Science",
		},
	}

	mockResponse := make([]dto.CategoryResponse, 0)

	for _, category := range categoryList {
		mockResponse = append(mockResponse, dto.ConvertToCategoryResponse(category))
	}

	t.Run("Success GetCategorys", func(t *testing.T) {
		repo.On("Get", filtername).Return(categoryList, nil).Once()
		service := NewCategoryUsecase(repo)

		res, err := service.Get(filtername)
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error GetCategorys", func(t *testing.T) {
		repo.On("Get", filtername).Return([]model.Category{}, assert.AnError).Once()
		service := NewCategoryUsecase(repo)

		res, err := service.Get(filtername)
		assert.Error(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetCategoryByID(t *testing.T) {
    repo := new(mocks.CategoryData)
    categoryID := 1

    returnData := model.Category{
        Model: gorm.Model{ID: uint(categoryID)},
        Name: 		"Science Fiction",
    }

    mockResponse := dto.ConvertToCategoryResponse(returnData)

    t.Run("Success Get by ID", func(t *testing.T) {
        repo.On("GetById", categoryID).Return(returnData, nil).Once()
        service := NewCategoryUsecase(repo)

        res, err := service.GetById(categoryID)
        assert.NoError(t, err)
        assert.Equal(t, mockResponse, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error Get by ID", func(t *testing.T) {
        repo.On("GetById", categoryID).Return(model.Category{}, assert.AnError).Once()
        service := NewCategoryUsecase(repo)

        res, err := service.GetById(categoryID)
        assert.Error(t, err)
        assert.Equal(t, dto.CategoryResponse{}, res)
        repo.AssertExpectations(t)
    })
}

func TestUpdateCategory(t *testing.T) {
    repo := new(mocks.CategoryData)
    updateData := dto.CategoryRequest{
		Name: 		"Updated Science Fiction",
	}
    categoryID := 1

	returnData := model.Category{
		Model:gorm.Model{ID:1},
		Name:    updateData.Name,
	}

    // Mock the Get and Update methods
    repo.On("GetById", categoryID).Return(model.Category{Model:gorm.Model{ID:uint(categoryID)}}, nil).Once()
    repo.On("Update", mock.Anything, categoryID).Return(nil).Once()
    repo.On("GetById", categoryID).Return(returnData, nil).Once()

    service := NewCategoryUsecase(repo)

    t.Run("Success Update", func(t *testing.T) {
        res, err := service.Update(updateData, categoryID)
        assert.NoError(t, err)
        expectedResponse := dto.ConvertToCategoryResponse(returnData)
        assert.Equal(t, expectedResponse, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error GetByID", func(t *testing.T) {
        repo.On("GetById", categoryID).Return(model.Category{}, assert.AnError).Once()
        res, err := service.Update(updateData, categoryID)
        assert.Error(t, err)
        assert.Equal(t, dto.CategoryResponse{}, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error Update", func(t *testing.T) {
        repo.On("GetById", categoryID).Return(model.Category{Model:gorm.Model{ID:uint(categoryID)}}, nil).Once()
        repo.On("Update", mock.Anything, categoryID).Return(assert.AnError).Once()
        res, err := service.Update(updateData, categoryID)
        assert.Error(t, err)
        assert.Equal(t, dto.CategoryResponse{}, res)
        repo.AssertExpectations(t)
    })
}

func TestDeleteCategory(t *testing.T) {
    repo := new(mocks.CategoryData)
    categoryID := 1

    // Mock the GetById and Delete methods
    repo.On("GetById", categoryID).Return(model.Category{Model:gorm.Model{ID:uint(categoryID)}}, nil).Once()
    repo.On("Delete", categoryID).Return(nil).Once()

    service := NewCategoryUsecase(repo)

    t.Run("Success Delete", func(t *testing.T) {
        err := service.Delete(categoryID)
        assert.NoError(t, err)
        repo.AssertExpectations(t)
    })

    t.Run("Error GetByID", func(t *testing.T) {
        repo.On("GetById", categoryID).Return(model.Category{}, assert.AnError).Once()
        err := service.Delete(categoryID)
        assert.Error(t, err)
        repo.AssertExpectations(t)
    })

    t.Run("Error Delete", func(t *testing.T) {
        repo.On("GetById", categoryID).Return(model.Category{Model:gorm.Model{ID:uint(categoryID)}}, nil).Once()
        repo.On("Delete", categoryID).Return(assert.AnError).Once()
        err := service.Delete(categoryID)
        assert.Error(t, err)
        repo.AssertExpectations(t)
    })
}