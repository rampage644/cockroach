package sql

import (
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

	var sl parser.StatementList
	body, _ := rows[0].Next()
	sbody, _ := body.(*parser.DString)
	sqlStr := string(*sbody)
	sl, err = parser.Parse(sqlStr[:len(sqlStr)-1])
	stmts := NewStatementList(sl)

	if err != nil {
		return errors.Wrap(err, "call-procedure error")
	}

	for _, stmt := range stmts {
		err := e.executor.execRegularStmtInOpenTxn(e.session, stmt, nil, false, false, false, 0, statementResultWriter)
		if err != nil {
			return errors.Wrap(err, "call-procedure error")
		}
	}

	return nil
}
