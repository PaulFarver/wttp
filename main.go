package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Could not parse command line arguments: ", os.Args[1:])
		os.Exit(1)
	}
	arg := os.Args[1]
	code, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Code is not an integer: ", err)
		os.Exit(1)
	}
	fmt.Println(http.StatusText(code))
}
