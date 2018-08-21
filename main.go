package main

import (
	"fmt"
	"time"

	"github.com/d5c5ceb0/newchain/chain"
	"github.com/d5c5ceb0/newchain/storage"
)

func main() {
	db, err := storage.NewDb(storage.LevelDB, "./db/db")
	if err != nil {
		return
	}
	blockchain := chain.NewBlockChain(db)

	for i := 0; i < 100000; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(blockchain.CurrentBlockHash())
	}

	return
}
