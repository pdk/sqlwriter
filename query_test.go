package sqlwriter

import (
	"database/sql"
	"testing"
)

func TestSelect(t *testing.T) {
	q := Select("a", "b", "c").
		From("the_table").
		String()

	e := "select\n" +
		"    a,\n" +
		"    b,\n" +
		"    c\n" +
		"from the_table"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

func TestWhere(t *testing.T) {
	q := Select("*").From("boink").
		Where("flavor = 'grape'",
			"color = 'purple'").
		String()

	e := "select\n" +
		"    *\n" +
		"from boink\n" +
		"where flavor = 'grape'\n" +
		"and color = 'purple'"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

func TestInnerJoin(t *testing.T) {
	q := Select("ack").
		From("account").
		InnerJoin("user", "user.account_id = account.account_id",
			"user.active = true").
		String()

	e := "select\n" +
		"    ack\n" +
		"from account\n" +
		"inner join user on user.account_id = account.account_id\n" +
		"and user.active = true"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

func TestOuterJoin(t *testing.T) {
	q := Select("ack").
		From("account").
		OuterJoin("user", "user.account_id = account.account_id",
			"user.active = true").
		String()

	e := "select\n" +
		"    ack\n" +
		"from account\n" +
		"left outer join user on user.account_id = account.account_id\n" +
		"and user.active = true"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

func TestGroupBy(t *testing.T) {
	q := Select("ack", "blip", "sum(glorp) as blift").
		From("foofah").
		GroupBy("ack", "blip").
		String()

	e := "select\n" +
		"    ack,\n" +
		"    blip,\n" +
		"    sum(glorp) as blift\n" +
		"from foofah\n" +
		"group by\n" +
		"    ack,\n" +
		"    blip"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

func TestEqual(t *testing.T) {
	q := Select("ack").
		From("blip").
		Equal("id", 1)

	e := "select\n" +
		"    ack\n" +
		"from blip\n" +
		"where id = $1"

	sql := q.SQL()

	if sql != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}

	vals := q.BindValues()

	if len(vals) != 1 {
		t.Errorf("expected 1 bind value, but got %d", len(vals))
	}

	val := vals[0]
	one, ok := val.(int)

	if !ok {
		t.Errorf("expected to get int bind value, but got %T", val)
	}

	if one != 1 {
		t.Errorf("expected to find value 1, but got %d", one)
	}
}

func TestExpr(t *testing.T) {
	q := Select("ack").
		From("blip").
		Expr("dollars", GreaterThanOrEquals, "1000000")

	sql := q.SQL()

	e := "select\n" +
		"    ack\n" +
		"from blip\n" +
		"where dollars >= $1"

	if sql != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

func TestBetween(t *testing.T) {
	q := Select("ack").
		From("blip").
		Between("dollars", 1000, 1000000)

	sql := q.SQL()

	e := "select\n" +
		"    ack\n" +
		"from blip\n" +
		"where dollars between $1 and $2"

	if sql != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}

	vals := q.BindValues()

	if len(vals) != 2 {
		t.Errorf("expected 2 bind values, but got %d", len(vals))
	}

}

func TestLimit(t *testing.T) {
	q := Select().From("foo").Limit(42).SQL()

	e := "select *\n" +
		"from foo\n" +
		"limit 42"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

func TestOrderBy(t *testing.T) {
	q := Select().From("foo").OrderBy("ack", "blip").SQL()

	e := "select *\n" +
		"from foo\n" +
		"order by\n" +
		"    ack,\n" +
		"    blip"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

func TestDistictOn(t *testing.T) {
	q := Select().DistinctOn("alpha", "beta").From("foo").SQL()

	e := "select distinct on (alpha, beta) *\n" +
		"from foo"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}

type Record struct {
	Name    string
	Age     int
	Address sql.NullString `db:"address"`
}

func TestSelectStruct(t *testing.T) {
	q := Select(DBNames(Record{})...).
		From("latable").SQL()

	e := "select\n" +
		"    name,\n" +
		"    age,\n" +
		"    address\n" +
		"from latable"

	if q != e {
		t.Errorf("expected <<%s>>, got <<%s>>", e, q)
	}
}
