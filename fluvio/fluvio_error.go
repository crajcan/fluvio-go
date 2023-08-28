package fluvio

import (
	"strings"
)

var ErrNoRecord = NewFluvioError("no records received")

// Used as default when Offset doesn't fit in any of the enumerable types
var ErrInvalidOffsetType = NewFluvioError("received invalid offset type")

type FluvioError struct {
	msg string
}

func (e *FluvioError) Error() string {
	return e.msg
}

func NewFluvioError(msg string) error {
	return &FluvioError{strings.ToLower(msg)}
}
