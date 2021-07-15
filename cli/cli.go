package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/gguldduck111/nomadcoin/explorer"
	"github.com/gguldduck111/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following commands\n\n")
	fmt.Printf("-port: 	Set the PORT of the server\n")
	fmt.Printf("-mode: 	Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "test", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "all":
		go rest.Start(4000)
		explorer.Start(5000)
	case "rest":
		//start rest api
		rest.Start(*port)
	case "html":
		//start html explorer
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
