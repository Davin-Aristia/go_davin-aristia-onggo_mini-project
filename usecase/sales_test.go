package usecase

import (
    "go-mini-project/model"
    mocks "go-mini-project/mock"
    "testing"
    "errors"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestGetSales(t *testing.T) {
    salesRepo := new(mocks.SalesData)

    t.Run("Success Get Sales", func(t *testing.T) {
        user := "1"
        invoice := ""
        expectedSales := []model.Sales{}

        salesRepo.On("Get", invoice, mock.AnythingOfType("int")).Return(expectedSales, nil).Once()

        usecase := NewSalesUsecase(salesRepo, nil)

        sales, err := usecase.Get(invoice, user)
        assert.NoError(t, err)
        assert.Len(t, sales, len(expectedSales))
        salesRepo.AssertExpectations(t)
    })

    t.Run("Error Get Sales", func(t *testing.T) {
        user := "1"
        invoice := ""
        expectedError := errors.New("an error message")

        salesRepo.On("Get", invoice, mock.AnythingOfType("int")).Return([]model.Sales{}, expectedError).Once()

        usecase := NewSalesUsecase(salesRepo, nil)

        sales, err := usecase.Get(invoice, user)
        assert.Error(t, err)
        assert.Empty(t, sales)
        salesRepo.AssertExpectations(t)
    })
}

func TestGetSalesById(t *testing.T) {
    salesRepo := new(mocks.SalesData)

    t.Run("Success Get Sales by ID", func(t *testing.T) {
        userID := 1
        salesID := 1
        role := "admin"
        expectedSales := model.Sales{}

        salesRepo.On("GetById", salesID).Return(expectedSales, nil).Once()

        usecase := NewSalesUsecase(salesRepo, nil)

        sales, err := usecase.GetById(salesID, userID, role)
        assert.NoError(t, err)
        assert.NotEmpty(t, sales)
        salesRepo.AssertExpectations(t)
    })

    t.Run("Error Get Sales by ID", func(t *testing.T) {
        userID := 1
        salesID := 1
        role := "customer"
        expectedError := errors.New("an error message")

        salesRepo.On("GetById", salesID).Return(model.Sales{}, expectedError).Once()

        usecase := NewSalesUsecase(salesRepo, nil)

        sales, err := usecase.GetById(salesID, userID, role)
        assert.Error(t, err)
        assert.Empty(t, sales)
        salesRepo.AssertExpectations(t)
    })
}