package sqlwriter

import (
	"reflect"
	"strings"
)

// DBNames returns the names of database columns of a struct. Names are either
// tagged with db:... or just the lowercase field name. Composed structs will be
// skipped (unless they are tagged ala `db:"blah"`), so that composed structs
// can be used. Will panic if passed a non-struct.
func DBNames(structList ...interface{}) []string {
	names := []string{}

	for _, thing := range structList {
		t := reflect.TypeOf(thing)

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			n := f.Tag.Get("db")
			if n != "" {
				names = append(names, n)
				continue
			}

			if reflect.ValueOf(thing).Field(i).Kind() == reflect.Struct {
				continue
			}

			names = append(names, strings.ToLower(f.Name))
		}
	}
	return names
}
