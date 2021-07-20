package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) == 1 {
		r := io.TeeReader(os.Stdin, os.Stdout)
		io.ReadAll(r)
		return
	}

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			printStr("ERROR: open " + filename + ": no such file or directory\n")
			os.Exit(1)
			return
		}
		printStr(string(data))
	}
}
