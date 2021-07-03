package model

import (
	"github.com/pingcap/parser/ast"
)

type colX struct {
	colNames []string
}

func (v *colX) Enter(in ast.Node) (ast.Node, bool) {
	if name, ok := in.(*ast.ColumnName); ok {
		v.colNames = append(v.colNames, name.Name.O)
	}
	return in, false
}

func (v *colX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func ExtractColumn(rootNode *ast.StmtNode) []string {
	if _, ok := (*rootNode).(*ast.SelectStmt); !ok {
		return nil
	}
	v := &colX{}
	(*rootNode).Accept(v)
	return v.colNames
}
