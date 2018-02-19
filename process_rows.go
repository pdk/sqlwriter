package sqlwriter

import (
	"database/sql"

	"github.com/pkg/errors"
)

type RowHandler func(*sql.Rows) error

func ProcessRows(db *sql.DB, query Statement, scanner RowHandler) (int, error) {
	rows, err := db.Query(query.SQL(), query.BindValues()...)
	if err != nil {
		return 0, errors.Wrapf(err, "query failed %s", query.SQL())
	}

	counter := 0
	for rows.Next() {
		err = scanner(rows)
		if err != nil {
			return counter, errors.Wrapf(err, "scanner failed for query %s", query.SQL())
		}
		counter++
	}

	return counter, nil
}
