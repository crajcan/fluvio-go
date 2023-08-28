package consumer

//TODO: Consider moving P

type Record struct {
	Offset int64
	Key    []byte
	Value  []byte
}