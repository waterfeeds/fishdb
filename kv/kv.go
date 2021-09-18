package kv

type KV interface {
	Get(key []byte) ([]byte, error)

	Put(key []byte, value []byte) error

	Delete(key []byte) error
}
