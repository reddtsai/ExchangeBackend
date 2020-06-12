package model

import (
	"errors"

	"github.com/nsqio/go-nsq"
)

type messageQueue struct {
	producer *nsq.Producer
	q        chan messageBody
}

type messageBody struct {
	topic string
	body  []byte
}

func newNSQ() (*messageQueue, error) {
	c := nsq.NewConfig()
	producer, err := nsq.NewProducer("192.168.1.73:4150", c)
	if err != nil {
		return nil, err
	}
	queue := &messageQueue{
		producer: producer,
		q:        make(chan messageBody, 128),
	}
	go queue.publish()

	return queue, nil
}

func (m *messageQueue) push(topic string, data []byte) error {
	mb := messageBody{
		topic: topic,
		body:  data,
	}
	select {
	case m.q <- mb:
	default:
		return errors.New("message body was full")
	}
	return nil
}

func (m *messageQueue) publish() {
	for {
		b, ok := <-m.q
		if !ok {
			// TODO
			break
		}
		m.producer.Publish(b.topic, b.body)
	}
}

func (m *messageQueue) close() {
	close(m.q)
	m.producer.Stop()
}
