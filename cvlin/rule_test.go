package cvlin

import (
	"testing"
	"reflect"
)

func TestLoadRule(t *testing.T) {
	path := "../example/rule.toml"
	expect := map[string]rule {
		"id"    : {Pattern: "A0[0-9]",  Notnull: false},
		"name"  : {Pattern: "*",        Notnull: true},
		"point" : {Pattern: "^[0-9]+$", Notnull: true},
	}

	res, err := LoadRule(path)

	if !(err == nil) {
		t.Errorf("Failed to parse toml (some error occurs). %s", err)
	}

	// This always returns false. Why?
	// reflect.DeepEqual(res, expect)

	if !(len(expect) == len(res)) {
		t.Errorf("Failed to parse toml (length differs). Expected: %s Got: %s", len(expect),len(res))
	}

	for k := range res {
		if !reflect.DeepEqual(expect[k], res[k]) {
			t.Errorf("Failed to parse toml (content differs). Expected: %s Got: %s", expect[k], res[k])
		}
	}
}


func TestLoadRule_LoadError(t *testing.T) {
	path := "/path/to/nowhere.toml"

	res, err := LoadRule(path)

	if !(res == nil && err != nil) {
		t.Errorf("Failed to raise error properly. result: %s, error: %s", res, err)
	}
}
