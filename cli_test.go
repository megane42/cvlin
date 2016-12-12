package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./cvlin -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("Unexpected Status Code: expected %d but got %d", ExitCodeOK, status)
	}

	expected := fmt.Sprintf("cvlin version %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("Unexpected output: expected %q but got %q", expected, errStream.String(),)
	}
}

func TestRun_ruleFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./cvlin --rule example/rule.toml example/subject.csv", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("Unexpected Status Code: expected %d but got %d", ExitCodeOK, status)
	}
}
