package sqlwriter_test

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	q "github.com/pdk/sqlwriter"
)

type ColStuff struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
}

type Member struct {
	ColStuff
	Age     int `db:"age"`
	Name    string
	Address sql.NullString `db:"street_address"`
}

func TestDBNames(t *testing.T) {

	names := q.DBNames(ColStuff{}, Member{})

	if len(names) != 5 {
		t.Errorf("expected 5 names, got %d", len(names))
	}

	expected := []string{
		"id",
		"created_at",
		"age",
		"name",
		"street_address",
	}

	if !reflect.DeepEqual(names, expected) {
		t.Errorf("expected %s, but got %s", expected, names)
	}
}
