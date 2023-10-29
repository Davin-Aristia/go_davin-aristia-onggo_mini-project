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

func TestCreateVendor(t *testing.T) {
	repo := new(mocks.VendorData)
	insertData := dto.VendorRequest{
		Name: 		"Davin",
		Email:      "davin@gmail.com",
		Address:   "Jalan Melayang",
		PhoneNumber: "1234567890",
	}

	returnData := model.Vendor{
		Model:gorm.Model{ID:1},
		Name:    insertData.Name,
		Email:   insertData.Email,
		Address:   insertData.Address,
		PhoneNumber:   insertData.PhoneNumber,
	}

	mockResponse := dto.ConvertToVendorResponse(returnData)

	t.Run("Success Create", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(returnData, nil).Once()
		service := NewVendorUsecase(repo)

		res, err := service.Create(insertData)
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Create", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(model.Vendor{}, assert.AnError).Once()
		service := NewVendorUsecase(repo)

		res, err := service.Create(insertData)
		assert.Error(t, err)
		assert.Equal(t, dto.VendorResponse{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetVendors(t *testing.T) {
	repo := new(mocks.VendorData)
	filtername := "vin"

	vendorList := []model.Vendor{
		{
			Model:gorm.Model{ID:1},
			Name: 		"Davin",
			Email:      "davin@gmail.com",
			Address:   "Jalan Melayang",
			PhoneNumber: "1234567890",
		},
		{
			Model:gorm.Model{ID:2},
			Name: 		"Kevin",
			Email:      "kevin@gmail.com",
			Address:   "Jalan Tenggelam",
			PhoneNumber: "1234567790",
		},
	}

	mockResponse := make([]dto.VendorResponse, 0)

	for _, vendor := range vendorList {
		mockResponse = append(mockResponse, dto.ConvertToVendorResponse(vendor))
	}

	t.Run("Success GetVendors", func(t *testing.T) {
		repo.On("Get", filtername).Return(vendorList, nil).Once()
		service := NewVendorUsecase(repo)

		res, err := service.Get(filtername)
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error GetVendors", func(t *testing.T) {
		repo.On("Get", filtername).Return([]model.Vendor{}, assert.AnError).Once()
		service := NewVendorUsecase(repo)

		res, err := service.Get(filtername)
		assert.Error(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetVendorByID(t *testing.T) {
    repo := new(mocks.VendorData)
    vendorID := 1

    returnData := model.Vendor{
        Model: gorm.Model{ID: uint(vendorID)},
        Name: 		"Davin",
		Email:      "davin@gmail.com",
		Address:   "Jalan Melayang",
		PhoneNumber: "1234567890",
    }

    mockResponse := dto.ConvertToVendorResponse(returnData)

    t.Run("Success Get by ID", func(t *testing.T) {
        repo.On("GetById", vendorID).Return(returnData, nil).Once()
        service := NewVendorUsecase(repo)

        res, err := service.GetById(vendorID)
        assert.NoError(t, err)
        assert.Equal(t, mockResponse, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error Get by ID", func(t *testing.T) {
        repo.On("GetById", vendorID).Return(model.Vendor{}, assert.AnError).Once()
        service := NewVendorUsecase(repo)

        res, err := service.GetById(vendorID)
        assert.Error(t, err)
        assert.Equal(t, dto.VendorResponse{}, res)
        repo.AssertExpectations(t)
    })
}

func TestUpdateVendor(t *testing.T) {
    repo := new(mocks.VendorData)
    updateData := dto.VendorRequest{
		Name: 		"Updated Davin",
		Email:      "davin@gmail.com",
		Address:   "Updated Jalan Melayang",
		PhoneNumber: "1234567890",
	}
    vendorID := 1

	returnData := model.Vendor{
		Model:gorm.Model{ID:1},
		Name:    updateData.Name,
		Email:   updateData.Email,
		Address:   updateData.Address,
		PhoneNumber:   updateData.PhoneNumber,
	}

    // Mock the Get and Update methods
    repo.On("GetById", vendorID).Return(model.Vendor{Model:gorm.Model{ID:uint(vendorID)}}, nil).Once()
    repo.On("Update", mock.Anything, vendorID).Return(nil).Once()
    repo.On("GetById", vendorID).Return(returnData, nil).Once()

    service := NewVendorUsecase(repo)

    t.Run("Success Update", func(t *testing.T) {
        res, err := service.Update(updateData, vendorID)
        assert.NoError(t, err)
        expectedResponse := dto.ConvertToVendorResponse(returnData)
        assert.Equal(t, expectedResponse, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error GetByID", func(t *testing.T) {
        repo.On("GetById", vendorID).Return(model.Vendor{}, assert.AnError).Once()
        res, err := service.Update(updateData, vendorID)
        assert.Error(t, err)
        assert.Equal(t, dto.VendorResponse{}, res)
        repo.AssertExpectations(t)
    })

    t.Run("Error Update", func(t *testing.T) {
        repo.On("GetById", vendorID).Return(model.Vendor{Model:gorm.Model{ID:uint(vendorID)}}, nil).Once()
        repo.On("Update", mock.Anything, vendorID).Return(assert.AnError).Once()
        res, err := service.Update(updateData, vendorID)
        assert.Error(t, err)
        assert.Equal(t, dto.VendorResponse{}, res)
        repo.AssertExpectations(t)
    })
}

func TestDeleteVendor(t *testing.T) {
    repo := new(mocks.VendorData)
    vendorID := 1

    // Mock the GetById and Delete methods
    repo.On("GetById", vendorID).Return(model.Vendor{Model:gorm.Model{ID:uint(vendorID)}}, nil).Once()
    repo.On("Delete", vendorID).Return(nil).Once()

    service := NewVendorUsecase(repo)

    t.Run("Success Delete", func(t *testing.T) {
        err := service.Delete(vendorID)
        assert.NoError(t, err)
        repo.AssertExpectations(t)
    })

    t.Run("Error GetByID", func(t *testing.T) {
        repo.On("GetById", vendorID).Return(model.Vendor{}, assert.AnError).Once()
        err := service.Delete(vendorID)
        assert.Error(t, err)
        repo.AssertExpectations(t)
    })

    t.Run("Error Delete", func(t *testing.T) {
        repo.On("GetById", vendorID).Return(model.Vendor{Model:gorm.Model{ID:uint(vendorID)}}, nil).Once()
        repo.On("Delete", vendorID).Return(assert.AnError).Once()
        err := service.Delete(vendorID)
        assert.Error(t, err)
        repo.AssertExpectations(t)
    })
}