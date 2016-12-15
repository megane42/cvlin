package rule

import (
	"testing"
	"reflect"
	"github.com/BurntSushi/toml"
)

func TestLoadRule(t *testing.T) {
	path := "../../example/rule.toml"
	expect := []string {`A0[0-9]`, `.*`, `^\d+$`}

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


func TestConvertToOrderedSlice(t *testing.T) {
	rulemap := map[string]string {
		"ccc": "333",
		"bbb": "222",
		"aaa": "111",
	}
	keys := []toml.Key {
		{"aaa"},
		{"bbb"},
		{"ccc"},
	}
	expect := []string {"111", "222", "333"} // The order is important

	result := convertToOrderedSlice(rulemap, keys)

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Failed to convert map to slice. Expected: %v Got: %v", expect, result)
	}
}
