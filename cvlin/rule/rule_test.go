package rule

import (
	"regexp"
	"testing"
	"reflect"
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

func TestLoadRule(t *testing.T) {
	tomlBytes, _ := ioutil.ReadFile("../../example/rule.toml")
	expect := []*regexp.Regexp {
		regexp.MustCompile(`A0[0-9]`),
		regexp.MustCompile(`.*`),
		regexp.MustCompile(`^\d+$`),
	}

	result, err := LoadRule(string(tomlBytes))

	if err != nil {
		t.Errorf("Failed to parse toml (unexpected error occurs): %s", err)
	}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Failed to parse toml (content differs). Expected: %s Got: %s", expect, result)
	}
}


func TestLoadRuleWithInvalidRegexp(t *testing.T) {
	toml := `
    validRule1  = "A0[0-9]"
    invalidRule = "*"
    validRule2  = "A0[0-9]"
    `

	_, err := LoadRule(toml)

	if err == nil {
		t.Errorf("Failed to raise error.")
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
