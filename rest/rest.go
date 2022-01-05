package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gguldduck111/nomadcoin/Util"
	"github.com/gguldduck111/nomadcoin/blockchain"
	"github.com/gorilla/mux"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url `json:"url"`
	Method      string
	Description string
	Payload     string `json:"payload,omitempty"`
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type balanceResponse struct{
	Address string 	`json:"address"`
	Amount int 		`json:"amount"`
}

type addTxPayload struct {
	To string `json:"to"`
	Amount int `json:"amount"`
}

func documentation(writer http.ResponseWriter, request *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "See A Block",
		},
		{
			URL: 		url("/balance/{address}"),
			Method:      "GET",
			Description: "Get TxOuts for an Address",
		},
	}
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		return
	}
}

func blocks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(blockchain.Blockchain().Blocks()) 
	case "POST":
		blockchain.Blockchain().AddBlock()
		w.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash:= vars["hash"]
	
	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func status(rw http.ResponseWriter, r *http.Request){
	json.NewEncoder(rw).Encode(blockchain.Blockchain())
}

func balance(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount:=blockchain.Blockchain().BalanceByAddress(address)
		json.NewEncoder(rw).Encode(balanceResponse{address,amount})
	default:
		Util.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain().UTxOutsByAddress(address)))
	}
}

func mempool(rw http.ResponseWriter,r *http.Request){
	Util.HandleErr(json.NewEncoder(rw).Encode(blockchain.Mempool.Txs))
}

func transactions(rw http.ResponseWriter,r *http.Request){
	var payload addTxPayload
	Util.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchain.Mempool.AddTx(payload.To,payload.Amount)
	if err != nil{
		json.NewEncoder(rw).Encode(errorResponse{"not enough fund"})
	}
	rw.WriteHeader(http.StatusCreated)
}

func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f 0-9]+}", block).Methods("GET")
	router.HandleFunc("/balance/{address}", balance)
	router.HandleFunc("/mempool", mempool)
	router.HandleFunc("/transactions", transactions).Methods("POST")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
