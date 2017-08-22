package parser

import (
	"bytes"
)

type Conditional struct {
}

// Implements Statement interface
func (node *Conditional) StatementType() StatementType { return Unknown }
func (node *Conditional) StatementTag() string         { return "Conditional" }

func (node *Conditional) Format(buf *bytes.Buffer, f FmtFlags) {
	buf.WriteString("IF THEN ELSE")
}

func (node *Conditional) String() string {
	return AsString(node)
}
