package cvlin

import (
	"testing"
	"reflect"
)

func TestParse(t *testing.T) {
	path := "../example/subject.csv"
	expect := [][]string{{"A01", "Shibuya Rin", "100"}, {"A02", "Sakuma Mayu", "200"}}

	res, err := Parse(path)

	if !(reflect.DeepEqual(res, expect) && err == nil) {
		t.Errorf("Failed to load and parse csv. result: %s, error: %s", res, err)
	}
}

func TestParse_LoadError(t *testing.T) {
	path := "/path/to/nowhere/test.csv"

	res, err := Parse(path)

	if !(res == nil && err != nil) {
		t.Errorf("Failed to raise error properly. result: %s, error: %s", res, err)
	}
}

func TestParse_ParseError(t *testing.T) {
	path := "/tmp/"

	res, err := Parse(path)

	if !(res == nil && err != nil) {
		t.Errorf("Failed to raise error properly. result: %s, error: %s", res, err)
	}
}
