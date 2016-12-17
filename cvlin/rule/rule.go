package rule

import (
	"regexp"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

func LoadRule(tomlStr string) ([]string, error) {
	var rulemap map[string]string

	meta, err := toml.Decode(tomlStr, &rulemap)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse rule file")
	}

	orderedRules := convertToOrderedSlice(rulemap, meta.Keys())

	for _, rule := range orderedRules {
		if _, err := regexp.Compile(rule); err != nil {
			return nil, errors.Wrap(err, "The rule file conatins invalid regexp")
		}
	}

	return orderedRules, nil
}

// Convert a map to a slice which is ordered as same as the original TOML document
func convertToOrderedSlice(rulemap map[string]string, keys []toml.Key) []string {
	ordered := make([]string, 0)
	for _, key := range keys {
		ordered = append(ordered, rulemap[key[0]])
	}
	return ordered
}
