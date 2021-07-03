package model

import (
	"github.com/pingcap/parser/ast"
)

type tableX struct {
	tableNames []string
}

func (v *tableX) Enter(in ast.Node) (ast.Node, bool) {
	if name, ok := in.(*ast.TableName); ok {
		v.tableNames = append(v.tableNames, name.Name.O)
	}
	return in, false
}

func (v *tableX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func ExtractTableName(rootNode *ast.StmtNode) []string {
	if _, ok := (*rootNode).(*ast.SelectStmt); !ok {
		return nil
	}
	v := &tableX{}
	(*rootNode).Accept(v)
	return v.tableNames
}
