// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import mongo "go.mongodb.org/mongo-driver/mongo"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Open provides a mock function with given fields: ctx
func (_m *Client) Open(ctx context.Context) (*mongo.Client, error) {
	ret := _m.Called(ctx)

	var r0 *mongo.Client
	if rf, ok := ret.Get(0).(func(context.Context) *mongo.Client); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
