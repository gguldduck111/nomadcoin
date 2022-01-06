package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/gguldduck111/nomadcoin/Util"
)

func Start(){
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	Util.HandleErr(err)
	message := "I love you"
	hashedMessage := Util.Hash(message)

	hashAsByte, err := hex.DecodeString(hashedMessage)
	Util.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader,privateKey,hashAsByte)
	Util.HandleErr(err)

	fmt.Printf("R:%d S:%d",r,s)
}