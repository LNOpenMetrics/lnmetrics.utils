package db

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Database struct {
	instance *leveldb.DB
	pathDb   *string
}

var instance Database

// Deprecated: The db implementation was designed to be a Singleton
// However, this need to be a client rules, so we need to
// give the possibility to build an instance with `NewInstance` method
// and if the client want to use the Singleton pattern it need to
// implement the pattern.
func GetInstance() *Database {
	return &instance
}

func NewInstance(homeDir string) (*Database, error) {
	instance := Database{}
	if err := instance.InitDB(homeDir); err != nil {
		return nil, err
	}
	return &instance, nil
}

func (this *Database) Ready() bool {
	return this.instance != nil
}

func (this *Database) InitDB(homedir string) error {
	path := homedir + "/db"
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return err
	}
	this.instance = db
	this.pathDb = &path
	return nil
}

// Put a string value in the db with the specified key
func (this *Database) PutValue(key string, value string) error {
	return this.PutValueInBytes(key, []byte(value))
}

// Put a value defined in bytes in the db with the specified key
func (this *Database) PutValueInBytes(key string, value []byte) error {
	return this.instance.Put([]byte(key), value, nil)
}

// Get a value as string with the specified key, or return a not nil error
// in case of error.
func (this *Database) GetValue(key string) (string, error) {
	value, err := this.GetValueInBytes(key)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// Get a value as bytes with the specified key, or return a not nil error.
func (this *Database) GetValueInBytes(key string) ([]byte, error) {
	value, err := this.instance.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// Delete value with the specified key
func (this *Database) DeleteValue(key string) error {
	return this.instance.Delete([]byte(key), nil)
}

// Iterate through the index from startKey included to endKey not included, and execute the callback defined from the user.
func (this *Database) IterateThrough(startKey string, endKey string, callback func(string) error) error {
	iter := this.instance.NewIterator(&util.Range{Start: []byte(startKey), Limit: []byte(endKey)}, nil)
	for iter.Next() {
		if err := callback(string(iter.Value())); err != nil {
			return err
		}
	}
	return nil
}

// Iterate through the index from startKey, and run the callback specified by the user,
// a not nil value is returned if any error occurs.
func (this *Database) IterateFrom(startKey string, callback func(string) error) error {
	iter := this.instance.NewIterator(util.BytesPrefix([]byte(startKey)), nil)
	for iter.Next() {
		if err := callback(string(iter.Value())); err != nil {
			return err
		}
	}
	return nil
}

// Take as result the raw iterator to iterate over the database content
// The user need to take kare of the Release at the end of the usage
func (this *Database) GetRawIterator() iterator.Iterator {
	return this.instance.NewIterator(nil, nil)
}

// Take as result the list of keys that are in the stored in the database
func (this *Database) ListOfKeys() ([]*string, error) {
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
func (this *Database) CloseDatabase() error {
	return this.instance.Close()
}

// Erase the root of the database
func (this *Database) EraseDatabase() error {
	return os.RemoveAll(*this.pathDb)
}

// Close and erase the database
func (this *Database) EraseAfterCloseDatabse() error {
	if err := this.CloseDatabase(); err != nil {
		return err
	}
	return this.EraseDatabase()
}
