package mock

import (
    "github.com/stretchr/testify/mock"
    "go-mini-project/model"
    "gorm.io/gorm"
)

type SalesData struct {
    mock.Mock
}

func (_m *SalesData) Get(invoice string, user int) ([]model.Sales, error) {
    ret := _m.Called(invoice, user)

    var r0 []model.Sales
    if rf, ok := ret.Get(0).(func(string, int) []model.Sales); ok {
        r0 = rf(invoice, user)
    } else {
        r0 = ret.Get(0).([]model.Sales)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(string, int) error); ok {
        r1 = rf(invoice, user)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *SalesData) GetById(id int) (model.Sales, error) {
    ret := _m.Called(id)

    var r0 model.Sales
    if rf, ok := ret.Get(0).(func(int) model.Sales); ok {
        r0 = rf(id)
    } else {
        r0 = ret.Get(0).(model.Sales)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(int) error); ok {
        r1 = rf(id)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *SalesData) GenerateNextInvoice() (string, error) {
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

func (_m *SalesData) BeginTransaction() *gorm.DB {
    ret := _m.Called()

    var r0 *gorm.DB
    if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
        r0 = rf()
    } else {
        r0 = ret.Get(0).(*gorm.DB)
    }

    return r0
}

func (_m *SalesData) CreateWithTransaction(sales model.Sales, tx *gorm.DB) (model.Sales, error) {
    ret := _m.Called(sales, tx)

    var r0 model.Sales
    if rf, ok := ret.Get(0).(func(model.Sales, *gorm.DB) model.Sales); ok {
        r0 = rf(sales, tx)
    } else {
        r0 = ret.Get(0).(model.Sales)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(model.Sales, *gorm.DB) error); ok {
        r1 = rf(sales, tx)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *SalesData) CreateSalesDetailWithTransaction(salesDetail model.SalesDetail, tx *gorm.DB) error {
    ret := _m.Called(salesDetail, tx)

    var r0 error
    if rf, ok := ret.Get(0).(func(model.SalesDetail, *gorm.DB) error); ok {
        r0 = rf(salesDetail, tx)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

func (_m *SalesData) DeductBookStockWithTransaction(bookID uint, quantity int, tx *gorm.DB) error {
    ret := _m.Called(bookID, quantity, tx)

    var r0 error
    if rf, ok := ret.Get(0).(func(uint, int, *gorm.DB) error); ok {
        r0 = rf(bookID, quantity, tx)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

func (_m *SalesData) UpdateTotalWithTransaction(ID uint, total float64, tx *gorm.DB) error {
    ret := _m.Called(ID, total, tx)

    var r0 error
    if rf, ok := ret.Get(0).(func(uint, float64, *gorm.DB) error); ok {
        r0 = rf(ID, total, tx)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

func (_m *SalesData) CommitTransaction(tx *gorm.DB) error {
    ret := _m.Called(tx)

    var r0 error
    if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
        r0 = rf(tx)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}