package partition_consumer

import (
	"github.com/avinassh/fluvio-go/fluvio/c_interface"
	"github.com/avinassh/fluvio-go/fluvio/consumer/consumer_config"
	"github.com/avinassh/fluvio-go/fluvio/consumer/partition_consumer_stream"
	"github.com/avinassh/fluvio-go/fluvio/fluvio_error"
	"github.com/avinassh/fluvio-go/fluvio/offset"
)

type PartitionConsumer struct {
	Wrapper *c_interface.PartitionConsumerWrapper
}

func (pc *PartitionConsumer) Stream(offset offset.Offset) (*partition_consumer_stream.PartitionConsumerStream, error) {
	return pc.StreamWithConfig(offset, nil)
}

func (pc *PartitionConsumer) StreamWithConfig(off offset.Offset, config *consumer_config.ConsumerConfig) (*partition_consumer_stream.PartitionConsumerStream, error) {
	errPtr := c_interface.FluvioErrorNew()
	defer c_interface.FluvioErrorFree(errPtr)
	var offsetPtr *c_interface.OffsetWrapper
	switch o := off.(type) {
	case *offset.OffsetFromBeginning:
		offsetPtr = c_interface.OffsetFromBeginning(c_interface.Uint32_t(o.Value))
		defer c_interface.OffsetFree(offsetPtr)
	case *offset.OffsetFromEnd:
		offsetPtr = c_interface.OffsetFromEnd(c_interface.Uint32_t(o.Value))
		defer c_interface.OffsetFree(offsetPtr)
	case *offset.OffsetAbsolute:
		offsetPtr = c_interface.OffsetAbsolute(c_interface.Int64_t(o.Value), errPtr)

		message := c_interface.FluvioErrorMsg(errPtr)
		if message != nil {
			return nil, fluvio_error.NewFluvioError(c_interface.GoString(message))
		}
		defer c_interface.OffsetFree(offsetPtr)
	default:
		return nil, fluvio_error.ErrInvalidOffsetType
	}
	var stream *c_interface.PartitionConsumerStream
	if config == nil {
		stream = c_interface.Stream(pc.Wrapper, offsetPtr, errPtr)
	} else {
		stream = c_interface.StreamWithConfig(pc.Wrapper, offsetPtr, config.Wrapper, errPtr)
	}

	message := c_interface.FluvioErrorMsg(errPtr)
	if message != nil {
		return nil, fluvio_error.NewFluvioError(c_interface.GoString(message))
	}
	return &partition_consumer_stream.PartitionConsumerStream{Wrapper: stream}, nil
}

func (pc *PartitionConsumer) Close() {
	c_interface.PartitionConsumerFree(pc.Wrapper)
}
