package query

import (
	"fmt"
	"strings"
)

func SelectQuery(table string, columns []string, condition string, limit int, offset int) string {
	selectClause := "SELECT " + strings.Join(columns, ", ")
	fromClause := "FROM " + table
	whereClause := ""
	limitClause := ""
	offsetClause := ""

	if condition != "" {
		whereClause = "WHERE " + condition
	}

	if limit > 0 {
		limitClause = fmt.Sprintf("LIMIT %d", limit)
	}

	if offset > 0 {
		offsetClause = fmt.Sprintf("OFFSET %d", offset)
	}

	clauses := []string{selectClause, fromClause, whereClause, limitClause, offsetClause}

	return strings.Join(clauses, " ")
}

func DeleteQuery(table, condition string) string {
	deleteClause := "DELETE"
	fromClause := "FROM " + table

	whereClause := ""
	if condition != "" {
		whereClause = "WHERE " + condition
	}

	return fmt.Sprintf("%s %s %s", deleteClause, fromClause, whereClause)
}
