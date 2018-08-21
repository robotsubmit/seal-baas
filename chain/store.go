package chain

import (
	"encoding/binary"
	"errors"

	"github.com/d5c5ceb0/newchain/crypto"
	"github.com/d5c5ceb0/newchain/storage"
	"github.com/d5c5ceb0/newchain/types"
)

type DbPrefix byte

const (
	BlockHashByHeight DbPrefix = 0x00
	BlockByHash       DbPrefix = 0x01
	CurrentBlock      DbPrefix = 0x02
	AccoutByAddress   DbPrefix = 0x03
)

type Store struct {
	db storage.Database
}

func NewStore(db storage.Database) *Store {
	return &Store{
		db: db,
	}
}

func (this *Store) StoreBlock(block *types.Block) error {
	hash, _ := block.Hash()
	height := block.GetHeight()
	data, err := block.Marshal()
	if err != nil {
		return errors.New("block Marshal error")
	}

	hashPrefix := []byte{byte(BlockByHash)}
	if err := this.db.Put(append(hashPrefix, hash.ToBytes()...), data); err != nil {
		return errors.New("store block error")
	}

	index := make([]byte, 8)
	binary.LittleEndian.PutUint64(index, height)
	heightPrefix := []byte{byte(BlockHashByHeight)}
	if err := this.db.Put(append(heightPrefix, index...), hash.ToBytes()); err != nil {
		return errors.New("store blockhash error")
	}

	return nil
}

func (this *Store) GetBlockByHash(hash *crypto.Digest) (*types.Block, error) {
	prefix := []byte{byte(BlockByHash)}
	data, err := this.db.Get(append(prefix, hash.ToBytes()...))
	if err != nil {
		return nil, err
	}

	var b types.Block
	err = b.Unmarshal(data)
	return &b, err
}

func (this *Store) GetBlockByHeight(height uint64) (*types.Block, error) {
	index := make([]byte, 8)
	binary.LittleEndian.PutUint64(index, height)
	prefix := []byte{byte(BlockHashByHeight)}
	hash, err := this.db.Get(append(prefix, index...))
	if err != nil {
		return nil, err
	}
	var h crypto.Digest
	copy(h[:], hash)
	return this.GetBlockByHash(&h)
}

func (this *Store) GetCurrentBlock() (*types.Block, error) {
	prefix := []byte{byte(CurrentBlock)}
	data, err := this.db.Get(prefix)
	if err != nil {
		return nil, err
	}

	var b types.Block
	err = b.Unmarshal(data)
	return &b, err
}

func (this *Store) GetAccount() {
}
