// This file was generated by counterfeiter
package dbngfakes

import (
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeConn struct {
	BeginStub        func() (dbng.Tx, error)
	beginMutex       sync.RWMutex
	beginArgsForCall []struct{}
	beginReturns     struct {
		result1 dbng.Tx
		result2 error
	}
	beginReturnsOnCall map[int]struct {
		result1 dbng.Tx
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	DriverStub        func() driver.Driver
	driverMutex       sync.RWMutex
	driverArgsForCall []struct{}
	driverReturns     struct {
		result1 driver.Driver
	}
	driverReturnsOnCall map[int]struct {
		result1 driver.Driver
	}
	ExecStub        func(query string, args ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		query string
		args  []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 error
	}
	pingReturnsOnCall map[int]struct {
		result1 error
	}
	PrepareStub        func(query string) (*sql.Stmt, error)
	prepareMutex       sync.RWMutex
	prepareArgsForCall []struct {
		query string
	}
	prepareReturns struct {
		result1 *sql.Stmt
		result2 error
	}
	prepareReturnsOnCall map[int]struct {
		result1 *sql.Stmt
		result2 error
	}
	QueryStub        func(query string, args ...interface{}) (*sql.Rows, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		query string
		args  []interface{}
	}
	queryReturns struct {
		result1 *sql.Rows
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 *sql.Rows
		result2 error
	}
	QueryRowStub        func(query string, args ...interface{}) *sql.Row
	queryRowMutex       sync.RWMutex
	queryRowArgsForCall []struct {
		query string
		args  []interface{}
	}
	queryRowReturns struct {
		result1 *sql.Row
	}
	queryRowReturnsOnCall map[int]struct {
		result1 *sql.Row
	}
	SetMaxIdleConnsStub        func(n int)
	setMaxIdleConnsMutex       sync.RWMutex
	setMaxIdleConnsArgsForCall []struct {
		n int
	}
	SetMaxOpenConnsStub        func(n int)
	setMaxOpenConnsMutex       sync.RWMutex
	setMaxOpenConnsArgsForCall []struct {
		n int
	}
	StatsStub        func() sql.DBStats
	statsMutex       sync.RWMutex
	statsArgsForCall []struct{}
	statsReturns     struct {
		result1 sql.DBStats
	}
	statsReturnsOnCall map[int]struct {
		result1 sql.DBStats
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConn) Begin() (dbng.Tx, error) {
	fake.beginMutex.Lock()
	ret, specificReturn := fake.beginReturnsOnCall[len(fake.beginArgsForCall)]
	fake.beginArgsForCall = append(fake.beginArgsForCall, struct{}{})
	fake.recordInvocation("Begin", []interface{}{})
	fake.beginMutex.Unlock()
	if fake.BeginStub != nil {
		return fake.BeginStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.beginReturns.result1, fake.beginReturns.result2
}

func (fake *FakeConn) BeginCallCount() int {
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	return len(fake.beginArgsForCall)
}

func (fake *FakeConn) BeginReturns(result1 dbng.Tx, result2 error) {
	fake.BeginStub = nil
	fake.beginReturns = struct {
		result1 dbng.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) BeginReturnsOnCall(i int, result1 dbng.Tx, result2 error) {
	fake.BeginStub = nil
	if fake.beginReturnsOnCall == nil {
		fake.beginReturnsOnCall = make(map[int]struct {
			result1 dbng.Tx
			result2 error
		})
	}
	fake.beginReturnsOnCall[i] = struct {
		result1 dbng.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeConn) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConn) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) Driver() driver.Driver {
	fake.driverMutex.Lock()
	ret, specificReturn := fake.driverReturnsOnCall[len(fake.driverArgsForCall)]
	fake.driverArgsForCall = append(fake.driverArgsForCall, struct{}{})
	fake.recordInvocation("Driver", []interface{}{})
	fake.driverMutex.Unlock()
	if fake.DriverStub != nil {
		return fake.DriverStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.driverReturns.result1
}

func (fake *FakeConn) DriverCallCount() int {
	fake.driverMutex.RLock()
	defer fake.driverMutex.RUnlock()
	return len(fake.driverArgsForCall)
}

func (fake *FakeConn) DriverReturns(result1 driver.Driver) {
	fake.DriverStub = nil
	fake.driverReturns = struct {
		result1 driver.Driver
	}{result1}
}

func (fake *FakeConn) DriverReturnsOnCall(i int, result1 driver.Driver) {
	fake.DriverStub = nil
	if fake.driverReturnsOnCall == nil {
		fake.driverReturnsOnCall = make(map[int]struct {
			result1 driver.Driver
		})
	}
	fake.driverReturnsOnCall[i] = struct {
		result1 driver.Driver
	}{result1}
}

func (fake *FakeConn) Exec(query string, args ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Exec", []interface{}{query, args})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(query, args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.execReturns.result1, fake.execReturns.result2
}

func (fake *FakeConn) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeConn) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].query, fake.execArgsForCall[i].args
}

func (fake *FakeConn) ExecReturns(result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) ExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Ping() error {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pingReturns.result1
}

func (fake *FakeConn) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeConn) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) PingReturnsOnCall(i int, result1 error) {
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) Prepare(query string) (*sql.Stmt, error) {
	fake.prepareMutex.Lock()
	ret, specificReturn := fake.prepareReturnsOnCall[len(fake.prepareArgsForCall)]
	fake.prepareArgsForCall = append(fake.prepareArgsForCall, struct {
		query string
	}{query})
	fake.recordInvocation("Prepare", []interface{}{query})
	fake.prepareMutex.Unlock()
	if fake.PrepareStub != nil {
		return fake.PrepareStub(query)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.prepareReturns.result1, fake.prepareReturns.result2
}

func (fake *FakeConn) PrepareCallCount() int {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return len(fake.prepareArgsForCall)
}

func (fake *FakeConn) PrepareArgsForCall(i int) string {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return fake.prepareArgsForCall[i].query
}

func (fake *FakeConn) PrepareReturns(result1 *sql.Stmt, result2 error) {
	fake.PrepareStub = nil
	fake.prepareReturns = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) PrepareReturnsOnCall(i int, result1 *sql.Stmt, result2 error) {
	fake.PrepareStub = nil
	if fake.prepareReturnsOnCall == nil {
		fake.prepareReturnsOnCall = make(map[int]struct {
			result1 *sql.Stmt
			result2 error
		})
	}
	fake.prepareReturnsOnCall[i] = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Query(query string, args ...interface{}) (*sql.Rows, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Query", []interface{}{query, args})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(query, args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.queryReturns.result1, fake.queryReturns.result2
}

func (fake *FakeConn) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *FakeConn) QueryArgsForCall(i int) (string, []interface{}) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].query, fake.queryArgsForCall[i].args
}

func (fake *FakeConn) QueryReturns(result1 *sql.Rows, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) QueryReturnsOnCall(i int, result1 *sql.Rows, result2 error) {
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 *sql.Rows
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) QueryRow(query string, args ...interface{}) *sql.Row {
	fake.queryRowMutex.Lock()
	ret, specificReturn := fake.queryRowReturnsOnCall[len(fake.queryRowArgsForCall)]
	fake.queryRowArgsForCall = append(fake.queryRowArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("QueryRow", []interface{}{query, args})
	fake.queryRowMutex.Unlock()
	if fake.QueryRowStub != nil {
		return fake.QueryRowStub(query, args...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.queryRowReturns.result1
}

func (fake *FakeConn) QueryRowCallCount() int {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return len(fake.queryRowArgsForCall)
}

func (fake *FakeConn) QueryRowArgsForCall(i int) (string, []interface{}) {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return fake.queryRowArgsForCall[i].query, fake.queryRowArgsForCall[i].args
}

func (fake *FakeConn) QueryRowReturns(result1 *sql.Row) {
	fake.QueryRowStub = nil
	fake.queryRowReturns = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *FakeConn) QueryRowReturnsOnCall(i int, result1 *sql.Row) {
	fake.QueryRowStub = nil
	if fake.queryRowReturnsOnCall == nil {
		fake.queryRowReturnsOnCall = make(map[int]struct {
			result1 *sql.Row
		})
	}
	fake.queryRowReturnsOnCall[i] = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *FakeConn) SetMaxIdleConns(n int) {
	fake.setMaxIdleConnsMutex.Lock()
	fake.setMaxIdleConnsArgsForCall = append(fake.setMaxIdleConnsArgsForCall, struct {
		n int
	}{n})
	fake.recordInvocation("SetMaxIdleConns", []interface{}{n})
	fake.setMaxIdleConnsMutex.Unlock()
	if fake.SetMaxIdleConnsStub != nil {
		fake.SetMaxIdleConnsStub(n)
	}
}

func (fake *FakeConn) SetMaxIdleConnsCallCount() int {
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	return len(fake.setMaxIdleConnsArgsForCall)
}

func (fake *FakeConn) SetMaxIdleConnsArgsForCall(i int) int {
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	return fake.setMaxIdleConnsArgsForCall[i].n
}

func (fake *FakeConn) SetMaxOpenConns(n int) {
	fake.setMaxOpenConnsMutex.Lock()
	fake.setMaxOpenConnsArgsForCall = append(fake.setMaxOpenConnsArgsForCall, struct {
		n int
	}{n})
	fake.recordInvocation("SetMaxOpenConns", []interface{}{n})
	fake.setMaxOpenConnsMutex.Unlock()
	if fake.SetMaxOpenConnsStub != nil {
		fake.SetMaxOpenConnsStub(n)
	}
}

func (fake *FakeConn) SetMaxOpenConnsCallCount() int {
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	return len(fake.setMaxOpenConnsArgsForCall)
}

func (fake *FakeConn) SetMaxOpenConnsArgsForCall(i int) int {
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	return fake.setMaxOpenConnsArgsForCall[i].n
}

func (fake *FakeConn) Stats() sql.DBStats {
	fake.statsMutex.Lock()
	ret, specificReturn := fake.statsReturnsOnCall[len(fake.statsArgsForCall)]
	fake.statsArgsForCall = append(fake.statsArgsForCall, struct{}{})
	fake.recordInvocation("Stats", []interface{}{})
	fake.statsMutex.Unlock()
	if fake.StatsStub != nil {
		return fake.StatsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.statsReturns.result1
}

func (fake *FakeConn) StatsCallCount() int {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return len(fake.statsArgsForCall)
}

func (fake *FakeConn) StatsReturns(result1 sql.DBStats) {
	fake.StatsStub = nil
	fake.statsReturns = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeConn) StatsReturnsOnCall(i int, result1 sql.DBStats) {
	fake.StatsStub = nil
	if fake.statsReturnsOnCall == nil {
		fake.statsReturnsOnCall = make(map[int]struct {
			result1 sql.DBStats
		})
	}
	fake.statsReturnsOnCall[i] = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeConn) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.driverMutex.RLock()
	defer fake.driverMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConn) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ dbng.Conn = new(FakeConn)
