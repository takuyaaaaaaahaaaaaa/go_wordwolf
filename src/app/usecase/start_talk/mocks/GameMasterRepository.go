// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/go-server-dev/src/app/domain"
	mock "github.com/stretchr/testify/mock"
)

// GameMasterRepository is an autogenerated mock type for the GameMasterRepository type
type GameMasterRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: _a0
func (_m *GameMasterRepository) Save(_a0 *domain.GameMaster) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.GameMaster) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
