package main

import (
	"encoding/json"
	"fmt"
	"github.com/gguldduck111/nomadcoin/Util"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL string
	Method string
	Description string
}

func documentation(writer http.ResponseWriter, request *http.Request) {
	data := URLDescription{
		URL: "/",
		Method: "GET",
		Description: "See Documentation",
	}
	b, err := json.MarshalIndent(data,""," ")
	Util.HandleErr(err)

	fmt.Printf("%s",b)
}

func main()  {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s",port)
	log.Fatal(http.ListenAndServe(port, nil))
}