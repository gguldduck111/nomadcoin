package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/gguldduck111/nomadcoin/Util"
	"github.com/gguldduck111/nomadcoin/db"
)

const difficulty int =2 

type Block struct {
	Data     	string 	`json:"data"`
	Hash     	string 	`json:"hash"`
	PrevHash 	string 	`json:"prevHash,omitempty"`
	Height   	int    	`json:"height"`
	Difficulty 	int 	`json:"difficulty"`
	Nonce		int 	`json:"nonce"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, Util.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte){
	Util.FromBytes(b, data)
}

func (b *Block) mine(){
	target := strings.Repeat("0",b.Difficulty)
	for{
		blockAsString := fmt.Sprint(b)
		hash := fmt.Sprintf("%x",sha256.Sum256([]byte(blockAsString)))
		fmt.Sprintf("Block as string:%s\nHash:%s\nNonce:%d\n\n\n",blockAsString,hash,b.Nonce)
		if strings.HasPrefix(hash,target) {
			b.Hash = hash
			break
		}else{
			b.Nonce++
		}
	}
}

func FindBlock(hash string) (*Block,error){
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil,ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil 
}

func CreateBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
		Difficulty: difficulty,
		Nonce: 0,
	}

	block.mine()
	block.persist()
	return block
}
