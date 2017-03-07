// This file was generated by counterfeiter
package migrationsfakes

import (
	"database/sql"
	"sync"

	"github.com/concourse/atc/db/migrations"
)

type FakeLimitedTx struct {
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
	StmtStub        func(stmt *sql.Stmt) *sql.Stmt
	stmtMutex       sync.RWMutex
	stmtArgsForCall []struct {
		stmt *sql.Stmt
	}
	stmtReturns struct {
		result1 *sql.Stmt
	}
	stmtReturnsOnCall map[int]struct {
		result1 *sql.Stmt
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLimitedTx) Exec(query string, args ...interface{}) (sql.Result, error) {
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

func (fake *FakeLimitedTx) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeLimitedTx) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].query, fake.execArgsForCall[i].args
}

func (fake *FakeLimitedTx) ExecReturns(result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeLimitedTx) ExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
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

func (fake *FakeLimitedTx) Prepare(query string) (*sql.Stmt, error) {
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

func (fake *FakeLimitedTx) PrepareCallCount() int {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return len(fake.prepareArgsForCall)
}

func (fake *FakeLimitedTx) PrepareArgsForCall(i int) string {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return fake.prepareArgsForCall[i].query
}

func (fake *FakeLimitedTx) PrepareReturns(result1 *sql.Stmt, result2 error) {
	fake.PrepareStub = nil
	fake.prepareReturns = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeLimitedTx) PrepareReturnsOnCall(i int, result1 *sql.Stmt, result2 error) {
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

func (fake *FakeLimitedTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
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

func (fake *FakeLimitedTx) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *FakeLimitedTx) QueryArgsForCall(i int) (string, []interface{}) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].query, fake.queryArgsForCall[i].args
}

func (fake *FakeLimitedTx) QueryReturns(result1 *sql.Rows, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeLimitedTx) QueryReturnsOnCall(i int, result1 *sql.Rows, result2 error) {
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

func (fake *FakeLimitedTx) QueryRow(query string, args ...interface{}) *sql.Row {
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

func (fake *FakeLimitedTx) QueryRowCallCount() int {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return len(fake.queryRowArgsForCall)
}

func (fake *FakeLimitedTx) QueryRowArgsForCall(i int) (string, []interface{}) {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return fake.queryRowArgsForCall[i].query, fake.queryRowArgsForCall[i].args
}

func (fake *FakeLimitedTx) QueryRowReturns(result1 *sql.Row) {
	fake.QueryRowStub = nil
	fake.queryRowReturns = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *FakeLimitedTx) QueryRowReturnsOnCall(i int, result1 *sql.Row) {
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

func (fake *FakeLimitedTx) Stmt(stmt *sql.Stmt) *sql.Stmt {
	fake.stmtMutex.Lock()
	ret, specificReturn := fake.stmtReturnsOnCall[len(fake.stmtArgsForCall)]
	fake.stmtArgsForCall = append(fake.stmtArgsForCall, struct {
		stmt *sql.Stmt
	}{stmt})
	fake.recordInvocation("Stmt", []interface{}{stmt})
	fake.stmtMutex.Unlock()
	if fake.StmtStub != nil {
		return fake.StmtStub(stmt)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stmtReturns.result1
}

func (fake *FakeLimitedTx) StmtCallCount() int {
	fake.stmtMutex.RLock()
	defer fake.stmtMutex.RUnlock()
	return len(fake.stmtArgsForCall)
}

func (fake *FakeLimitedTx) StmtArgsForCall(i int) *sql.Stmt {
	fake.stmtMutex.RLock()
	defer fake.stmtMutex.RUnlock()
	return fake.stmtArgsForCall[i].stmt
}

func (fake *FakeLimitedTx) StmtReturns(result1 *sql.Stmt) {
	fake.StmtStub = nil
	fake.stmtReturns = struct {
		result1 *sql.Stmt
	}{result1}
}

func (fake *FakeLimitedTx) StmtReturnsOnCall(i int, result1 *sql.Stmt) {
	fake.StmtStub = nil
	if fake.stmtReturnsOnCall == nil {
		fake.stmtReturnsOnCall = make(map[int]struct {
			result1 *sql.Stmt
		})
	}
	fake.stmtReturnsOnCall[i] = struct {
		result1 *sql.Stmt
	}{result1}
}

func (fake *FakeLimitedTx) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	fake.stmtMutex.RLock()
	defer fake.stmtMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLimitedTx) recordInvocation(key string, args []interface{}) {
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

var _ migrations.LimitedTx = new(FakeLimitedTx)
