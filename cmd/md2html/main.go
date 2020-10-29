package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	err = md.Convert(b, os.Stdout)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
}
