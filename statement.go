package sqlwriter

type Statement interface {
	SQL() string
	BindValues() []interface{}
}

// RawStatement type is provided as a base/last resort type to be able to ad-hoc
// make up whatever.
type RawStatement struct {
	sqlStatement string
	bindValues   []interface{}
}

func NewStatement(sqlStatement string, bindValues ...interface{}) *RawStatement {
	return &RawStatement{
		sqlStatement,
		bindValues,
	}
}

func (s *RawStatement) SQL() string {
	return s.sqlStatement
}

func (s *RawStatement) BindValues() []interface{} {
	return s.bindValues
}
