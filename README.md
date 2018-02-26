# edyscoin Implementation
This is an implementation of the edyscoin cryptocurrency, created for educational purposes.

## Installation
1. Make sure you have the most recent version of Golang installed. I used version 1.9.2 at the time writing this README.
	* See (https://golang.org/doc/install) for specific details for your operating system.
	* I am on MacOS so I installed via Homebrew.
2. Within this directory, run `go install edyscoin`.
3. Run using `./bin/edyscoin [local host:port] [remote host:port]`
	* For the first node, the local and remote address should be the same, i.e. `./bin/edyscoin localhost:1111 localhost:1111`.

## Commands
* `list [peers|transaction|blockchain]`
	- Display either the currently connect peer list, current transaction list, or display the blockchain.
* `handshake [remote host:port]`
	- pings the remote node and adds it to the peer list. This command is rarely useful except for testing.
* `transaction [sender] [recipient] [amount]`
	- Adds a new transaction to the current blockchain and broadcasts the transaction to all other nodes as well.
* `mine`
	- Mine the current block and broadcast it to the rest of the network. The other nodes with verify the blockchain via the Consensus algorithm.

## Concepts Explored and to Explore in the Future
* blockchain - DONE
* transaction - DONE
* mining (new coins, difficulty target) - DONE
* consensus (proof of work, proof of stake) - DONE
* merkle tree
* transaction fees
* ownership (private keys)
* supply (coinbase)
* wallets (persistent storage)
* scalability (segregated witness/lightning network)
* block size bottleneck
* forks
* double spending
* 51% attack
 
## Is cryptocurrency a currency?
Money is:
* a store of value
* a medium of exchange
* unit of account

Further research and observation is needed.
