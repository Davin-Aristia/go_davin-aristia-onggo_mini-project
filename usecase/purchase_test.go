package usecase

import (
    "go-mini-project/model"
    mocks "go-mini-project/mock"
    "testing"
    "errors"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestGetPurchase(t *testing.T) {
    purchaseRepo := new(mocks.PurchaseData)

    t.Run("Success Get Purchase", func(t *testing.T) {
        user := "1"
        invoice := ""
        expectedPurchase := []model.Purchase{}

       
        purchaseRepo.On("Get", invoice, mock.AnythingOfType("int")).Return(expectedPurchase, nil).Once()

        usecase := NewPurchaseUsecase(purchaseRepo, nil)

       
        purchase, err := usecase.Get(invoice, user)
        assert.NoError(t, err)
        assert.Len(t, purchase, len(expectedPurchase))
        purchaseRepo.AssertExpectations(t)
    })

    t.Run("Error Get Purchase", func(t *testing.T) {
        user := "1"
        invoice := ""
        expectedError := errors.New("an error message")

        purchaseRepo.On("Get", invoice, mock.AnythingOfType("int")).Return([]model.Purchase{}, expectedError).Once()

        usecase := NewPurchaseUsecase(purchaseRepo, nil)

        purchase, err := usecase.Get(invoice, user)
        assert.Error(t, err)
        assert.Empty(t, purchase)
        purchaseRepo.AssertExpectations(t)
    })
}

func TestGetPurchaseById(t *testing.T) {
    purchaseRepo := new(mocks.PurchaseData)

    t.Run("Success Get Purchase by ID", func(t *testing.T) {
        purchaseID := 1
        expectedPurchase := model.Purchase{}

        purchaseRepo.On("GetById", purchaseID).Return(expectedPurchase, nil).Once()

        usecase := NewPurchaseUsecase(purchaseRepo, nil)

        purchase, err := usecase.GetById(purchaseID)
        assert.NoError(t, err)
        assert.NotEmpty(t, purchase)
        purchaseRepo.AssertExpectations(t)
    })

    t.Run("Error Get Purchase by ID", func(t *testing.T) {
        purchaseID := 1
        expectedError := errors.New("an error message")

        purchaseRepo.On("GetById", purchaseID).Return(model.Purchase{}, expectedError).Once()

        usecase := NewPurchaseUsecase(purchaseRepo, nil)

        purchase, err := usecase.GetById(purchaseID)
        assert.Error(t, err)
        assert.Empty(t, purchase)
        purchaseRepo.AssertExpectations(t)
    })
}