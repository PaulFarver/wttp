package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	code, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Code is not an integer: ", err)
	}
	fmt.Println(http.StatusText(code))
}
