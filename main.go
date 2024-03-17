package main

import (
	"Server/internal/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	//logger := logging.GetLogger()
	//logger.Info("START")
	//cfg := config.GetConfig()
	//time.Sleep(15 * time.Second)
	//conn, err := amqp.Dial("amqp://rmuser:rmpassword@rabbitmq:5672/") // Создаем подключение к RabbitMQ
	//if err != nil {
	//	log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	//}
	//
	//defer func() {
	//	_ = conn.Close() // Закрываем подключение в случае удачной попытки
	//}()
	//
	//ch, err := conn.Channel()
	//if err != nil {
	//	log.Fatalf("failed to open channel. Error: %s", err)
	//}
	//
	//defer func() {
	//	_ = ch.Close() // Закрываем канал в случае удачной попытки открытия
	//}()
	//
	//queue, err := ch.QueueDeclare(
	//	"",    // name
	//	false, // durable
	//	false, // delete when unused
	//	false, // exclusive
	//	false, // no-wait
	//	nil,   // arguments
	//)
	//queueName := queue.Name
	//if err != nil {
	//	log.Fatalf("failed to declare a queue. Error: %s", err)
	//}
	//
	//err = ch.QueueBind(
	//	queueName,    // queue name
	//	"",           // routing key
	//	"amq.fanout", // exchange
	//	false,
	//	nil,
	//)
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//body := "Hello World!"
	//err = ch.PublishWithContext(ctx,
	//	"amq.fanout", // exchange
	//	"",           // routing key
	//	false,        // mandatory
	//	false,        // immediate
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body:        []byte(body),
	//	})
	//if err != nil {
	//	log.Fatalf("failed to publish a message. Error: %s", err)
	//}
	//
	//log.Printf(" [x] Sent %s\n", body)
	//log.Printf(" [x] QueueName %s\n", queueName)
	//
	//messages, err := ch.Consume(
	//	queueName, // queue
	//	"",        // consumer
	//	true,      // auto-ack
	//	false,     // exclusive
	//	false,     // no-local
	//	false,     // no-wait
	//	nil,       // args
	//)
	//if err != nil {
	//	log.Fatalf("failed to register a consumer. Error: %s", err)
	//}
	//
	//go func() {
	//	for message := range messages {
	//		log.Printf("received a message: %s\n", message.Body)
	//	}
	//}()

	queue := store.NewQueue()
	defer queue.Close()
	queue.Init()
	queue.AddConsumer(func(message []byte) {
		log.Printf("message: %s\n", message)
	})

	router := httprouter.New()
	handler := store.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
