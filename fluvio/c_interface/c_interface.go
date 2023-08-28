package c_interface

/*
#cgo LDFLAGS: -L../../src -lfluvio_go
#include "../../src/fluvio_go.h"
*/
import "C"
import (
	"unsafe"
)

type TopicProducerWrapper C.TopicProducerWrapper
type FluvioWrapper C.FluvioWrapper
type FluvioErrorWrapper C.FluvioErrorWrapper
type PartitionConsumerWrapper C.PartitionConsumerWrapper
type PartitionConsumerStream C.PartitionConsumerStream
type ConsumerConfigWrapper C.ConsumerConfigWrapper
type OffsetWrapper C.OffsetWrapper
type RecordWrapper C.RecordWrapper

type Uint8_t C.uint8_t

func FluvioConnect(errPtr *FluvioErrorWrapper) *FluvioWrapper {
	return (*FluvioWrapper)(C.fluvio_connect((*C.FluvioErrorWrapper)(errPtr)))
}
func FluvioTopicProducer(f *FluvioWrapper, topic *C.char, errPtr *FluvioErrorWrapper) *TopicProducerWrapper {
	return (*TopicProducerWrapper)(C.fluvio_topic_producer((*C.FluvioWrapper)(f), topic, (*C.FluvioErrorWrapper)(errPtr)))
}

func FluvioErrorNew() *FluvioErrorWrapper {
	return (*FluvioErrorWrapper)(C.fluvio_error_new())
}

func FluvioErrorFree(errPtr *FluvioErrorWrapper) {
	C.fluvio_error_free((*C.FluvioErrorWrapper)(errPtr))
}

func TopicProducerSend(tp *TopicProducerWrapper, key *Uint8_t, keyLen C.size_t, value *Uint8_t, valueLen C.size_t, errPtr *FluvioErrorWrapper) {
	C.topic_producer_send(
		(*C.TopicProducerWrapper)(tp),
		(*C.uint8_t)(key),
		keyLen,
		(*C.uint8_t)(value),
		valueLen,
		(*C.FluvioErrorWrapper)(errPtr),
	)
}

func TopicProducerFree(tp *TopicProducerWrapper) {
	C.topic_producer_free((*C.TopicProducerWrapper)(tp))
}

func FluvioErrorMsg(errPtr *FluvioErrorWrapper) *C.char {
	return ((*C.FluvioErrorWrapper)(errPtr)).msg
}

func Size_t(size int) C.size_t {
	return C.size_t(size)
}

func Uint32_t(value int64) C.uint32_t {
	return C.uint32_t(value)
}

func Int64_t(value int64) C.int64_t {
	return C.int64_t(value)
}

func Int32_t(value int32) C.int32_t {
	return C.int32_t(value)
}

func Int(value int) C.int {
	return C.int(value)
}

func GoString(str *C.char) string {
	return C.GoString(str)
}

func CString(str string) *C.char {
	return C.CString(str)
}

func Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

func FluvioFree(f *FluvioWrapper) {
	C.fluvio_free((*C.FluvioWrapper)(f))
}

func ConsumerConfigWithWasmFilter(wasmFile *C.char, errPtr *FluvioErrorWrapper) *ConsumerConfigWrapper {
	return (*ConsumerConfigWrapper)(C.consumer_config_with_wasm_filter(wasmFile, (*C.FluvioErrorWrapper)(errPtr)))
}

func OffsetFromBeginning(offset C.uint32_t) *OffsetWrapper {
	return (*OffsetWrapper)(C.offset_from_beginning(offset))
}

func OffsetFromEnd(offset C.uint32_t) *OffsetWrapper {
	return (*OffsetWrapper)(C.offset_from_end(offset))
}

func OffsetAbsolute(offset C.int64_t, errPtr *FluvioErrorWrapper) *OffsetWrapper {
	return (*OffsetWrapper)(C.offset_absolute(offset, (*C.FluvioErrorWrapper)(errPtr)))
}

func OffsetFree(offset *OffsetWrapper) {
	C.offset_free((*C.OffsetWrapper)(offset))
}

func Stream(pc *PartitionConsumerWrapper, offset *OffsetWrapper, errPtr *FluvioErrorWrapper) *PartitionConsumerStream {
	return (*PartitionConsumerStream)(C.partition_consumer_stream((*C.PartitionConsumerWrapper)(pc), (*C.OffsetWrapper)(offset), (*C.FluvioErrorWrapper)(errPtr)))
}

func StreamWithConfig(pc *PartitionConsumerWrapper, offset *OffsetWrapper, config *ConsumerConfigWrapper, errPtr *FluvioErrorWrapper) *PartitionConsumerStream {
	return (*PartitionConsumerStream)(C.partition_consumer_stream_with_config(
		(*C.PartitionConsumerWrapper)(pc),
		(*C.OffsetWrapper)(offset),
		(*C.ConsumerConfigWrapper)(config),
		(*C.FluvioErrorWrapper)(errPtr)),
	)
}

func PartitionConsumer(f *FluvioWrapper, topic *C.char, partition C.int32_t, errPtr *FluvioErrorWrapper) *PartitionConsumerWrapper {
	return (*PartitionConsumerWrapper)(C.fluvio_partition_consumer((*C.FluvioWrapper)(f), topic, partition, (*C.FluvioErrorWrapper)(errPtr)))
}

func PartitionConsumerFree(pc *PartitionConsumerWrapper) {
	C.partition_consumer_free((*C.PartitionConsumerWrapper)(pc))
}

func PartitionConsumerStreamNext(stream *PartitionConsumerStream, errPtr *FluvioErrorWrapper) *RecordWrapper {
	return (*RecordWrapper)(C.partition_consumer_stream_next((*C.PartitionConsumerStream)(stream), (*C.FluvioErrorWrapper)(errPtr)))
}

func PartitionConsumerStreamFree(stream *PartitionConsumerStream) {
	C.partition_consumer_stream_free((*C.PartitionConsumerStream)(stream))
}

func RecordFree(record *RecordWrapper) {
	C.record_free((*C.RecordWrapper)(record))
}

func RecordOffset(record *RecordWrapper) C.int64_t {
	return (*C.RecordWrapper)(record).offset
}

func RecordValue(record *RecordWrapper) *Uint8_t {
	return (*Uint8_t)((*C.RecordWrapper)(record).value)
}

func RecordValueLen(record *RecordWrapper) C.size_t {
	return (*C.RecordWrapper)(record).value_len
}

func RecordKey(record *RecordWrapper) *Uint8_t {
	return (*Uint8_t)((*C.RecordWrapper)(record).key)
}

func RecordKeyLen(record *RecordWrapper) C.size_t {
	return (*C.RecordWrapper)(record).key_len
}

func ConsumerConfigFree(config *ConsumerConfigWrapper) {
	C.consumer_config_free((*C.ConsumerConfigWrapper)(config))
}