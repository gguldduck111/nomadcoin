package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"math/big"

	"github.com/gguldduck111/nomadcoin/Util"
)

const (
	signature 		string = "d3b34fc64d6834cf9b5325159e4748cd2a4a80d13b7b4130db056b200b62a83dc9f610c77e2f2c09b251e19c2c79f0b20489750c3e7cc32a008e3b27d2b3c37c"
	privateKey 		string = "3077020101042046b39a7d46192c9f47478f762b0709ec52a1fe12dbfe4ca1041a4eabf10fb499a00a06082a8648ce3d030107a144034200049d013dbbf9017aae2a19c115d49b304e8554ebb09e71198f12af92d756501e472e278223c70fbedbd96529e880ea6a1060cf9d41c566f5a1bca756e83aacd263"
	hashedMessage 	string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
)

func Start(){
	privBytes, err := hex.DecodeString(privateKey)
	Util.HandleErr(err)
	_, err = x509.ParseECPrivateKey(privBytes) 
	Util.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
}