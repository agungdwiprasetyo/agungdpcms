// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
import mock "github.com/stretchr/testify/mock"

import shared "github.com/agungdwiprasetyo/agungdpcms/shared"

// Profile is an autogenerated mock type for the Profile type
type Profile struct {
	mock.Mock
}

// FindByResumeID provides a mock function with given fields: resumeID
func (_m *Profile) FindByResumeID(resumeID int) <-chan *domain.Profile {
	ret := _m.Called(resumeID)

	var r0 <-chan *domain.Profile
	if rf, ok := ret.Get(0).(func(int) <-chan *domain.Profile); ok {
		r0 = rf(resumeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *domain.Profile)
		}
	}

	return r0
}

// Remove provides a mock function with given fields: data
func (_m *Profile) Remove(data *domain.Profile) shared.Result {
	ret := _m.Called(data)

	var r0 shared.Result
	if rf, ok := ret.Get(0).(func(*domain.Profile) shared.Result); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(shared.Result)
	}

	return r0
}

// Save provides a mock function with given fields: data
func (_m *Profile) Save(data *domain.Profile) shared.Result {
	ret := _m.Called(data)

	var r0 shared.Result
	if rf, ok := ret.Get(0).(func(*domain.Profile) shared.Result); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(shared.Result)
	}

	return r0
}
