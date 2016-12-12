package cvlin

import (
	"testing"
)

func TestValidate(t *testing.T) {
	rule := map[string]rule {
		"id"    : {Pattern: "A0[0-9]",  Notnull: false},
		"name"  : {Pattern: "*",        Notnull: true},
		"point" : {Pattern: "^[0-9]+$", Notnull: true},
	}

	subject := [][]string{{"A01", "Shibuya Rin", "100"}, {"A02", "Sakuma Mayu", "200"}}

	res, err := validate(rule, subject)

	if !(err == nil) {
		t.Errorf("Failed to validate (unexpected error occurs). %s", err)
	}

	if !(res == true) {
		t.Errorf("Failed to validate (false negative). %s", res)
	}
}
