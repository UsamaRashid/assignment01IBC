package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateHash(inputBlock *Block) string {
	// return "123"
	var toReturn string
	for i := range inputBlock.Data.Transactions {
		toReturn = toReturn + (inputBlock.Data.Transactions[i])
	}
	toReturn = fmt.Sprintf("%x", sha256.Sum256([]byte(toReturn)))
	return toReturn
}
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {

	if chainHead == nil {
		// If first Block
		//making a new Block to put in chain
		fmt.Print("FIRST BLOCK ADDED \n")

		var newBlock Block
		newBlock.Data = dataToInsert
		newBlock.PrevPointer = chainHead
		newBlock.PrevHash = ""
		chainHead = &newBlock
		newBlock.CurrentHash = CalculateHash(chainHead)
		// {dataToInsert, chainHead, "", CalculateHash(chainHead)}

		return &newBlock
	} else {
		fmt.Print("123 :::::::\n")
		// newBlock := Block{dataToInsert, chainHead, CalculateHash(chainHead), "chainHead.PrevPointer.CurrentHash"}
		var newBlock Block
		newBlock.Data = dataToInsert
		newBlock.PrevPointer = chainHead
		chainHead = &newBlock
		newBlock.PrevHash = chainHead.PrevPointer.CurrentHash
		newBlock.CurrentHash = CalculateHash(chainHead)
		fmt.Print("123 ::::::: 456")

		return &newBlock

	}
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	for i := 0; chainHead != nil; i++ {
		fmt.Print("CHANGE BLOCK:, ", chainHead.Data.Transactions, "\n")
		for i := range chainHead.Data.Transactions {
			if chainHead.Data.Transactions[i] == oldTrans {
				chainHead.Data.Transactions[i] = newTrans
				return
			}
		}
		chainHead = chainHead.PrevPointer
	}
}
func ListBlocks(chainHead *Block) {
	fmt.Print("\n############################################################################\n")
	fmt.Print("List Transactions")
	fmt.Print("\n############################################################################\n")
	for index := chainHead; index != nil; index = index.PrevPointer {
		fmt.Print("\n Transactions: ", index.Data.Transactions)
		fmt.Print("\n Prev Hash: ", index.PrevHash)
		fmt.Print("\n Curr Hash: ", index.CurrentHash)
		fmt.Print("\n")
	}
	fmt.Print("\n############################################################################\n")
}
func VerifyChain(chainHead *Block) {
	// ...
	for i := 0; chainHead.PrevPointer != nil; i++ {
		var verifyHash string = CalculateHash(chainHead.PrevPointer)
		if verifyHash != chainHead.PrevHash {
			fmt.Print("BlockChain is comprised!")
			return
		}
		chainHead = chainHead.PrevPointer
	}
	println("BlockChain is verified!")
}
