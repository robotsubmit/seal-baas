package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	BitsPerKey    = 10
	OpenFileLimit = 64
)

func NewLevelDB(file string) (*LDBStore, error) {
	opts := opt.Options{
		NoSync:                 false,
		Filter:                 filter.NewBloomFilter(BitsPerKey),
		OpenFilesCacheCapacity: OpenFileLimit,
	}

	db, err := leveldb.OpenFile(file, &opts)

	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}

	if err != nil {
		return nil, err
	}

	return &LDBStore{
		fn: file,
		db: db,
	}, nil
}

type LDBStore struct {
	fn string
	db *leveldb.DB
}

func (self *LDBStore) Put(key []byte, value []byte) error {
	return self.db.Put(key, value, nil)
}

func (self *LDBStore) Get(key []byte) ([]byte, error) {
	dat, err := self.db.Get(key, nil)
	return dat, err
}

func (self *LDBStore) Delete(key []byte) error {
	return self.db.Delete(key, nil)
}

func (self *LDBStore) Close() error {
	if err := self.Commit(); err != nil {
		return err
	}

	return self.db.Close()
}

func (self *LDBStore) Commit() error {
	return nil
}
