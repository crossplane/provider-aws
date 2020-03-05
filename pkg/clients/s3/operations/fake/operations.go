// Code generated by mockery v1.0.0. DO NOT EDIT.

package fake

import (
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	operations "github.com/crossplane/provider-aws/pkg/clients/s3/operations"
	mock "github.com/stretchr/testify/mock"
)

// Operations is an autogenerated mock type for the Operations type
type Operations struct {
	mock.Mock
}

// CreateBucketRequest provides a mock function with given fields: _a0
func (_m *Operations) CreateBucketRequest(_a0 *s3.CreateBucketInput) operations.CreateBucketRequest {
	ret := _m.Called(_a0)

	var r0 operations.CreateBucketRequest
	if rf, ok := ret.Get(0).(func(*s3.CreateBucketInput) operations.CreateBucketRequest); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(operations.CreateBucketRequest)
		}
	}

	return r0
}

// DeleteBucketRequest provides a mock function with given fields: _a0
func (_m *Operations) DeleteBucketRequest(_a0 *s3.DeleteBucketInput) operations.DeleteBucketRequest {
	ret := _m.Called(_a0)

	var r0 operations.DeleteBucketRequest
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketInput) operations.DeleteBucketRequest); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(operations.DeleteBucketRequest)
		}
	}

	return r0
}

// GetBucketVersioningRequest provides a mock function with given fields: _a0
func (_m *Operations) GetBucketVersioningRequest(_a0 *s3.GetBucketVersioningInput) operations.GetBucketVersioningRequest {
	ret := _m.Called(_a0)

	var r0 operations.GetBucketVersioningRequest
	if rf, ok := ret.Get(0).(func(*s3.GetBucketVersioningInput) operations.GetBucketVersioningRequest); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(operations.GetBucketVersioningRequest)
		}
	}

	return r0
}

// PutBucketACLRequest provides a mock function with given fields: _a0
func (_m *Operations) PutBucketACLRequest(_a0 *s3.PutBucketAclInput) operations.PutBucketACLRequest {
	ret := _m.Called(_a0)

	var r0 operations.PutBucketACLRequest
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAclInput) operations.PutBucketACLRequest); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(operations.PutBucketACLRequest)
		}
	}

	return r0
}

// PutBucketVersioningRequest provides a mock function with given fields: _a0
func (_m *Operations) PutBucketVersioningRequest(_a0 *s3.PutBucketVersioningInput) operations.PutBucketVersioningRequest {
	ret := _m.Called(_a0)

	var r0 operations.PutBucketVersioningRequest
	if rf, ok := ret.Get(0).(func(*s3.PutBucketVersioningInput) operations.PutBucketVersioningRequest); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(operations.PutBucketVersioningRequest)
		}
	}

	return r0
}
