package cvlin

import (
	"regexp"
	"fmt"
	"io/ioutil"
	"github.com/pkg/errors"
	"github.com/megane42/cvlin/cvlin/rule"
	"github.com/megane42/cvlin/cvlin/subject"
)


func Run(rulePath, subjectPath string) (bool, error) {

	ruleStr, err := readRuleFile(rulePath)
	if err != nil {
		return false, err
	}

	rules, err := rule.LoadRule(ruleStr)
	if err != nil {
		return false, err
	}

	subj, err := subject.LoadSubject(subjectPath)
	if err != nil {
		return false, err
	}

	return validate(rules, subj)
}


func validate(rules []*regexp.Regexp, subject [][]string) (bool, error) {
	for j, row := range subject {

		// validate if #rules is match with #columns
		if len(rules) != len(row) {
			return false, fmt.Errorf("The number of rules (%d) doesn't correspond to the number of columns (%d).", len(rules), len(row))
		}

		// validate if a value satisfies a rule
		for i, col := range row {
			if !(*rules[i]).MatchString(col) {
				return false, fmt.Errorf("Invalid. ( line: %d, column: %d, value: %s, rules: %s )", j, i, col, rules[i])
			}
		}
	}

	return true, nil
}


func readRuleFile(rulePath string) (string, error) {
	var tomlBytes []byte
	var err error

	if rulePath == "" {
		tomlBytes, err = Asset("default_rules.toml")
		if err != nil {
			return "", errors.Wrap(err, "Failed to load default rule")
		}

	} else {
		tomlBytes, err = ioutil.ReadFile(rulePath)
		if err != nil {
			return "", errors.Wrap(err, "Failed to load rule file")
		}
	}

	return string(tomlBytes), nil
}
