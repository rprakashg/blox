// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: handler/store/etcd_txstore.go

package mocks

import (
	context "context"
	clientv3 "github.com/coreos/etcd/clientv3"
	concurrency "github.com/coreos/etcd/clientv3/concurrency"
	gomock "github.com/golang/mock/gomock"
)

// Mock of EtcdTXStore interface
type MockEtcdTXStore struct {
	ctrl     *gomock.Controller
	recorder *_MockEtcdTXStoreRecorder
}

// Recorder for MockEtcdTXStore (not exported)
type _MockEtcdTXStoreRecorder struct {
	mock *MockEtcdTXStore
}

func NewMockEtcdTXStore(ctrl *gomock.Controller) *MockEtcdTXStore {
	mock := &MockEtcdTXStore{ctrl: ctrl}
	mock.recorder = &_MockEtcdTXStoreRecorder{mock}
	return mock
}

func (_m *MockEtcdTXStore) EXPECT() *_MockEtcdTXStoreRecorder {
	return _m.recorder
}

func (_m *MockEtcdTXStore) GetV3Client() *clientv3.Client {
	ret := _m.ctrl.Call(_m, "GetV3Client")
	ret0, _ := ret[0].(*clientv3.Client)
	return ret0
}

func (_mr *_MockEtcdTXStoreRecorder) GetV3Client() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetV3Client")
}

func (_m *MockEtcdTXStore) NewSTMRepeatable(_param0 context.Context, _param1 *clientv3.Client, _param2 func(concurrency.STM) error) (*clientv3.TxnResponse, error) {
	ret := _m.ctrl.Call(_m, "NewSTMRepeatable", _param0, _param1, _param2)
	ret0, _ := ret[0].(*clientv3.TxnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEtcdTXStoreRecorder) NewSTMRepeatable(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewSTMRepeatable", arg0, arg1, arg2)
}
