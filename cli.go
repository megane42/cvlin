package main

import (
	"flag"
	"fmt"
	"io"
	"github.com/megane42/cvlin/cvlin"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK      int = 0
	ExitCodeError   int = 1 + iota
	ExitCodeInvalid
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		rulePath string
		version  bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&rulePath, "rule", "", "Path to rule file")
	flags.StringVar(&rulePath, "r",    "", "Path to rule file (Short)")
	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	// Validate CSV
	subjectPath := flags.Arg(0)
	_, err := cvlin.Run(rulePath, subjectPath)
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err)
		return ExitCodeError
	}

	fmt.Fprintf(cli.outStream, "Valid.\n")
	return ExitCodeOK
}
