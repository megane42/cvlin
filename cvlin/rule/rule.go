package rule

import (
	"regexp"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

func LoadRule(tomlStr string) ([]*regexp.Regexp, error) {
	var rulemap map[string]string

	meta, err := toml.Decode(tomlStr, &rulemap)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse rule file")
	}

	orderedRules := convertToOrderedSlice(rulemap, meta.Keys())
	validRules, err := convertToSliceOfRegexp(orderedRules)
	if err != nil {
		return nil, errors.Wrap(err, "The rule file conatins invalid regexp")
	}

	return validRules, nil
}

// Convert a map to a slice which is ordered as same as the original TOML document
func convertToOrderedSlice(rulemap map[string]string, keys []toml.Key) []string {
	ordered := make([]string, 0)
	for _, key := range keys {
		ordered = append(ordered, rulemap[key[0]])
	}
	return ordered
}

// Convert slice of strings to slice of *regexp.Regexp
func convertToSliceOfRegexp(rules []string) ([]*regexp.Regexp, error) {
	validRules := make([]*regexp.Regexp, 0)

	for _, rule := range rules {
		reg, err := regexp.Compile(rule)
		if err != nil {
			return nil, err
		}
		validRules = append(validRules, reg)
	}

	return validRules, nil
}
