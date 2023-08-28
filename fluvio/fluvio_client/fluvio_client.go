package fluvio_client

import (
	"unsafe"

	"github.com/avinassh/fluvio-go/fluvio/c_interface"
	"github.com/avinassh/fluvio-go/fluvio/consumer/partition_consumer"
	"github.com/avinassh/fluvio-go/fluvio/fluvio_error"
	"github.com/avinassh/fluvio-go/fluvio/topic_producer"
)

func Hello() string {
	return "Hello, world."
}

type Fluvio struct {
	wrapper *c_interface.FluvioWrapper
}

func Connect() (*Fluvio, error) {
	errPtr := c_interface.FluvioErrorNew()
	defer c_interface.FluvioErrorFree(errPtr)
	f := c_interface.FluvioConnect(errPtr)

	message := c_interface.FluvioErrorMsg(errPtr)
	if message != nil {
		return nil, fluvio_error.NewFluvioError(c_interface.GoString(message))
	}
	return &Fluvio{
		wrapper: f,
	}, nil

}

func (f *Fluvio) TopicProducer(topic string) (*topic_producer.TopicProducer, error) {
	topicPtr := c_interface.CString(topic)
	defer c_interface.Free(unsafe.Pointer(topicPtr))
	errPtr := c_interface.FluvioErrorNew()
	defer c_interface.FluvioErrorFree(errPtr)
	t := c_interface.FluvioTopicProducer(f.wrapper, topicPtr, errPtr)

	message := c_interface.FluvioErrorMsg(errPtr)
	if message != nil {
		return nil, fluvio_error.NewFluvioError(c_interface.GoString(message))
	}
	return &topic_producer.TopicProducer{
		Wrapper: t,
	}, nil
}

func (f *Fluvio) Close() {
	c_interface.FluvioFree(f.wrapper)
}

func (f *Fluvio) PartitionConsumer(topic string, partition int32) (*partition_consumer.PartitionConsumer, error) {
	topicPtr := c_interface.CString(topic)
	defer c_interface.Free(unsafe.Pointer(topicPtr))
	errPtr := c_interface.FluvioErrorNew()
	defer c_interface.FluvioErrorFree(errPtr)
	consumer := c_interface.PartitionConsumer(f.wrapper, topicPtr, c_interface.Int32_t(partition), errPtr)

	message := c_interface.FluvioErrorMsg(errPtr)
	if message != nil {
		return nil, fluvio_error.NewFluvioError(c_interface.GoString(message))
	}
	return &partition_consumer.PartitionConsumer{Wrapper: consumer}, nil
}
