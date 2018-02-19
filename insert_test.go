package sqlwriter

import (
	"testing"
)

func TestInsert(t *testing.T) {

	i := Insert("users", "id", "name").
		Values(12, "Alice").
		SQL()

	e := "insert into users (id, name) values ($1, $2)"

	if i != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, i)
	}
}

func TestReturning(t *testing.T) {

	i := Insert("users", "name").
		Values("Alice").
		Returning("id").
		SQL()

	e := "insert into users (name) values ($1) returning id"

	if i != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, i)
	}
}

func TestColumns(t *testing.T) {

	i := Insert("users").
		Columns("id", "name").
		Values(12, "Alice").
		SQL()

	e := "insert into users (id, name) values ($1, $2)"

	if i != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, i)
	}
}

func TestCol(t *testing.T) {

	i := Insert("users").
		Col("id", 12).
		Col("name", "alice").
		SQL()

	e := "insert into users (id, name) values ($1, $2)"

	if i != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, i)
	}
}

func TestBindValues(t *testing.T) {
	v := Insert("users").
		Col("id", 12).
		Col("name", "alice").
		BindValues()

	if len(v) != 2 {
		t.Errorf("expected 2 values, but got %d", len(v))

	} else {

		v1, ok := (v[0]).(int)
		if !ok {
			t.Errorf("expected first value to be an integer, but got %T", v[0])
		}
		if v1 != 12 {
			t.Errorf("expected first value to be 12, but got %d", v1)
		}

		v2, ok := (v[1]).(string)
		if !ok {
			t.Errorf("expected first value to be an integer, but got %T", v[1])
		}
		if v2 != "alice" {
			t.Errorf("expected first value to be 12, but got %s", v2)
		}

	}
}

func TestWithoutValues(t *testing.T) {
	i := Insert("users", "id", "name").SQL()

	e := "insert into users (id, name) values ($1, $2)"

	if i != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, i)
	}
}
