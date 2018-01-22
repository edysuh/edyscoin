package main

import (
	// "libedyscoin"
	"log"
	"flag"
)

// start the first node with `./edyscoin [remote addr:port] [local addr:port]`
func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Fatal("improper usage: ./edyscoin [remote address:port] " +
			"[local addr:port]")
	}

	// fmt.Printf("%s %s\n", args[0], args[1])
}
