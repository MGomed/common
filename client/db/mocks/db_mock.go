// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	db "github.com/MGomed/common/client/db"
	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// DB mocks base method.
func (m *MockClient) DB() db.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(db.DB)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockClientMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockClient)(nil).DB))
}

// MockTxManager is a mock of TxManager interface.
type MockTxManager struct {
	ctrl     *gomock.Controller
	recorder *MockTxManagerMockRecorder
}

// MockTxManagerMockRecorder is the mock recorder for MockTxManager.
type MockTxManagerMockRecorder struct {
	mock *MockTxManager
}

// NewMockTxManager creates a new mock instance.
func NewMockTxManager(ctrl *gomock.Controller) *MockTxManager {
	mock := &MockTxManager{ctrl: ctrl}
	mock.recorder = &MockTxManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxManager) EXPECT() *MockTxManagerMockRecorder {
	return m.recorder
}

// ReadCommitted mocks base method.
func (m *MockTxManager) ReadCommitted(ctx context.Context, f db.Handler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCommitted", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadCommitted indicates an expected call of ReadCommitted.
func (mr *MockTxManagerMockRecorder) ReadCommitted(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCommitted", reflect.TypeOf((*MockTxManager)(nil).ReadCommitted), ctx, f)
}

// MockTransactor is a mock of Transactor interface.
type MockTransactor struct {
	ctrl     *gomock.Controller
	recorder *MockTransactorMockRecorder
}

// MockTransactorMockRecorder is the mock recorder for MockTransactor.
type MockTransactorMockRecorder struct {
	mock *MockTransactor
}

// NewMockTransactor creates a new mock instance.
func NewMockTransactor(ctrl *gomock.Controller) *MockTransactor {
	mock := &MockTransactor{ctrl: ctrl}
	mock.recorder = &MockTransactorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactor) EXPECT() *MockTransactorMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockTransactor) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, txOptions)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockTransactorMockRecorder) BeginTx(ctx, txOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockTransactor)(nil).BeginTx), ctx, txOptions)
}

// MockSQLExecer is a mock of SQLExecer interface.
type MockSQLExecer struct {
	ctrl     *gomock.Controller
	recorder *MockSQLExecerMockRecorder
}

// MockSQLExecerMockRecorder is the mock recorder for MockSQLExecer.
type MockSQLExecerMockRecorder struct {
	mock *MockSQLExecer
}

// NewMockSQLExecer creates a new mock instance.
func NewMockSQLExecer(ctrl *gomock.Controller) *MockSQLExecer {
	mock := &MockSQLExecer{ctrl: ctrl}
	mock.recorder = &MockSQLExecerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLExecer) EXPECT() *MockSQLExecerMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockSQLExecer) Exec(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockSQLExecerMockRecorder) Exec(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockSQLExecer)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockSQLExecer) Query(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockSQLExecerMockRecorder) Query(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockSQLExecer)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockSQLExecer) QueryRow(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockSQLExecerMockRecorder) QueryRow(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockSQLExecer)(nil).QueryRow), varargs...)
}

// ScanAll mocks base method.
func (m *MockSQLExecer) ScanAll(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanAll", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanAll indicates an expected call of ScanAll.
func (mr *MockSQLExecerMockRecorder) ScanAll(ctx, dest, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanAll", reflect.TypeOf((*MockSQLExecer)(nil).ScanAll), varargs...)
}

// ScanOne mocks base method.
func (m *MockSQLExecer) ScanOne(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanOne", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanOne indicates an expected call of ScanOne.
func (mr *MockSQLExecerMockRecorder) ScanOne(ctx, dest, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanOne", reflect.TypeOf((*MockSQLExecer)(nil).ScanOne), varargs...)
}

// MockNamedExecer is a mock of NamedExecer interface.
type MockNamedExecer struct {
	ctrl     *gomock.Controller
	recorder *MockNamedExecerMockRecorder
}

// MockNamedExecerMockRecorder is the mock recorder for MockNamedExecer.
type MockNamedExecerMockRecorder struct {
	mock *MockNamedExecer
}

// NewMockNamedExecer creates a new mock instance.
func NewMockNamedExecer(ctrl *gomock.Controller) *MockNamedExecer {
	mock := &MockNamedExecer{ctrl: ctrl}
	mock.recorder = &MockNamedExecerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamedExecer) EXPECT() *MockNamedExecerMockRecorder {
	return m.recorder
}

// ScanAll mocks base method.
func (m *MockNamedExecer) ScanAll(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanAll", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanAll indicates an expected call of ScanAll.
func (mr *MockNamedExecerMockRecorder) ScanAll(ctx, dest, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanAll", reflect.TypeOf((*MockNamedExecer)(nil).ScanAll), varargs...)
}

// ScanOne mocks base method.
func (m *MockNamedExecer) ScanOne(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanOne", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanOne indicates an expected call of ScanOne.
func (mr *MockNamedExecerMockRecorder) ScanOne(ctx, dest, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanOne", reflect.TypeOf((*MockNamedExecer)(nil).ScanOne), varargs...)
}

// MockQueryExecer is a mock of QueryExecer interface.
type MockQueryExecer struct {
	ctrl     *gomock.Controller
	recorder *MockQueryExecerMockRecorder
}

// MockQueryExecerMockRecorder is the mock recorder for MockQueryExecer.
type MockQueryExecerMockRecorder struct {
	mock *MockQueryExecer
}

// NewMockQueryExecer creates a new mock instance.
func NewMockQueryExecer(ctrl *gomock.Controller) *MockQueryExecer {
	mock := &MockQueryExecer{ctrl: ctrl}
	mock.recorder = &MockQueryExecerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryExecer) EXPECT() *MockQueryExecerMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockQueryExecer) Exec(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockQueryExecerMockRecorder) Exec(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockQueryExecer)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockQueryExecer) Query(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockQueryExecerMockRecorder) Query(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockQueryExecer)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockQueryExecer) QueryRow(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockQueryExecerMockRecorder) QueryRow(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockQueryExecer)(nil).QueryRow), varargs...)
}

// MockPinger is a mock of Pinger interface.
type MockPinger struct {
	ctrl     *gomock.Controller
	recorder *MockPingerMockRecorder
}

// MockPingerMockRecorder is the mock recorder for MockPinger.
type MockPingerMockRecorder struct {
	mock *MockPinger
}

// NewMockPinger creates a new mock instance.
func NewMockPinger(ctrl *gomock.Controller) *MockPinger {
	mock := &MockPinger{ctrl: ctrl}
	mock.recorder = &MockPingerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPinger) EXPECT() *MockPingerMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockPinger) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockPingerMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockPinger)(nil).Ping), ctx)
}

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockDB) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, txOptions)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockDBMockRecorder) BeginTx(ctx, txOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockDB)(nil).BeginTx), ctx, txOptions)
}

// Close mocks base method.
func (m *MockDB) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDB)(nil).Close))
}

// Exec mocks base method.
func (m *MockDB) Exec(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockDBMockRecorder) Exec(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDB)(nil).Exec), varargs...)
}

// Ping mocks base method.
func (m *MockDB) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockDBMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDB)(nil).Ping), ctx)
}

// Query mocks base method.
func (m *MockDB) Query(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDBMockRecorder) Query(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDB)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockDB) QueryRow(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockDBMockRecorder) QueryRow(ctx, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockDB)(nil).QueryRow), varargs...)
}

// ScanAll mocks base method.
func (m *MockDB) ScanAll(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanAll", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanAll indicates an expected call of ScanAll.
func (mr *MockDBMockRecorder) ScanAll(ctx, dest, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanAll", reflect.TypeOf((*MockDB)(nil).ScanAll), varargs...)
}

// ScanOne mocks base method.
func (m *MockDB) ScanOne(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanOne", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanOne indicates an expected call of ScanOne.
func (mr *MockDBMockRecorder) ScanOne(ctx, dest, q interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanOne", reflect.TypeOf((*MockDB)(nil).ScanOne), varargs...)
}
