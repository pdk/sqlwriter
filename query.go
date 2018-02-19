package sqlwriter

import (
	"fmt"
	"strconv"
	"strings"
)

type tableJoin struct {
	joinType   string
	tableName  string
	conditions []string
}

type Query struct {
	bindCounter int
	distinctOn  []string
	selects     []string
	tableName   string
	wheres      []string
	joins       []tableJoin
	orderBys    []string
	groupBys    []string
	limit       int
	bindValues  []interface{}
}

func (q *Query) DistinctOn(columns ...string) *Query {
	q.distinctOn = append(q.distinctOn, columns...)
	return q
}

func Select(expressions ...string) *Query {
	q := &Query{
		limit: -1,
	}
	return q.Select(expressions...)
}

func (q *Query) Select(expressions ...string) *Query {
	q.selects = append(q.selects, expressions...)
	return q
}

func (q *Query) From(tableName string) *Query {
	q.tableName = tableName
	return q
}

func (q *Query) Where(expressions ...string) *Query {
	q.wheres = append(q.wheres, expressions...)
	return q
}

func (q *Query) Equal(columnName string, value interface{}) *Query {

	return q.Expr(columnName, Equals, value)
}

type Operator string

const (
	Equals              Operator = "="
	NotEqual            Operator = "<>"
	GreaterThan         Operator = ">"
	GreaterThanOrEquals Operator = ">="
	LessThan            Operator = "<"
	LessThanOrEquals    Operator = "<="
)

func (q *Query) Expr(columnName string, operator Operator, value interface{}) *Query {
	q.bindCounter++

	q.wheres = append(q.wheres,
		fmt.Sprintf("%s %s $%d", columnName, operator, q.bindCounter))
	q.bindValues = append(q.bindValues, value)

	return q
}

func (q *Query) Between(columnName string, value1, value2 interface{}) *Query {
	q.bindCounter += 2
	q.wheres = append(q.wheres,
		fmt.Sprintf("%s between $%d and $%d", columnName, q.bindCounter-1, q.bindCounter))
	q.bindValues = append(q.bindValues, value1, value2)
	return q
}

func (q *Query) InnerJoin(tableName string, conditions ...string) *Query {
	q.joins = append(q.joins, tableJoin{"inner", tableName, conditions})
	return q
}

func (q *Query) OuterJoin(tableName string, conditions ...string) *Query {
	q.joins = append(q.joins, tableJoin{"left outer", tableName, conditions})
	return q
}

func (q *Query) GroupBy(expressions ...string) *Query {
	q.groupBys = append(q.groupBys, expressions...)
	return q
}

func (q *Query) OrderBy(expressions ...string) *Query {
	q.orderBys = append(q.orderBys, expressions...)
	return q
}

func (q *Query) Limit(limit int) *Query {
	q.limit = limit
	return q
}

func (q *Query) SQL() string {
	queryString := "select"

	if len(q.distinctOn) > 0 {
		queryString += " distinct on (" + strings.Join(q.distinctOn, ", ") + ")"
	}

	if len(q.selects) == 0 {
		queryString += " *"
	} else {
		queryString += "\n    " + strings.Join(q.selects, ",\n    ")
	}

	queryString += "\nfrom " + q.tableName

	if len(q.joins) > 0 {
		for _, j := range q.joins {
			queryString += "\n" + j.joinType + " join " + j.tableName +
				" on " + strings.Join(j.conditions, "\nand ")
		}
	}

	if len(q.wheres) > 0 {
		queryString += "\nwhere " + strings.Join(q.wheres, "\nand ")
	}

	if len(q.orderBys) > 0 {
		queryString += "\norder by\n    " + strings.Join(q.orderBys, ",\n    ")
	}

	if len(q.groupBys) > 0 {
		queryString += "\ngroup by\n    " + strings.Join(q.groupBys, ",\n    ")
	}

	if q.limit > -1 {
		queryString += "\nlimit " + strconv.Itoa(q.limit)
	}

	return queryString
}

func (q *Query) String() string {
	return q.SQL()
}

func (q *Query) BindValues() []interface{} {
	return q.bindValues
}
