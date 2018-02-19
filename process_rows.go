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

// The following is a (non-tested) example of how to use the ProcessRows
// function, above.

type recordType struct {
	Name    string
	Address string
}

func queryRecords(db *sql.DB) ([]recordType, error) {

	var accumulator []recordType
	scanner := func(r *sql.Rows) error {
		nextRecord := recordType{}
		err := r.Scan(&nextRecord.Name, &nextRecord.Address)
		accumulator = append(accumulator, nextRecord)
		return err
	}

	query := Select(DBNames(recordType{})...).
		From("latable").
		Expr("name", GreaterThanOrEquals, "Bobby")

	_, err := ProcessRows(db, query, scanner)

	return accumulator, err
}
