package store

import (
	"Server/internal/queues"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
	"time"
)

type queue struct {
	connection *amqp.Connection
	queue      amqp.Queue
	channel    *amqp.Channel
	mu         sync.RWMutex
	eventMap   map[string]struct{}
}

type event struct {
	uuid      string
	message   string
	timestamp time.Time
}

func NewQueue() queues.Queue {
	return &queue{}
}

func (e event) GetEventUUID() string {
	return e.uuid
}
func (q *queue) Init() {
	time.Sleep(15 * time.Second)
	conn, err := amqp.Dial("amqp://rmuser:rmpassword@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %q", err)
	}

	q.connection = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %q", err)
	}

	q.channel = ch

	queue, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	q.queue = queue

	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %q", err)
	}

	err = ch.QueueBind(
		queue.Name,
		"",
		"amq.fanout",
		false,
		nil,
	)

	q.eventMap = make(map[string]struct{})
}

func (q *queue) Close() {
	_ = q.connection.Close()
	_ = q.channel.Close()
}

func (q *queue) AddConsumer(f func(e queues.Event)) {
	messages, err := q.channel.Consume(
		q.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %q", err)
	}

	go func() {
		for message := range messages {
			var e event
			err := json.Unmarshal(message.Body, &e)
			if err != nil {
				log.Fatalf("failed to unmarshal an e. Error: %q", err)
			}
			f(e)
		}
	}()
}

func (q *queue) Publish(message []byte) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id := uuid.New().String()
	event := event{uuid: id, message: string(message), timestamp: time.Now()}
	q.mu.Lock()
	q.eventMap[id] = struct{}{}
	q.mu.Unlock()
	jsonString, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("failed to marshal an event. Error: %q", err)
	}
	err = q.channel.PublishWithContext(ctx,
		"amq.fanout",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonString,
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %q", err)
	}
}
