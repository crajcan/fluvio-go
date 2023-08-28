package consumer

import (
	"github.com/avinassh/fluvio-go/fluvio"
	"github.com/avinassh/fluvio-go/fluvio/c_interface"
)

type PartitionConsumer struct {
	Wrapper *c_interface.PartitionConsumerWrapper
}

func (pc *PartitionConsumer) Stream(offset fluvio.Offset) (*PartitionConsumerStream, error) {
	return pc.StreamWithConfig(offset, nil)
}

func (pc *PartitionConsumer) StreamWithConfig(offset fluvio.Offset, config *ConsumerConfig) (*PartitionConsumerStream, error) {
	errPtr := c_interface.FluvioErrorNew()
	defer c_interface.FluvioErrorFree(errPtr)
	var offsetPtr *c_interface.OffsetWrapper
	switch o := offset.(type) {
	case *fluvio.OffsetFromBeginning:
		offsetPtr = c_interface.OffsetFromBeginning(c_interface.Uint32_t(o.Value))
		defer c_interface.OffsetFree(offsetPtr)
	case *fluvio.OffsetFromEnd:
		offsetPtr = c_interface.OffsetFromEnd(c_interface.Uint32_t(o.Value))
		defer c_interface.OffsetFree(offsetPtr)
	case *fluvio.OffsetAbsolute:
		offsetPtr = c_interface.OffsetAbsolute(c_interface.Int64_t(o.Value), errPtr)

		message := c_interface.FluvioErrorMsg(errPtr)
		if message != nil {
			return nil, fluvio.NewFluvioError(c_interface.GoString(message))
		}
		defer c_interface.OffsetFree(offsetPtr)
	default:
		return nil, fluvio.ErrInvalidOffsetType
	}
	var stream *c_interface.PartitionConsumerStream
	if config == nil {
		stream = c_interface.Stream(pc.Wrapper, offsetPtr, errPtr)
	} else {
		stream = c_interface.StreamWithConfig(pc.Wrapper, offsetPtr, config.Wrapper, errPtr)
	}

	message := c_interface.FluvioErrorMsg(errPtr)
	if message != nil {
		return nil, fluvio.NewFluvioError(c_interface.GoString(message))
	}
	return &PartitionConsumerStream{Wrapper: stream}, nil
}

func (pc *PartitionConsumer) Close() {
	c_interface.PartitionConsumerFree(pc.Wrapper)
}
