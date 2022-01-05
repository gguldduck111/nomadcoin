package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/gguldduck111/nomadcoin/Util"
	"github.com/gguldduck111/nomadcoin/db"
)

type Block struct {
	Hash     	string 	`json:"hash"`
	PrevHash 	string 	`json:"prevHash,omitempty"`
	Height   	int    	`json:"height"`
	Difficulty 	int 	`json:"difficulty"`
	Nonce		int 	`json:"nonce"`
	Timestamp 	int		`json:"timestamp"`
	Transactions []*Tx 	`json:"transactions"`
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
		b.Timestamp = int(time.Now().Unix())
		hash := Util.Hash(b)
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

func CreateBlock(prevHash string, height int, diff int) *Block {
	block := &Block{
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
		Difficulty: diff,
		Nonce: 0,
	}

	block.mine()
	block.Transactions = Mempool.TxToConfirm()
	block.persist()
	return block
}
