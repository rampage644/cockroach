package sql

import (
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

func (e *Engine) ExecuteProcedure(name string) (res Result, err error) {
	err = errors.Errorf("execute procedure is not implemented yet")
	return res, err
}
