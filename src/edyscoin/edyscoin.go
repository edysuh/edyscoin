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
	res, err := lnode.DoHandshake(args[1])
	if err != nil {
		log.Fatal("ERR-> remote node offline or incorrect address")
	}
	fmt.Print("OK-> resp: "+ res.SenderId.ToString() + " from addr: " + res.SenderAddr + "\n")

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

	case "peers":
		str := ""
		for k, v := range lnode.Peers {
			str += fmt.Sprintf("%+v %+v\n", k, v)
		}
		return str

	case "handshake":
		if len(toks) != 2 {
			return "ERR-> usage: `handshake [remote node host:port]`"
		}
		res, err := lnode.DoHandshake(toks[1])
		if err != nil {
			return "ERR-> remote node offline or incorrect address"
		}
		return "OK-> response from node id: "+ res.SenderId.ToString() +
			" from addr: " + res.SenderAddr

	case "broadcast":
		if len(toks) != 2 {
			return "ERR-> usage: `broadcast [string]`"
		}
		msg := &libedyscoin.Message{
			MsgId: libedyscoin.NewId(lnode.Address),
			SenderId: lnode.Id,
			SenderAddr: lnode.Address,
			Method: "Broadcast",
			Args: toks[1],
			Seenlist: make(map[libedyscoin.Id]bool),
			Success: false,
		}
		lnode.DoBroadcast(msg)
		return "OK-> broadcast to all nodes:\n"
		// res, err := lnode.DoBroadcast(toks[1], q, make(map[libedyscoin.Id]bool))
		// if err != nil {
		// 	return "broadcast error"
		// }
		// return "OK-> broadcast to all nodes:\n" + res.ResNode.ToString()
	}
	return "ERR-> command not recognized"
}
