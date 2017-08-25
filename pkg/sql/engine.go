package sql

import (
	_ "net/http/pprof"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"github.com/pkg/errors"
)

type Engine struct {
	executor *Executor
	session  *Session
}

func MakeEngine(executor *Executor, sess *Session) *Engine {
	return &Engine{
		executor: executor,
		session:  sess,
	}
}

func (e *Engine) ExecuteProcedure(name string, statementResultWriter StatementResultWriter) error {
	ex := InternalExecutor{LeaseManager: e.executor.cfg.LeaseManager}
	rows, err := ex.QueryRowInTransaction(
		e.session.Ctx(), "call-procedure", e.session.TxnState.mu.txn,
		"SELECT body FROM system.proc WHERE name = $1", name)

	if err != nil {
		return errors.Wrap(err, "call-procedure error")
	}

	body, _ := rows[0].Next()
	sbody, _ := body.(*parser.DString)
	stmts := string(*sbody)
	stmts = stmts[:len(stmts)-1]

	err = e.executor.execRequest(e.session, stmts, nil, copyMsgNone)

	return err
}
