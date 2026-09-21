package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const version = "0.1.0"

func Greet(name string) string { return "hello, " + name }

// run returns the process exit code so main stays a one-liner and the flag
// handling is reachable from tests, which cannot observe os.Exit.
func run(args []string, out io.Writer) int {
	fs := flag.NewFlagSet("grappler-testbed", flag.ContinueOnError)
	fs.SetOutput(out)
	showVersion := fs.Bool("version", false, "print the version and exit")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *showVersion {
		fmt.Fprintln(out, version)
		return 0
	}
	fmt.Fprintln(out, Greet("world"))
	return 0
}

func main() { os.Exit(run(os.Args[1:], os.Stdout)) }
