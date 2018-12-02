package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"sync"
)
var (
	wg sync.WaitGroup
)
func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.44.128:9092"}, nil)
	// 快捷写法 string.Split("192.168.14.4:9092", ",")初步估计会返回一个[]string{}
	if err != nil {
		fmt.Println("Failed start consumer", err)
	}
	partitionsList, err := consumer.Partitions("testTopic")
	if err != nil {
		fmt.Println("Failed to get the list of partitons", err)
	}
	fmt.Println("Partitions list:", partitionsList)
	for partition := range partitionsList {
		pc, err := consumer.ConsumePartition("testTopic", int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer){
			wg.Add(1)
			for msg:= range pc.Messages(){
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
