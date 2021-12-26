package blockchain

import (
	"fmt"
	"sync"

	"github.com/gguldduck111/nomadcoin/Util"
	"github.com/gguldduck111/nomadcoin/db"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte)  {
	Util.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(Util.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := CreateBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			// Search for checkpoint on the DB
			checkpoint := db.Checkpoint();
			if checkpoint == nil {
				b.AddBlock("Genesis")
			}else{
				// Restore b from bytes
				b.restore(checkpoint) 
			}
		})
	}

	fmt.Println(b.NewestHash)
	return b
}
