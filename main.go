package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gguldduck111/nomadcoin/Util"
	"github.com/gguldduck111/nomadcoin/blockchain"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL `json:"url"`
	Method      string
	Description string
	Payload     string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string
}

func documentation(writer http.ResponseWriter, request *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         URL("/blocks/{id}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		return
	}
}

func blocks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(blockchain.GetBlockchain().AllBlocks())
		Util.HandleErr(err)
	case "POST":
		var addBlockBody AddBlockBody
		Util.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
