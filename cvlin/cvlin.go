package cvlin

import (
	"regexp"
	"fmt"
)

type InvalidNumOfRulesError struct {
	rules []rule
	row   []string
}

func (e InvalidNumOfRulesError) Error() string {
	return fmt.Sprintf("Error: The number of rules (%d) doesn't match with the number of columns (%d).", len(e.rules), len(e.row))
}


func Run(rulePath, subjectPath string) (bool, error) {
	rule, err := LoadRule(rulePath)
	if err != nil {
		return false, err
	}

	subj, err := LoadSubject(subjectPath)
	if err != nil {
		return false, err
	}

	return validate(rule, subj)
}

func validate(rules []rule, subject [][]string) (bool, error) {
	result := true
	for _, row := range subject {
		if len(rules) != len(row) {
			return false, InvalidNumOfRulesError{rules, row}
		}

		for i, col := range row {
			reg := regexp.MustCompile(rules[i].Pattern)
			result = result && reg.MatchString(col)
		}
	}
	return result, nil
}
