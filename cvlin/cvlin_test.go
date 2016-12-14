package cvlin

import (
	"testing"
)

func TestValidate(t *testing.T) {
	rule := []rule {
		{Pattern: "A0[0-9]",  Notnull: false},
		{Pattern: ".*",       Notnull: true},
		{Pattern: "^[0-9]+$", Notnull: true},
	}

	subject := [][]string {
		{"A01", "Shibuya Rin", "100"},
		{"A02", "Sakuma Mayu", "200"},
	}

	res, err := validate(rule, subject)

	if err != nil {
		t.Errorf("Failed to validate (unexpected error occurs). %s", err)
	}

	if res != true {
		t.Errorf("Failed to validate (false negative). %s", res)
	}
}

func TestValidate_InvalidNumOfRules(t *testing.T) {
	// 2 rules for 3 cols
	rule := []rule {
		{Pattern: "A0[0-9]",  Notnull: false},
		{Pattern: ".*",       Notnull: true},
	}

	subject := [][]string {
		{"A01", "Shibuya Rin", "100"},
		{"A02", "Sakuma Mayu", "200"},
	}

	res, err := validate(rule, subject)

	if err == nil {
		t.Errorf("Failed to validate (failed to raise error). %s", err)
	}

	if res == true {
		t.Errorf("Failed to validate (false positive). %s", res)
	}
}
