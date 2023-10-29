package mock

import (
    "github.com/stretchr/testify/mock"
    "go-mini-project/model"
    "gorm.io/gorm"
)

type PurchaseData struct {
    mock.Mock
}

func (_m *PurchaseData) Get(invoice string, user int) ([]model.Purchase, error) {
    ret := _m.Called(invoice, user)

    var r0 []model.Purchase
    if rf, ok := ret.Get(0).(func(string, int) []model.Purchase); ok {
        r0 = rf(invoice, user)
    } else {
        r0 = ret.Get(0).([]model.Purchase)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(string, int) error); ok {
        r1 = rf(invoice, user)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *PurchaseData) GetById(id int) (model.Purchase, error) {
    ret := _m.Called(id)

    var r0 model.Purchase
    if rf, ok := ret.Get(0).(func(int) model.Purchase); ok {
        r0 = rf(id)
    } else {
        r0 = ret.Get(0).(model.Purchase)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(int) error); ok {
        r1 = rf(id)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *PurchaseData) GenerateNextPurchaseOrder() (string, error) {
    ret := _m.Called()

    var r0 string
    if rf, ok := ret.Get(0).(func() string); ok {
        r0 = rf()
    } else {
        r0 = ret.String(0)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func() error); ok {
        r1 = rf()
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *PurchaseData) BeginTransaction() *gorm.DB {
    ret := _m.Called()

    var r0 *gorm.DB
    if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
        r0 = rf()
    } else {
        r0 = ret.Get(0).(*gorm.DB)
    }

    return r0
}

func (_m *PurchaseData) CreateWithTransaction(purchase model.Purchase, tx *gorm.DB) (model.Purchase, error) {
    ret := _m.Called(purchase, tx)

    var r0 model.Purchase
    if rf, ok := ret.Get(0).(func(model.Purchase, *gorm.DB) model.Purchase); ok {
        r0 = rf(purchase, tx)
    } else {
        r0 = ret.Get(0).(model.Purchase)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(model.Purchase, *gorm.DB) error); ok {
        r1 = rf(purchase, tx)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *PurchaseData) CreatePurchaseDetailWithTransaction(purchaseDetail model.PurchaseDetail, tx *gorm.DB) error {
    ret := _m.Called(purchaseDetail, tx)

    var r0 error
    if rf, ok := ret.Get(0).(func(model.PurchaseDetail, *gorm.DB) error); ok {
        r0 = rf(purchaseDetail, tx)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

func (_m *PurchaseData) AddBookStockWithTransaction(bookID uint, quantity int, tx *gorm.DB) error {
    ret := _m.Called(bookID, quantity, tx)

    var r0 error
    if rf, ok := ret.Get(0).(func(uint, int, *gorm.DB) error); ok {
        r0 = rf(bookID, quantity, tx)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

func (_m *PurchaseData) UpdateTotalWithTransaction(ID uint, total float64, tx *gorm.DB) error {
    ret := _m.Called(ID, total, tx)

    var r0 error
    if rf, ok := ret.Get(0).(func(uint, float64, *gorm.DB) error); ok {
        r0 = rf(ID, total, tx)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

func (_m *PurchaseData) CommitTransaction(tx *gorm.DB) error {
    ret := _m.Called(tx)

    var r0 error
    if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
        r0 = rf(tx)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}