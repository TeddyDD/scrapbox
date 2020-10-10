package main

import (
	"flag"
	"fmt"
	"os"

	gonanoid "github.com/matoous/go-nanoid"
)

func main() {
	c := flag.Int("c", 16, "count")
	flag.Parse()
	gen, err := gonanoid.ID(*c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	fmt.Println(gen)
}
