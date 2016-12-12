package cvlin

import (
	"github.com/BurntSushi/toml"
)

type rule struct {
	Pattern string
	Notnull bool
}

func LoadRule(path string) (map[string]rule, error) {
	var r map[string]rule

	// load toml
	if _, err := toml.DecodeFile(path, &r); err != nil {
		return nil, err
	}

	// set default pattern
	for i := range r {
		if r[i].Pattern == "" {
			// What I really want to do here is just `r[i].Pattern = "*"`, but it's not allowed.
			// http://stackoverflow.com/questions/32751537/
			// http://stackoverflow.com/questions/15984423/
			r[i] = rule{Pattern: "*", Notnull: r[i].Notnull}
		}
	}

	return r, nil
}
