package cvlin

import (
	"testing"
	"reflect"
)

func TestLoadRule(t *testing.T) {
	path := "../example/rule.toml"
	expect := []rule {
		{Pattern: "A0[0-9]",  Notnull: false},
		{Pattern: "*",        Notnull: true},
		{Pattern: "^[0-9]+$", Notnull: true},
	}

	result, err := LoadRule(path)

	if err != nil {
		t.Errorf("Failed to parse toml (unexpected error occurs): %s", err)
	}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Failed to parse toml (content differs). Expected: %s Got: %s", expect, result)
	}
}


func TestLoadRule_LoadError(t *testing.T) {
	path := "/path/to/nowhere.toml"

	result, err := LoadRule(path)

	if !(result == nil && err != nil) {
		t.Errorf("Failed to raise error properly. result: %s, error: %s", result, err)
	}
}


func TestSort(t *testing.T) {
	order := []string{"key1", "key2", "key3"}
	target := map[string]rule {
		"key3" : {Pattern: "3"},
		"key2" : {Pattern: "2"},
		"key1" : {Pattern: "1"},
	}

	expect := []rule{{Pattern: "1"}, {Pattern: "2"}, {Pattern: "3"}}
	result := sort(target, order)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Failed to sort. Expected: %s Got: %s", expect, result)
	}
}
