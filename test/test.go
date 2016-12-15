package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func main() {
	type Rule struct {
		Pattern string
	}

	type Config map[string]Rule

	var conf Config

	if _, err := toml.DecodeFile("./test.toml", &conf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conf)
	return
}
