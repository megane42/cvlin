package cvlin

import (
	"github.com/BurntSushi/toml"
)

type rule struct {
	Pattern string
	Notnull bool
}

func LoadRule(path string) ([]rule, error) {
	var rulemap map[string]rule

	meta, err := toml.DecodeFile(path, &rulemap)
	if err != nil {
		return nil, err
	}

	// Set default pattern
	for i := range rulemap {
		if rulemap[i].Pattern == "" {
			// What I really want to do here is just `rulemap[i].Pattern = "*"`, but it's not allowed.
			// http://stackoverflow.com/questions/32751537/
			// http://stackoverflow.com/questions/15984423/
			rulemap[i] = rule{Pattern: "*", Notnull: rulemap[i].Notnull}
		}
	}

	titles  := extractTitles(meta.Keys())
	ordered := sort(rulemap, titles)
	return ordered, nil
}


// Extract the titles of key groups.
// Note that the titles have the same order of appearance in the original TOML document
func extractTitles(keys []toml.Key) []string {
	titles := make([]string, 0)
	for _, key := range keys {
		// If a key has only 1 element, it indicates the title of a key group
		if len(key) == 1 {
			titles = append(titles, key[0])
		}
	}
	return titles
}

// Sort the rules by same order of given keys.
func sort(rmap map[string]rule, keys []string) []rule {
	ordered := make([]rule, 0)
	for _, key := range keys {
		ordered = append(ordered, rmap[key])
	}
	return ordered
}
