// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	database "github.com/fsvxavier/unico/database"

	mock "github.com/stretchr/testify/mock"

	models "github.com/fsvxavier/unico/internal/models"
)

// FeirasUsecases is an autogenerated mock type for the FeirasUsecases type
type FeirasUsecases struct {
	mock.Mock
}

// CreateFeira provides a mock function with given fields: feira
func (_m *FeirasUsecases) CreateFeira(feira models.FeiraLivre) (models.FeiraLivre, error) {
	ret := _m.Called(feira)

	var r0 models.FeiraLivre
	if rf, ok := ret.Get(0).(func(models.FeiraLivre) models.FeiraLivre); ok {
		r0 = rf(feira)
	} else {
		r0 = ret.Get(0).(models.FeiraLivre)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.FeiraLivre) error); ok {
		r1 = rf(feira)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFeira provides a mock function with given fields: id
func (_m *FeirasUsecases) DeleteFeira(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFeira provides a mock function with given fields: id
func (_m *FeirasUsecases) GetFeira(id string) (models.FeiraLivre, error) {
	ret := _m.Called(id)

	var r0 models.FeiraLivre
	if rf, ok := ret.Get(0).(func(string) models.FeiraLivre); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.FeiraLivre)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFeiraSearch provides a mock function with given fields: feira
func (_m *FeirasUsecases) GetFeiraSearch(feira models.SearchFeira) (database.Pagination, error) {
	ret := _m.Called(feira)

	var r0 database.Pagination
	if rf, ok := ret.Get(0).(func(models.SearchFeira) database.Pagination); ok {
		r0 = rf(feira)
	} else {
		r0 = ret.Get(0).(database.Pagination)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.SearchFeira) error); ok {
		r1 = rf(feira)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFeirasPagination provides a mock function with given fields: page
func (_m *FeirasUsecases) GetFeirasPagination(page string) (database.Pagination, error) {
	ret := _m.Called(page)

	var r0 database.Pagination
	if rf, ok := ret.Get(0).(func(string) database.Pagination); ok {
		r0 = rf(page)
	} else {
		r0 = ret.Get(0).(database.Pagination)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFeira provides a mock function with given fields: id, feira
func (_m *FeirasUsecases) UpdateFeira(id string, feira models.FeiraLivre) (models.FeiraLivre, error) {
	ret := _m.Called(id, feira)

	var r0 models.FeiraLivre
	if rf, ok := ret.Get(0).(func(string, models.FeiraLivre) models.FeiraLivre); ok {
		r0 = rf(id, feira)
	} else {
		r0 = ret.Get(0).(models.FeiraLivre)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, models.FeiraLivre) error); ok {
		r1 = rf(id, feira)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFeirasUsecases interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeirasUsecases creates a new instance of FeirasUsecases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeirasUsecases(t mockConstructorTestingTNewFeirasUsecases) *FeirasUsecases {
	mock := &FeirasUsecases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
