package sqlwriter

import (
	"fmt"
	"strings"
)

type InsertStatement struct {
	tableName       string
	columns         []string
	values          []interface{}
	returningColumn string
}

func Insert(tableName string, columns ...string) *InsertStatement {
	i := &InsertStatement{
		tableName: tableName,
		columns:   columns,
	}

	return i
}

func (i *InsertStatement) Columns(columnNames ...string) *InsertStatement {
	i.columns = append(i.columns, columnNames...)
	return i
}

func (i *InsertStatement) Col(columnName string, value interface{}) *InsertStatement {
	i.columns = append(i.columns, columnName)
	i.values = append(i.values, value)
	return i
}

func (i *InsertStatement) Values(values ...interface{}) *InsertStatement {
	i.values = append(i.values, values...)
	return i
}

func (i *InsertStatement) ColumnValues(columnValuePairs ...interface{}) *InsertStatement {

	for k := 0; k < len(columnValuePairs); k += 2 {
		// If this blows up with type assertion error, then the caller is not
		// passing in strings for column names, or got their name/value pairs
		// mixed up.
		i.columns = append(i.columns, columnValuePairs[k].(string))
		// If this blows up with "index out of range" then it means somebody is
		// calling this method with an uneven number of arguments, i.e. not
		// pairs.
		i.values = append(i.values, columnValuePairs[k+1])
	}

	return i
}

func (i *InsertStatement) Returning(columnName string) *InsertStatement {
	i.returningColumn = columnName
	return i
}

func (i *InsertStatement) bindMarkers() []string {
	// For postgres this is $1, $2, ...
	// It may be different for other databases.
	markers := []string{}
	for p := range i.columns {
		markers = append(markers, fmt.Sprintf("$%d", p+1))
	}
	return markers
}

func (i *InsertStatement) SQL() string {
	s := fmt.Sprintf("insert into %s (%s) values (%s)",
		i.tableName,
		strings.Join(i.columns, ", "),
		strings.Join(i.bindMarkers(), ", "))

	if i.returningColumn != "" {
		s = s + " returning " + i.returningColumn
	}

	return s
}

func (i *InsertStatement) BindValues() []interface{} {
	return i.values
}
