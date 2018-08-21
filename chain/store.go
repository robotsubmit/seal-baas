package chain

import "github.com/d5c5ceb0/newchain/storage"

type DbPrefix byte

const (
	BlockHashByHeight DbPrefix = 0x00
	BlockByHash       DbPrefix = 0x01
	CurrentBlock      DbPrefix = 0x02
	AccoutByAddress   DbPrefix = 0x03
)

type Store struct {
	db *storage.Database
}

func NewStore(db *storage.Database) *Store {
	return &Store{
		db: db,
	}
}

func (this *Store) StoreBlock() {
}
func (this *Store) GetBlockByHash() {
}
func (this *Store) GetBlockByHeight() {
}
func (this *Store) GetCurrentBlock() {
}
func (this *Store) GetAccount() {
}
