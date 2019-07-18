package storage

import (
	"log"

	"github.com/dgraph-io/badger"
)

type BDBStore struct {
	fn string
	db *badger.DB
}

func NewBadgerDB(file string) (*BDBStore, error) {
	db, err := badger.Open(badger.DefaultOptions(file))
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	return &BDBStore{
		fn: file,
		db: db,
	}, nil
}

func (self *BDBStore) Put(key []byte, value []byte) error {
	err := self.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})

	if err != nil {
		return err
	}
	return nil
}

func (self *BDBStore) Get(key []byte) (value []byte, err error) {

	var valCopy []byte
	err = self.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			if err != nil {
				return err
			}
			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return valCopy, nil
}

func (self *BDBStore) Delete(key []byte) error {
	err := self.db.View(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (self *BDBStore) Close() error {
	return self.db.Close()
}
