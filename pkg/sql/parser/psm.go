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

package parser

import "bytes"

// CreateFunction represents a CREATE FUNCTION statement.
type CreateFunction struct {
}

// Format implements the NodeFormatter interface.
func (node *CreateFunction) Format(buf *bytes.Buffer, f FmtFlags) {
	buf.WriteString("CREATE FUNCTION ")
}

type CallProcedure struct {
	Name string
}

func (node *CallProcedure) Format(buf *bytes.Buffer, f FmtFlags) {
	buf.WriteString("CALL ")
	buf.WriteString(node.Name)
	buf.WriteString("()")
}

// CreateProcedure represents a CREATE PROCEDURE statement.
type CreateProcedure struct {
	Name       string
	Parameters ParameterList
	Body       []Statement
}

type Parameter struct {
	Type string
	Name string
}

type ParameterList []Parameter

// Format implements the NodeFormatter interface.
func (node *CreateProcedure) Format(buf *bytes.Buffer, f FmtFlags) {
	buf.WriteString("CREATE PROCEDURE ")
	buf.WriteString(node.Name)
	buf.WriteString("(")
	for idx, param := range node.Parameters {
		if idx > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.Type)
		buf.WriteString(" ")
		buf.WriteString(param.Name)
	}
	buf.WriteString(") { ")
	for idx, stmt := range node.Body {
		if idx > 0 {
			buf.WriteString("; ")
		}
		stmt.Format(buf, f)
	}
	buf.WriteString(" }")
}
