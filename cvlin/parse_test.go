package cvlin

import (
	"testing"
	"reflect"
)

func TestParse(t *testing.T) {
	path := "../example/rule.csv"
	expect := [][]string{{"aaa", "000"}, {"bbb", "111"}}

	res, err := Parse(path)

	if !(reflect.DeepEqual(res, expect) && err == nil) {
		t.Errorf("Failed to load and parse csv. result: %d, error: %s", res, err)
	}
}

func TestParse_LoadError(t *testing.T) {
	path := "/path/to/nowhere/test.csv"

	res, err := Parse(path)

	if !(res == nil && err != nil) {
		t.Errorf("Failed to raise error properly. result: %d, error: %s", res, err)
	}
}

func TestParse_ParseError(t *testing.T) {
	path := "/tmp/"

	res, err := Parse(path)

	if !(res == nil && err != nil) {
		t.Errorf("Failed to raise error properly. result: %d, error: %s", res, err)
	}
}
