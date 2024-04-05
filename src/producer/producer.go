package main

import (
	"encoding/json"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type QualidadeAr struct {
	idSensor     string
	timestamp    string
	tipoPoluente string
	nivel        float32
}

func createSensor(id string, time string, polutant string, level float32) *QualidadeAr {
	s := &QualidadeAr{
		idSensor:     id,
		timestamp:    time,
		tipoPoluente: polutant,
		nivel:        level,
	}

	return s
}

func (a *QualidadeAr) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func createProducer() *ckafka.Producer{
	producer, err := ckafka.NewProducer(&ckafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"client.id":         "go-producer",
	})

	if err != nil {
		panic(err)
	}

	return producer
}

func main() {

	sensor := "sensor_001, 2024-04-04T12:34:56Z, PM2.5, 35.2"

	producer := createProducer()

	defer producer.Close()

	topic := "qualidadeAr"

	message := sensor

	for {

		producer.Produce(&ckafka.Message{
				TopicPartition: ckafka.TopicPartition{Topic: &topic,},
				Value:         []byte(message),
		}, nil)

		producer.Flush(2 * 1000)
	}

}