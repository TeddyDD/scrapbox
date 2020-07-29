package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yuin/goldmark"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	err = goldmark.Convert(b, os.Stdout)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
}
