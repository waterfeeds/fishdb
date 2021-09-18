package query

type Iterator interface {
	Seek() Iterator

	Next() bool

	At() ([]byte, error)

	Err() error
}

type Batch interface {

}

type Searcher interface {

}

type Query interface {
	Iterator
	Batch
}
