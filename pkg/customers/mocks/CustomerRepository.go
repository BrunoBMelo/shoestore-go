// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import customers "github.com/BrunoBMelo/shoestore/pkg/customers"
import mock "github.com/stretchr/testify/mock"

// CustomerRepository is an autogenerated mock type for the CustomerRepository type
type CustomerRepository struct {
	mock.Mock
}

// CheckIfIdentificationExist provides a mock function with given fields: identification
func (_m *CustomerRepository) CheckIfIdentificationExist(identification string) (bool, error) {
	ret := _m.Called(identification)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(identification)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identification)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckIfMailAlredyExist provides a mock function with given fields: mail
func (_m *CustomerRepository) CheckIfMailAlredyExist(mail string) (bool, error) {
	ret := _m.Called(mail)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(mail)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(mail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: customer
func (_m *CustomerRepository) Save(customer *customers.Customer) error {
	ret := _m.Called(customer)

	var r0 error
	if rf, ok := ret.Get(0).(func(*customers.Customer) error); ok {
		r0 = rf(customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}