package topic_producer

/*
#cgo LDFLAGS: -L../../src -lfluvio_go
#include "../../src/fluvio_go.h"
*/
import "C"
import (
	"unsafe"

	"github.com/avinassh/fluvio-go/fluvio/c_interface"
	"github.com/avinassh/fluvio-go/fluvio/fluvio_error"
)

type TopicProducer struct {
	Wrapper *c_interface.TopicProducerWrapper
}

func (t *TopicProducer) Send(key, value []byte) error {
	errPtr := c_interface.FluvioErrorNew()
	defer c_interface.FluvioErrorFree(errPtr)
	c_interface.TopicProducerSend(
		t.Wrapper,
		(*c_interface.Uint8_t)(unsafe.Pointer(&key[0])),
		c_interface.Size_t(len(key)),
		(*c_interface.Uint8_t)(unsafe.Pointer(&value[0])),
		c_interface.Size_t(len(value)),
		errPtr,
	)

	message := c_interface.FluvioErrorMsg(errPtr)
	// return nil if errPtr.msg is nil
	if message == nil {
		return nil
	}

	// return message
	return fluvio_error.NewFluvioError(c_interface.GoString(message))
}

func (t *TopicProducer) SendString(key, value string) error {
	return t.Send([]byte(key), []byte(value))
}

func (t *TopicProducer) Close() {
	c_interface.TopicProducerFree(t.Wrapper)
}


