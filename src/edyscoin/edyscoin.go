package main

import (
	"bufio"
	"libedyscoin"
	"fmt"
	"log"
	"flag"
	"os"
	"strconv"
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
	fmt.Printf("local node: %+v\n", lnode)
	res, err := lnode.DoHandshake(args[1])
	if err != nil {
		log.Fatal("ERR-> remote node offline or incorrect address")
	}
	fmt.Print("OK-> resp: "+ res.SenderId.ToString() + " from addr: " + res.SenderAddr + "\n")

	lnode.DoSyncBlockChain(args[1])

	startCLI(lnode)
}

func startCLI(lnode *libedyscoin.Node) {
	input := bufio.NewReader(os.Stdin)
	quit := false
	for !quit {
		fmt.Print("edyscoin> ")
		line, err := input.ReadString('\n')
		if err != nil {
			log.Fatal("input error: ", err)
		}
		line = strings.TrimSpace(line)
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
	toks := strings.Split(line, " ")
	command := toks[0]

	switch command {
	case "quit":
		fallthrough
	case "exit":
		return "quit"

	case "list":
		if len(toks) != 2 {
			return "ERR-> usage: `list [peers|transaction|blockchain]`"
		}
		if toks[1] == "peers" {
			str := ""
			for k, v := range lnode.Peers {
				str += fmt.Sprintf("%+v %+v\n", k, v)
			}
			return str
		} else if toks[1] == "transaction" || toks[1] == "txn" {
			lnode.BlockChain.ListTransactions()
		} else if toks[1] == "blockchain" || toks[1] == "bc" {
			lnode.BlockChain.DisplayBlockChain()
		} else {
			return "ERR-> usage: `list [peers|transaction|blockchain]`"
		}
		return ""

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
		msg := libedyscoin.NewMessage(lnode, "Broadcast")
		rnodes, _ := lnode.DoBroadcast(msg)
		return "OK-> broadcast to all nodes:\n" + fmt.Sprintf("%+v", rnodes)

	case "txn":
		fallthrough
	case "transaction":
		if len(toks) != 4 {
			return "ERR-> usage: `transaction [sender] [recipient] [amount]`"
		}
		amount, err := strconv.ParseInt(toks[3], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		txn := &libedyscoin.Transaction{
			Sender:    toks[1],
			Recipient: toks[2],
			Amount:    amount,
		}
		lnode.BlockChain.NewTransaction(txn)
		rnodes, _ := lnode.DoBroadcastNewTransaction(txn)
		return "OK-> broadcasted a transaction to all nodes:\n" + fmt.Sprintf("%+v", rnodes)

	case "mine":
		if len(toks) != 1 {
			return "ERR-> usage: `mine`"
		}
		mined := lnode.BlockChain.Mine()
		if mined {
			fmt.Printf("OK-> mined a new block!\n")
			lnode.BlockChain.DisplayBlockChain()
			rnodes, _ := lnode.DoBroadcastNewBlockChain(lnode.BlockChain)
			return "OK-> broadcasted new block to all nodes:\n" + fmt.Sprintf("%+v", rnodes)
		}
	}

	return "ERR-> command not recognized"
}
