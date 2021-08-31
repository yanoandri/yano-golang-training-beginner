// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/yanoandri/yano-golang-training-beginner/model"
)

// IPaymentCodeService is an autogenerated mock type for the IPaymentCodeService type
type IPaymentCodeService struct {
	mock.Mock
}

// CreatePaymentCode provides a mock function with given fields: payment
func (_m *IPaymentCodeService) CreatePaymentCode(payment model.PaymentCodes) (model.PaymentCodes, error) {
	ret := _m.Called(payment)

	var r0 model.PaymentCodes
	if rf, ok := ret.Get(0).(func(model.PaymentCodes) model.PaymentCodes); ok {
		r0 = rf(payment)
	} else {
		r0 = ret.Get(0).(model.PaymentCodes)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.PaymentCodes) error); ok {
		r1 = rf(payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpirePaymentCode provides a mock function with given fields:
func (_m *IPaymentCodeService) ExpirePaymentCode() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetActivePaymentCodeByCode provides a mock function with given fields: code
func (_m *IPaymentCodeService) GetActivePaymentCodeByCode(code string) (model.PaymentCodes, error) {
	ret := _m.Called(code)

	var r0 model.PaymentCodes
	if rf, ok := ret.Get(0).(func(string) model.PaymentCodes); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Get(0).(model.PaymentCodes)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentCodeById provides a mock function with given fields: id
func (_m *IPaymentCodeService) GetPaymentCodeById(id string) (model.PaymentCodes, error) {
	ret := _m.Called(id)

	var r0 model.PaymentCodes
	if rf, ok := ret.Get(0).(func(string) model.PaymentCodes); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.PaymentCodes)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
