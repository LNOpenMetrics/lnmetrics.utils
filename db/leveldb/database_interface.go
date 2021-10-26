package db

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type database struct {
	instance *leveldb.DB
	pathDb   *string
}

var instance database

func GetInstance() *database {
	return &instance
}

func (this *database) Ready() bool {
	return this.instance != nil
}

func (this *database) InitDB(homedir string) error {
	path := homedir + "/db"
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return err
	}
	this.instance = db
	this.pathDb = &path
	return nil
}

func (this *database) PutValue(key string, value string) error {
	return this.instance.Put([]byte(key), []byte(value), nil)
}

func (this *database) GetValue(key string) (string, error) {
	value, err := this.instance.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

func (this *database) DeleteValue(key string) error {
	return this.instance.Delete([]byte(key), nil)
}

// TODO Add method to iterate over a method.

// Take as result the raw iterator to iterate over the database content
// The user need to take kare of the Release at the end of the usage
func (this *database) GetRawIterator() iterator.Iterator {
	return this.instance.NewIterator(nil, nil)
}

// Take as result the list of keys that are in the stored in the database
func (this *database) ListOfKeys() ([]*string, error) {
	iter := this.GetRawIterator()
	defer iter.Release()
	keys := make([]*string, 0)
	for iter.Next() {
		key := string(iter.Key())
		keys = append(keys, &key)
	}
	err := iter.Error()
	return keys, err
}

// Close connection with the database
func (this *database) CloseDatabase() error {
	return this.instance.Close()
}

// Erase the root of the database
func (this *database) EraseDatabase() error {
	return os.RemoveAll(*this.pathDb)
}

func (this *database) EraseAfterCloseDatabse() error {
	if err := this.CloseDatabase(); err != nil {
		return err
	}
	return this.EraseDatabase()
}
