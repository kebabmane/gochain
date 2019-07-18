package storage

type DBType int

type Database interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Close() error
}

func NewDb(file string) (Database, error) {
	return NewBadgerDB(file)
}
