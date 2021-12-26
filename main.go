package main

import (
	"github.com/gguldduck111/nomadcoin/blockchain"
	"github.com/gguldduck111/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
