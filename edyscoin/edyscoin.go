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

// start the first node with `./edyscoin [local host:port] [remote host:port]`
func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Fatal("improper usage: ./edyscoin [local host:port] " +
			"[remote host:port]")
	}

	lnode := libedyscoin.NewNode(args[0])
	fmt.Printf("local node id and address: %v, %v\n", lnode.Id, lnode.Address)

	startCLI(lnode)
}

func startCLI(lnode *libedyscoin.Node) {
	input := bufio.NewReader(os.Stdin)
	quit := false
	for !quit {
		fmt.Print("edyscoin > ")
		line, err := input.ReadString('\n')
		if err != nil {
			log.Fatal("input error: ", err)
		}
		if line == "" {
			continue
		}

		resp := executeLine(lnode, line)
		if resp == "quit" {
			quit = true
		} else if resp != "" {
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
		if len(toks) != 2 {
			return "ERR -> usage: `handshake [remote node host:port]`"
		}
		res, err := lnode.DoHandshake(toks[1])
		if err != nil {
			return "ERR -> remote node offline or incorrect address"
		}
		return "OK -> response from node id:\n"+ res.NodeId.ToString() +
			"\nfrom addr: " + res.Address
	}
	return "ERR -> command not recognized"
}
