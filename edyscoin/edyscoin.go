package main

import (
	"libedyscoin"
	"fmt"
	"os"
)

// start the first node with `./edyscoin [remote node address:port]`
func main() {
	args := os.Args
	hh := libedyscoin.HttpHandler{}

	switch {
	case len(args) < 2:
		fmt.Println("improper usage: ./edyscoin [remote node address:port]")
		os.Exit(1)
	case len(args) == 2:
		hh.NewHandler("8080")
	case len(args) == 3:
		libedyscoin.Connect(args[1])
	}
}
