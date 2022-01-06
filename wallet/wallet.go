package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/gguldduck111/nomadcoin/Util"
)

const (
	signature 		string = "d3b34fc64d6834cf9b5325159e4748cd2a4a80d13b7b4130db056b200b62a83dc9f610c77e2f2c09b251e19c2c79f0b20489750c3e7cc32a008e3b27d2b3c37c"
	privateKey 		string = "3077020101042046b39a7d46192c9f47478f762b0709ec52a1fe12dbfe4ca1041a4eabf10fb499a00a06082a8648ce3d030107a144034200049d013dbbf9017aae2a19c115d49b304e8554ebb09e71198f12af92d756501e472e278223c70fbedbd96529e880ea6a1060cf9d41c566f5a1bca756e83aacd263"
	hashedMessage 	string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
)

func Start(){
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	Util.HandleErr(err)

	keyAsByte,err := x509.MarshalECPrivateKey(privateKey)
	Util.HandleErr(err)

	fmt.Printf("%x\n",keyAsByte)
	
	hashAsByte, err := hex.DecodeString(hashedMessage)
	Util.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader,privateKey,hashAsByte)
	Util.HandleErr(err)
	signature := append(r.Bytes(),s.Bytes()...)
	fmt.Printf("%x\n",signature)
	ok := ecdsa.Verify(&privateKey.PublicKey,hashAsByte,r,s)
	fmt.Println(ok)
}