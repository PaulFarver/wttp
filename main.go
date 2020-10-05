package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/goterm/term"
)

func main() {
	descriptionFlag := flag.Bool("long", false, "Print a long description of the HTTP status code")
	noColor := flag.Bool("nocolor", false, "Disable color in output")
	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	code, ok := codes[args[0]]
	if !ok {
		fmt.Println("Unknown HTTP code")
		os.Exit(1)
	}

	msg := code.Message
	if !*noColor {
		msg = addColor(args[0], code.Message)
	}

	fmt.Println(msg)
	if *descriptionFlag {
		fmt.Printf("\n%s\n", code.Description)
	}
}

func addColor(code, message string) string {
	switch string(code[0]) {
	case "1":
		return term.Bluef(message)
	case "2":
		return term.Greenf(message)
	case "3":
		return term.Cyanf(message)
	case "4":
		return term.Yellowf(message)
	case "5":
		return term.Redf(message)
	default:
		return message
	}
}
