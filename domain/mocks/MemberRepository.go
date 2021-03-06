// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	domain "github.com/junminhong/member-center-service/domain"

	mock "github.com/stretchr/testify/mock"
)

// MemberRepository is an autogenerated mock type for the MemberRepository type
type MemberRepository struct {
	mock.Mock
}

// GetEmail provides a mock function with given fields: ctx, memberUUID
func (_m *MemberRepository) GetEmail(ctx *gin.Context, memberUUID string) string {
	ret := _m.Called(ctx, memberUUID)

	var r0 string
	if rf, ok := ret.Get(0).(func(*gin.Context, string) string); ok {
		r0 = rf(ctx, memberUUID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Register provides a mock function with given fields: ctx, member
func (_m *MemberRepository) Register(ctx *gin.Context, member *domain.Member) bool {
	ret := _m.Called(ctx, member)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gin.Context, *domain.Member) bool); ok {
		r0 = rf(ctx, member)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
