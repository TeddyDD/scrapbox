package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

const cost = 10

func main() {
	user := flag.String("user", "", "username")
	password := flag.String("password", "", "password")
	flag.Parse()

	hash, err := bcrypt.GenerateFromPassword([]byte(*password), cost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't calculate hash '%s'\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s:%s%s%s\n",
		*user, string(hash[0:2]), "y", string(hash[3:]))
}
