package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"models"
	"os"
	"strconv"
	"time"
)

var rabbitChannel *amqp.Channel

// Load config setting from json
func initConfig(conf *models.ConfWorker) error{
	file, _ := os.Open("env.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("[initConfig]: decoder.Decode(&conf) Error:", err)
		return err
	}
	err = file.Close()
	if err != nil {
		fmt.Println("[initConfig]: file.Close() Error:", err)
		return err
	}
	return nil
}

func main() {
	workerConf := models.ConfWorker{}
	errInitConfig := initConfig(&workerConf)
	if errInitConfig != nil {
		fmt.Println("ERROR: initConfig(&gameConf):", errInitConfig)
		return
	}

	conn, errAmqpDial := amqp.Dial(workerConf.AmqpHost)
	if errAmqpDial != nil {
		fmt.Println("Failed to connect to RabbitMQ:", workerConf.AmqpHost, ",Error:", errAmqpDial)
		return
	}
	defer conn.Close()

	var err error
	rabbitChannel, err = conn.Channel()
	if err != nil {
		fmt.Printf("Create Channel failed, error: %v\n", err)
		return
	}

	ch, errChannel := conn.Channel()
	if errChannel != nil {
		fmt.Println("Failed to open a channel", errChannel)
		return
	}
	defer ch.Close()

	chString := make(chan string)

	for i:=0; i<len(workerConf.NameOfWorkerQueue); i++ {
		num := workerConf.NumberOfWorkerQueue[i]
		for num > 0 {
			<-time.After(100 * time.Millisecond)	// 减少并发，避免出错
			go ListenQueue(workerConf.NameOfWorkerQueue[i], num, chString)
			num--
		}
	}

	for r := range chString {
		fmt.Println(r)
	}

	rabbitChannel.Close()
	fmt.Println("Program exit abnormal.")
}

// 队列侦听
func ListenQueue(tag string, workerId int, ch chan<- string) {

	queue, errQueueDeclare := rabbitChannel.QueueDeclare(tag, false, true, false, false, nil)
	if errQueueDeclare != nil {
		fmt.Println("[ListenQueue]: QueueDeclare error:", errQueueDeclare)
	}

	fmt.Println("WorkerID: worker", strconv.Itoa(workerId))
	msg, err := rabbitChannel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("[ListenQueue]: rabbitChannel.Consume error:", err)
		ch<- "error"
		return
	}

	for d := range msg {
		doNothing()
		d.Ack(false)
	}

	ch <- "finish"
}

func doNothing() {}