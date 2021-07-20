package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Print("Too many arguments\n")
		return
	}

	if len(os.Args) == 1 {
		fmt.Print("File name missing\n")
		return
	}

	filename := os.Args[1]

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(string(dat))
}
