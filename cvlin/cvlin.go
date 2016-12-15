package cvlin

import (
	"regexp"
	"fmt"
	"github.com/megane42/cvlin/cvlin/rule"
	"github.com/megane42/cvlin/cvlin/subject"
)


func Run(rulePath, subjectPath string) (bool, error) {
	rule, err := rule.LoadRule(rulePath)
	if err != nil {
		return false, err
	}

	subj, err := subject.LoadSubject(subjectPath)
	if err != nil {
		return false, err
	}

	return validate(rule, subj)
}


func validate(rules []string, subject [][]string) (bool, error) {
	for j, row := range subject {

		// validate if #rules is match with #columns
		if len(rules) != len(row) {
			return false, fmt.Errorf("The number of rules (%d) doesn't correspond to the number of columns (%d).", len(rules), len(row))
		}

		// validate if a value satisfies a rule
		for i, col := range row {
			reg := regexp.MustCompile(rules[i])
			if !reg.MatchString(col) {
				return false, fmt.Errorf("Invalid. ( line: %d, column: %d, value: %s, rules: %s )", j, i, col, rules[i])
			}
		}
	}

	return true, nil
}
