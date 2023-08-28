package consumer_config

import (
	"unsafe"

	"github.com/avinassh/fluvio-go/fluvio/c_interface"
	"github.com/avinassh/fluvio-go/fluvio/fluvio_error"
)

type ConsumerConfig struct {
	Wrapper *c_interface.ConsumerConfigWrapper
}

func ConsumerConfigWithWasmFilter(wasmFile string) (*ConsumerConfig, error) {
	wasmFilePtr := c_interface.CString(wasmFile)
	defer c_interface.Free(unsafe.Pointer(wasmFilePtr))
	errPtr := c_interface.FluvioErrorNew()
	defer c_interface.FluvioErrorFree(errPtr)
	config := c_interface.ConsumerConfigWithWasmFilter(wasmFilePtr, errPtr)

	message := c_interface.FluvioErrorMsg(errPtr)
	if message != nil {
		return nil, fluvio_error.NewFluvioError((c_interface.GoString(message)))
	}
	return &ConsumerConfig{Wrapper: config}, nil
}

func(c *ConsumerConfig) Close() {
	c_interface.ConsumerConfigFree(c.Wrapper)
}