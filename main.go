package main

import (
	"flag"
	"fmt"
	"interpreter_made_in_go/repl"
	"io"
	"io/ioutil"
	"os"
	"os/user"
)

func init() {
	flag.Parse()
}

func main() {
	var filename string

	if args := flag.Args(); len(args) > 0 {
		filename = args[0]
	}

	var r io.Reader
	switch filename {
	case "":
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	default:
		f, err := os.Open(filename)
		if err != nil {
			fmt.Print(err)
		}
		defer f.Close()
		r = f

		b, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Print(err)
		}

		repl.StartFromTxtfile(string(b), os.Stdout)

	}

}
