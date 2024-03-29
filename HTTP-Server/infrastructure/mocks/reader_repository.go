// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "httpserver/domain"

import mock "github.com/stretchr/testify/mock"

// ReaderRepository is an autogenerated mock type for the ReaderRepository type
type ReaderRepository struct {
	mock.Mock
}

// DeleteReader provides a mock function with given fields: id
func (_m *ReaderRepository) DeleteReader(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReaderByID provides a mock function with given fields: id
func (_m *ReaderRepository) GetReaderByID(id int) (*domain.Reader, error) {
	ret := _m.Called(id)

	var r0 *domain.Reader
	if rf, ok := ret.Get(0).(func(int) *domain.Reader); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReaders provides a mock function with given fields:
func (_m *ReaderRepository) GetReaders() ([]*domain.Reader, error) {
	ret := _m.Called()

	var r0 []*domain.Reader
	if rf, ok := ret.Get(0).(func() []*domain.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertReader provides a mock function with given fields: reader
func (_m *ReaderRepository) InsertReader(reader *domain.Reader) (interface{}, error) {
	ret := _m.Called(reader)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*domain.Reader) interface{}); ok {
		r0 = rf(reader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Reader) error); ok {
		r1 = rf(reader)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReader provides a mock function with given fields: reader
func (_m *ReaderRepository) UpdateReader(reader *domain.Reader) (int, error) {
	ret := _m.Called(reader)

	var r0 int
	if rf, ok := ret.Get(0).(func(*domain.Reader) int); ok {
		r0 = rf(reader)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Reader) error); ok {
		r1 = rf(reader)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
