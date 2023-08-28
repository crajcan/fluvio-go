package offset

type Offset interface {
	isOffset()
}

type OffsetFromBeginning struct {
	Value int64
}

func (_ OffsetFromBeginning) isOffset() {}

type OffsetFromEnd struct {
	Value int64
}

func (_ OffsetFromEnd) isOffset() {}

type OffsetAbsolute struct {
	Value int64
}

func (_ OffsetAbsolute) isOffset() {}

func NewOffsetFromBeginning(value uint32) Offset {
	return &OffsetFromBeginning{Value: int64(value)}
}

func NewOffsetFromEnd(value uint32) Offset {
	return &OffsetFromEnd{Value: int64(value)}
}

func NewOffsetAbsolute(value int64) Offset {
	return &OffsetAbsolute{Value: value}
}
