package sqlwriter

import "testing"

func TestStatement(t *testing.T) {
	sql := "select ack from foo where color = $1"

	s := NewStatement(sql, "blue")

	if s.SQL() != sql {
		t.Errorf("expected sql %s but got %s", sql, s.SQL())
	}

	if len(s.BindValues()) != 1 {
		t.Errorf("expected 1 bind value, but got %d", len(s.BindValues()))
	}
}
