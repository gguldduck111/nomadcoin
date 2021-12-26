package main

import (
	"github.com/gguldduck111/nomadcoin/cli"
	"github.com/gguldduck111/nomadcoin/db"
)

func main() {
	defer db.DB().Close()

	cli.Start()
}
