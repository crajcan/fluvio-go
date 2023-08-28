package consumer

import "C"
import (
	"unsafe"

	"github.com/avinassh/fluvio-go/fluvio"
	"github.com/avinassh/fluvio-go/fluvio/c_interface"
)

type PartitionConsumerStream struct {
	Wrapper *c_interface.PartitionConsumerStream
}

func (pcs *PartitionConsumerStream) Next() (*Record, error) {
	errPtr := c_interface.FluvioErrorNew()
	defer c_interface.FluvioErrorFree(errPtr)
	result := c_interface.PartitionConsumerStreamNext(pcs.Wrapper, errPtr)
	message := c_interface.FluvioErrorMsg(errPtr)
	if message != nil {
		return nil, fluvio.NewFluvioError(c_interface.GoString(message))
	}
	if result == nil {
		return nil, fluvio.ErrNoRecord
	}
	defer c_interface.RecordFree(result)

	record := &Record{
		Offset: int64(c_interface.RecordOffset(result)),
		Value:  C.GoBytes(unsafe.Pointer(c_interface.RecordValue(result)), C.int(c_interface.RecordValueLen(result))),
	}
	if C.int(c_interface.RecordKeyLen(result)) > 0 {
		record.Key = C.GoBytes(unsafe.Pointer(c_interface.RecordKey(result)), C.int(c_interface.RecordKeyLen(result)))
	}
	return record, nil
}

func (pcs *PartitionConsumerStream) Close() {
	c_interface.PartitionConsumerStreamFree(pcs.Wrapper)
}
