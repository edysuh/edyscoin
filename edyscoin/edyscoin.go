package main

import (
	"bufio"
	"libedyscoin"
	"fmt"
	"log"
	"flag"
	"os"
	"strings"
)

// start the first node with `./edyscoin [local addr:port] [remote addr:port]`
func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Fatal("improper usage: ./edyscoin [local address:port] " +
			"[remote addr:port]")
	}

	lnode := libedyscoin.NewNode(args[0])
	fmt.Printf("from main: %v, %v\n", lnode.Id, lnode.Address)
	startCLI(lnode)
}

func startCLI(lnode *libedyscoin.Node) {
	input := bufio.NewReader(os.Stdin)
	quit := false
	for !quit {
		fmt.Print("edyscoin > ")
		line, err := input.ReadString('\n')
		if err != nil {
			log.Fatal("input error", err)
		}
		if line == "" {
			continue
		}

		resp := executeLine(lnode, line)
		if resp == "quit" {
			quit = true
		} else {
			fmt.Printf("%s\n", resp)
		}
	}
}

func executeLine(lnode *libedyscoin.Node, line string) string {
	line = strings.TrimSpace(line)
	toks := strings.Split(line, " ")
	command := toks[0]

	switch command {
	case "quit":
		fallthrough
	case "exit":
		return "quit"
	case "handshake":
		res, err := lnode.DoHandshake(toks[1])
		if err != nil {
			log.Fatal("error in handshake", err)
		}
		return "OK: "+ res.NodeId.String() + " " + res.Address
	}
	return "ERR: command not recognized"
}
