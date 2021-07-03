package model

import (
	"fmt"
	"github.com/pingcap/parser/ast"
)

type colX struct {
	colNames []string
}

func (v *colX) Enter(in ast.Node) (ast.Node, bool) {
	if boExpr, ok := in.(*ast.BinaryOperationExpr); ok {

		switch boExpr.Op.String() {
		case "eq":
			if leftExpr, ok := boExpr.L.(*ast.ColumnNameExpr); ok {
				v.colNames = append([]string{leftExpr.Name.Name.O}, v.colNames...)
			} else if rightExpr, ok := boExpr.R.(*ast.ColumnNameExpr); ok {
				v.colNames = append([]string{rightExpr.Name.Name.O}, v.colNames...)

			}
		case "gt", "ge", "lt", "le":
			if leftExpr, ok := boExpr.L.(*ast.ColumnNameExpr); ok {
				v.colNames = append(v.colNames, leftExpr.Name.Name.O)
			} else if rightExpr, ok := boExpr.R.(*ast.ColumnNameExpr); ok {
				v.colNames = append([]string{rightExpr.Name.Name.O}, v.colNames...)
			}
		default:
			fmt.Println("unsupported expressions: ", boExpr.Op.String())
		}
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
