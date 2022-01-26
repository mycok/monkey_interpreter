package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mycok/monkey_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! Welcome to the monkey programing language!\n", user.Username)
	fmt.Printf("Feel free to type commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
