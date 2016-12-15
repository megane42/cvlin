package rule

import (
	"github.com/BurntSushi/toml"
)

func LoadRule(path string) ([]string, error) {
	var rulemap map[string]string

	meta, err := toml.DecodeFile(path, &rulemap)
	if err != nil {
		return nil, err
	}

	return convertToOrderedSlice(rulemap, meta.Keys()), nil
}

// Convert a map to a slice which is ordered as same as the original TOML document
func convertToOrderedSlice(rulemap map[string]string, keys []toml.Key) []string {
	ordered := make([]string, 0)
	for _, key := range keys {
		ordered = append(ordered, rulemap[key[0]])
	}
	return ordered
}
