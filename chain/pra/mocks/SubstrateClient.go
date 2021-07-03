// Code generated by mockery v2.0.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/centrifuge/go-substrate-rpc-client/v3/gethrpc"

	types "github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

// SubstrateClient is an autogenerated mock type for the SubstrateClient type
type SubstrateClient struct {
	mock.Mock
}

// Call provides a mock function with given fields: result, method, args
func (_m *SubstrateClient) Call(result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBlockHash provides a mock function with given fields: blockNumber
func (_m *SubstrateClient) GetBlockHash(blockNumber uint64) (types.Hash, error) {
	ret := _m.Called(blockNumber)

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func(uint64) types.Hash); ok {
		r0 = rf(blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHashLatest provides a mock function with given fields:
func (_m *SubstrateClient) GetBlockHashLatest() (types.Hash, error) {
	ret := _m.Called()

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func() types.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
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

// GetFinalizedHead provides a mock function with given fields:
func (_m *SubstrateClient) GetFinalizedHead() (types.Hash, error) {
	ret := _m.Called()

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func() types.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
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

// GetHeader provides a mock function with given fields: hash
func (_m *SubstrateClient) GetHeader(hash types.Hash) (*types.Header, error) {
	ret := _m.Called(hash)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(types.Hash) *types.Header); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeaderLatest provides a mock function with given fields:
func (_m *SubstrateClient) GetHeaderLatest() (*types.Header, error) {
	ret := _m.Called()

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func() *types.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
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

// GetMetadata provides a mock function with given fields: blockHash
func (_m *SubstrateClient) GetMetadata(blockHash types.Hash) (*types.Metadata, error) {
	ret := _m.Called(blockHash)

	var r0 *types.Metadata
	if rf, ok := ret.Get(0).(func(types.Hash) *types.Metadata); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageRaw provides a mock function with given fields: key, blockHash
func (_m *SubstrateClient) GetStorageRaw(key types.StorageKey, blockHash types.Hash) (*types.StorageDataRaw, error) {
	ret := _m.Called(key, blockHash)

	var r0 *types.StorageDataRaw
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.Hash) *types.StorageDataRaw); ok {
		r0 = rf(key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.Hash) error); ok {
		r1 = rf(key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: ctx, namespace, subscribeMethodSuffix, unsubscribeMethodSuffix, notificationMethodSuffix, channel, args
func (_m *SubstrateClient) Subscribe(ctx context.Context, namespace string, subscribeMethodSuffix string, unsubscribeMethodSuffix string, notificationMethodSuffix string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, namespace, subscribeMethodSuffix, unsubscribeMethodSuffix, notificationMethodSuffix, channel)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *rpc.ClientSubscription
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, interface{}, ...interface{}) *rpc.ClientSubscription); ok {
		r0 = rf(ctx, namespace, subscribeMethodSuffix, unsubscribeMethodSuffix, notificationMethodSuffix, channel, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.ClientSubscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, interface{}, ...interface{}) error); ok {
		r1 = rf(ctx, namespace, subscribeMethodSuffix, unsubscribeMethodSuffix, notificationMethodSuffix, channel, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URL provides a mock function with given fields:
func (_m *SubstrateClient) URL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
