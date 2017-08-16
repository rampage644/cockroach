// Copyright 2017 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package sql

import (
	"bytes"

	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgerror"
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
)

// CreateFunction creates a persistent function.
// Privileges: None.
func (p *planner) CreateFunction(ctx context.Context, n *parser.CreateFunction) (planNode, error) {
	return nil, pgerror.Unimplemented(`create-function`, "CREATE FUNCTION is not implemented")
}

type createProcedureNode struct {
	n *parser.CreateProcedure
}

func (n *createProcedureNode) Start(params runParams) error {
	const insertStoredProcedure = `INSERT INTO system.proc ("name", "body") ` +
		`VALUES ($1, $2)`
	var buf bytes.Buffer
	for idx, stmt := range n.n.Body {
		if idx > 0 {
			buf.WriteString("; ")
		}
		stmt.Format(&buf, parser.FmtSimple)
	}

	rows, err := params.p.exec(params.ctx, insertStoredProcedure, n.n.Name, buf.String())
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.Errorf("%s: expected 1 result, found %d", insertStoredProcedure, rows)
	}
	return nil
}

func (*createProcedureNode) Next(runParams) (bool, error) { return false, nil }
func (*createProcedureNode) Close(context.Context)        {}
func (*createProcedureNode) Values() parser.Datums        { return parser.Datums{} }

// CreateProcedure creates a persistent stored procedure.
// Privileges: None.
func (p *planner) CreateProcedure(ctx context.Context, n *parser.CreateProcedure) (planNode, error) {
	return &createProcedureNode{
		n: n,
	}, nil
}

type callProcedureNode struct {
	n *parser.CallProcedure
}

func (p *planner) CallProcedure(ctx context.Context, n *parser.CallProcedure, desiredTypes []parser.Type) (planNode, error) {
	panic("Shouldn't be called ever")
}

func (n *callProcedureNode) Start(params runParams) error {
	panic("Shouldn't be called ever")
}

func (*callProcedureNode) Next(runParams) (bool, error) {
	panic("Shouldn't be called ever")
	return false, nil
}
func (*callProcedureNode) Close(context.Context) {
	panic("Shouldn't be called ever")
}
func (*callProcedureNode) Values() parser.Datums {
	panic("Shouldn't be called ever")
	return parser.Datums{}
}

func SaveStoreProcedure(node *parser.CreateProcedure) error {

	return pgerror.Unimplemented(`create-function`, "save store procedure is not implemented")
}
