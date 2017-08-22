package sql

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"golang.org/x/net/context"
)

type ConditionalNode struct {
}

func (p *planner) Conditional(ctx context.Context, node *Conditional) (planNode, error) {
	return nil, fmt.Errorf("we're not quite frobnicating yet...")
}

func (n *ConditionalNode) Start(params runParams) error {
	return fmt.Errorf("not implemented yet")
}

func (n *ConditionalNode) Next(params runParams) (bool, error) {
	return fmt.Errorf("not implemented yet")
}

func (n *ConditionalNode) Values() parser.Datums {
	return fmt.Errorf("not implemented yet")
}

func (n *ConditionalNode) Close(ctx context.Context) {
	return fmt.Errorf("not implemented yet")
}
