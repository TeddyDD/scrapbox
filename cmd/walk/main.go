package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	dirs := flag.Bool("dirs", false, "show only directories")
	nodirs := flag.Bool("nodirs", false, "don't show directories")
	zero := flag.Bool("0", false, `use \0 to separate paths`)
	flag.Parse()

	separator := "\n"
	if *zero {
		separator = "\x00"
	}

	paths := []string{}
	if len(flag.Args()) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		paths = append(paths, cwd)
	} else {
		paths = flag.Args()
	}

	out := bufio.NewWriter(os.Stdout)
	defer func() {
		err := out.Flush()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(2)
		}
	}()

	for _, p := range paths {
		filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if *dirs && !info.IsDir() {
				return nil
			}
			if *nodirs && info.IsDir() {
				return nil
			}
			fmt.Fprintf(out, "%s%s", path, separator)
			return nil
		})
	}
}
