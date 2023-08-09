- module fluvio
  - `func Connect() (*Fluvio, error)`❌1 
  - Fluvio struct❌2
    - `func (f *Fluvio) TopicProducer(topic string) (TopicProducer, error)`❌3
    - `func (f *Fluvio) ConsumerConfigWithWasmFilter(wasmFile string) (*ConsumerConfig, error)`❌4
    - `func (f *Fluvio) PartitionConsumer(topic string, partition int32) (*PartitionConsumer, error)`❌5
    - `func (f *Fluvio) Close()`✅
  - TopicProducer struct❌6
    - `func (t *TopicProducer) Send(key, value []byte) error`❌7
    - `func (t *TopicProducer) SendString(key, value String) error`❌8
    - `func (t *TopicProducer) Close()`
  - PartitionConsumer struct❌9
    - `func (pc *PartitionConsumer) Stream(offset Offset) (*PartitionConsumerStream, error)`✅
    - `func (pc *PartitionConsumer) StreamWithConfig(offset Offset, config *ConsumerConfig) (*PartitionConsumerStream, error)`✅
    - `func (pc *PartitionConsumer) Close()`✅
  - PartitionConsumerStream struct❌10
    - `func (pcs *PartitionConsumerStream) Next() (*Record, error)`✅
    - `func (pcs *PartitionConsumerStream) Close()`✅
  - ConsumerConfig struct❌11
    - `func (c *ConsumerConfig) Close()`✅

1. Connect() should be inside the fluvio struct file atleast
2. Fluvio Struct should been in a fluvio::fluvio submodule
3. TopicProducer() should be inside a fluvio::fluvio submodule 
4. ConsumerConfigWithWasmFilter should mirror consumer_config_buidler.smartmodule (may need create_smartmodule() and create_smartmodule_from_path() from fluvio/crates/fluvio-cli/src/client/smartmoduleinvocation.rs)
5. PartitionConsumer() should be inside a fluvio::fluvio submodule 
6. TopicProducer struct should be in a file called topic_producer.go
7. TopicProducer.Send() should have some response like ProduceOutput 
8. See if we can remove the need for SendString by using an interface for the value argument in Send
9. PartitionConsumer struct should be in a file called partition_consumer.go 
10. PartitionConsumerStream struct should be in a file called partition_consumer_stream.go
11. ConsumerConfig struct should be in a file called consumer_config.go