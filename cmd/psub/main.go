package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

func exitOnError(err error, code int) {
	fmt.Fprintf(os.Stderr, "%s", err.Error())
	os.Exit(1)
}

func main() {
	dir, err := ioutil.TempDir("", "psub-*")
	if err != nil {
		os.RemoveAll(dir)
		exitOnError(err, 1)
	}

	fifoPath := filepath.Join(dir, "psub.fifo")

	err = syscall.Mkfifo(fifoPath, 0666)
	if err != nil {
		os.RemoveAll(dir)
		exitOnError(err, 2)
	}

	var attr = os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			nil,
			nil,
		},
	}
	process, err := os.StartProcess("/bin/sh",
		[]string{
			"/bin/sh",
			"-c",
			fmt.Sprintf("cat > %s ; rm -rf %s", fifoPath, dir),
		},
		&attr)

	if err != nil {
		os.RemoveAll(dir)
		exitOnError(err, 3)
	}

	err = process.Release()
	if err != nil {
		os.RemoveAll(dir)
		exitOnError(err, 4)
	}
	absPath, err := filepath.Abs(fifoPath)
	if err != nil {
		os.RemoveAll(dir)
		exitOnError(err, 5)
	}

	fmt.Println(absPath)
}
