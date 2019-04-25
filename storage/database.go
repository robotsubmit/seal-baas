package storage

import "errors"

type DBType int

const (
	LevelDB DBType = 1
)

type Database interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Close() error
	Commit() error
}

func NewDb(typ DBType, file string) (Database, error) {
	switch typ {
	case LevelDB:
		return NewLevelDB(file)
	}

	return nil, errors.New("no database.")
}
