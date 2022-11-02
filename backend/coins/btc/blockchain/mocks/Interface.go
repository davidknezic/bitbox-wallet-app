// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	btcutil "github.com/btcsuite/btcd/btcutil"
	blockchain "github.com/digitalbitbox/bitbox-wallet-app/backend/coins/btc/blockchain"

	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	wire "github.com/btcsuite/btcd/wire"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Interface) Close() {
	_m.Called()
}

// ConnectionError provides a mock function with given fields:
func (_m *Interface) ConnectionError() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EstimateFee provides a mock function with given fields: _a0
func (_m *Interface) EstimateFee(_a0 int) (btcutil.Amount, error) {
	ret := _m.Called(_a0)

	var r0 btcutil.Amount
	if rf, ok := ret.Get(0).(func(int) btcutil.Amount); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(btcutil.Amount)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMerkle provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Interface) GetMerkle(_a0 chainhash.Hash, _a1 int, _a2 func([]blockchain.TXHash, int), _a3 func(error)) {
	_m.Called(_a0, _a1, _a2, _a3)
}

// Headers provides a mock function with given fields: _a0, _a1, _a2
func (_m *Interface) Headers(_a0 int, _a1 int, _a2 func([]*wire.BlockHeader, int)) {
	_m.Called(_a0, _a1, _a2)
}

// HeadersSubscribe provides a mock function with given fields: _a0, _a1
func (_m *Interface) HeadersSubscribe(_a0 func() func(error), _a1 func(*blockchain.Header)) {
	_m.Called(_a0, _a1)
}

// RegisterOnConnectionErrorChangedEvent provides a mock function with given fields: _a0
func (_m *Interface) RegisterOnConnectionErrorChangedEvent(_a0 func(error)) {
	_m.Called(_a0)
}

// RelayFee provides a mock function with given fields:
func (_m *Interface) RelayFee() (btcutil.Amount, error) {
	ret := _m.Called()

	var r0 btcutil.Amount
	if rf, ok := ret.Get(0).(func() btcutil.Amount); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(btcutil.Amount)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScriptHashGetHistory provides a mock function with given fields: _a0
func (_m *Interface) ScriptHashGetHistory(_a0 blockchain.ScriptHashHex) (blockchain.TxHistory, error) {
	ret := _m.Called(_a0)

	var r0 blockchain.TxHistory
	if rf, ok := ret.Get(0).(func(blockchain.ScriptHashHex) blockchain.TxHistory); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blockchain.TxHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(blockchain.ScriptHashHex) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScriptHashSubscribe provides a mock function with given fields: _a0, _a1, _a2
func (_m *Interface) ScriptHashSubscribe(_a0 func() func(error), _a1 blockchain.ScriptHashHex, _a2 func(string)) {
	_m.Called(_a0, _a1, _a2)
}

// TransactionBroadcast provides a mock function with given fields: _a0
func (_m *Interface) TransactionBroadcast(_a0 *wire.MsgTx) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*wire.MsgTx) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransactionGet provides a mock function with given fields: _a0
func (_m *Interface) TransactionGet(_a0 chainhash.Hash) (*wire.MsgTx, error) {
	ret := _m.Called(_a0)

	var r0 *wire.MsgTx
	if rf, ok := ret.Get(0).(func(chainhash.Hash) *wire.MsgTx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wire.MsgTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(chainhash.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInterface creates a new instance of Interface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewInterface(t testing.TB) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
