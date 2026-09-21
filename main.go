package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func Greet(name string) string { return "hello, " + name }

func run(out io.Writer, asJSON bool) error {
	greeting := Greet("world")
	if !asJSON {
		_, err := fmt.Fprintln(out, greeting)
		return err
	}
	return json.NewEncoder(out).Encode(struct {
		Greeting string `json:"greeting"`
	}{greeting})
}

func main() {
	asJSON := flag.Bool("json", false, "print the greeting as a JSON object")
	flag.Parse()
	if err := run(os.Stdout, *asJSON); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
