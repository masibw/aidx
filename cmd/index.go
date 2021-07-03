package cmd

import (
	"aidx/model"
	"fmt"
	"strings"
)

func GenIndex(query string) (idx string) {
	astNode, err := parse(query)
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	cols := model.ExtractColumn(astNode)
	idxCols := strings.Join(cols, ", ")
	tableNames := model.ExtractTableName(astNode)
	idx = fmt.Sprintf("CREATE INDEX idx ON %s(%s);", tableNames[0], idxCols)
	return
}
