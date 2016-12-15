package subject

import (
	"testing"
	"reflect"
)

func TestLoadSubject(t *testing.T) {
	path := "../../example/subject.csv"
	expect := [][]string{{"A01", "Shibuya Rin", "100"}, {"A02", "", "200"}}

	res, err := LoadSubject(path)

	if !(reflect.DeepEqual(res, expect) && err == nil) {
		t.Errorf("Failed to load and parse csv. result: %s, error: %s", res, err)
	}
}

func TestLoadSubject_LoadError(t *testing.T) {
	path := "/path/to/nowhere.csv"

	res, err := LoadSubject(path)

	if !(res == nil && err != nil) {
		t.Errorf("Failed to raise error properly. result: %s, error: %s", res, err)
	}
}

func TestLoadSubject_ParseError(t *testing.T) {
	path := "/tmp/"

	res, err := LoadSubject(path)

	if !(res == nil && err != nil) {
		t.Errorf("Failed to raise error properly. result: %s, error: %s", res, err)
	}
}
