package storage

var (
	BitsPerKey    = 10
	OpenFileLimit = 64
)

func NewLevelDB(file string) (*LevelDB, error) {
	opts := opt.Options{
		NoSync:                 false,
		Filter:                 filter.NewBloomFilter(BITSPERKEY),
		OpenFilesCacheCapacity: OpenFileLimit,
	}

	db, err := leveldb.OpenFile(file, &opts)

	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}

	if err != nil {
		return nil, err
	}

	return &LevelDB{
		fn: file,
		db: db,
	}, nil
}

type LevelDB struct {
	fn string
	db *leveldb.DB
}

func (self *LevelDB) Put(key []byte, value []byte) error {
	return self.db.Put(key, value, nil)
}

func (self *LevelDB) Get(key []byte) ([]byte, error) {
	dat, err := self.db.Get(key, nil)
	return dat, err
}

func (self *LevelDB) Delete(key []byte) error {
	return self.db.Delete(key, nil)
}

func (self *LevelDB) Close() error {
	if err := self.Commit(); err != nil {
		return err
	}

	return self.db.Close()
}

func (self *LevelDB) Commit() error {
	return nil
}
