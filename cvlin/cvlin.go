package cvlin

import (
	"regexp"
	"fmt"
	"github.com/megane42/cvlin/cvlin/rule"
	"github.com/megane42/cvlin/cvlin/subject"
)

type InvalidNumOfRulesError struct {
	rules []string
	row   []string
}

func (e InvalidNumOfRulesError) Error() string {
	return fmt.Sprintf("Error: The number of rules (%d) doesn't match with the number of columns (%d).", len(e.rules), len(e.row))
}

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
	result := true
	for _, row := range subject {

		// validate if #rules is match with #columns
		if len(rules) != len(row) {
			return false, InvalidNumOfRulesError{rules, row}
		}

		// validate if a value satisfies a rule
		for i, col := range row {
			reg := regexp.MustCompile(rules[i])
			result = result && reg.MatchString(col)
		}
	}
	return result, nil
}
