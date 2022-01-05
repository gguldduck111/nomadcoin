package blockchain

import (
	"sync"

	"github.com/gguldduck111/nomadcoin/Util"
	"github.com/gguldduck111/nomadcoin/db"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
	CurrentDifficulty int `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

const (
	defaultDifficulty int = 2
	difficultyInterval int = 5
	blockInterval int = 2
	allowedRange int = 2
)

func (b *blockchain) restore(data []byte)  {
	Util.FromBytes(b, data)
}

func (b *blockchain) AddBlock() {
	block := CreateBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)
}

func persistBlockchain(b *blockchain) {
	db.SaveCheckpoint(Util.ToBytes(b))
}


func recalculateDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
		newestBlock := allBlocks[0]
		lastRecalculateBlock := allBlocks[difficultyInterval - 1]
		actualTime := (newestBlock.Timestamp/60) - (lastRecalculateBlock.Timestamp/60)
		expectedTime := difficultyInterval * blockInterval
		if actualTime <= (expectedTime-allowedRange) {
			return b.CurrentDifficulty + 1
		} else if actualTime >= (expectedTime+allowedRange) {
			return b.CurrentDifficulty - 1		
		}
		return b.CurrentDifficulty
}

func getDifficulty(b * blockchain) int{
	if b.Height == 0 {
		return defaultDifficulty
	}else if b.Height % 5 == 0{
		// Recalculate the difficulty
		return recalculateDifficulty(b)
	}else{
		return b.CurrentDifficulty
	}
}

func Blocks(b *blockchain) []*Block{
	var blocks []*Block 
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		if block != nil {
			blocks = append(blocks,block)
			hashCursor = block.PrevHash
		} else{
			break
		}
	}
	return blocks
}

func UTxOutsByAddress(address string,b *blockchain) []*UTxOut{
	var uTxOuts []*UTxOut
	creatorIds := make(map[string]bool)

	 for _, block := range Blocks(b){
		 for _, tx := range block.Transactions{
			 for _,input := range tx.TxIns{
				 if input.Owner == address {
					 creatorIds[input.TxID] = true
				 }
			 }
			 for index, output := range tx.TxOuts{
				if output.Owner == address {
					if _, ok := creatorIds[tx.Id]; !ok {
						uTxOut := &UTxOut{tx.Id,index,output.Amount}
						if !isOnMempool(uTxOut) {
							uTxOuts = append(uTxOuts, uTxOut)
						}
					}
				}
			 }
		 }
	 }

	 return uTxOuts
}

func BalanceByAddress(address string,b *blockchain) int{
	var amount int
	txOuts := UTxOutsByAddress(address,b)
	for _,txOut := range txOuts{
		amount += txOut.Amount
	}
	return amount
} 

func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{Height: 0}
		// Search for checkpoint on the DB
		checkpoint := db.Checkpoint();
		if checkpoint == nil {
			b.AddBlock()
		}else{
			// Restore b from bytes
			b.restore(checkpoint) 
		}
	})
	return b
}
