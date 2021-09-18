package database

type Database interface {
	Open(path string) (*DB, error)

	Close() error
}

type DBOptions struct {

}

type DB struct {

}
