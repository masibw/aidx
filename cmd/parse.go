package cmd

import (
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/parser/test_driver"
)

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}
