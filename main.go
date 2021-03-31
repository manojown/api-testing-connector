package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/manojown/connector/model"
	"github.com/manojown/connector/service"
)

var knt int

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
	msg.Ack()
	var dataRecieved model.Payload
	_ = json.Unmarshal(msg.Payload(), &dataRecieved)
	task := service.ActionFinder(dataRecieved)
	task.Responder, _ = os.Hostname()
	// data := ResponseHandler(msg.Payload())
	out, _ := json.Marshal(task)
	fmt.Println("called in mqtt subsciber :::")
	// text := fmt.Sprintf("this is result msg #%d!", knt)
	// log.Println("DOne", text)
	knt++
	// str := fmt.Sprintf("nn/sensors/%s", msg.Topic())
	str := fmt.Sprintf("nn/sensors/1")

	fmt.Println("Topic is", str)
	token := client.Publish(str, 0, false, out)
	_ = token.Wait()
	if token.Error() != nil {
		log.Fatal(token.Error()) // Use your preferred logging technique (or just fmt.Printf)
	}

}

func main() {
	l := fmt.Println
	knt = 0
	topic := os.Args[1]
	userName := os.Args[2]
	password := os.Args[3] //"server/XVlBz"
	fmt.Println("userName", userName, password)
	//"server/XVlBz"
	log.Println("Topic", topic)
	var sentData model.Payload
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sentData.Server.ServerIP, _ = os.Hostname()
	sentData.Server.Topic = topic
	sentData.Action = "CONNECTION"

	opts := MQTT.NewClientOptions().AddBroker("mqtts://c4bfab5c0b76403f8b3a0259ed4ccdc3.s1.eu.hivemq.cloud:8883")
	opts.SetClientID("mac-go")
	opts.SetPassword(password)
	opts.SetUsername(userName)
	opts.SetDefaultPublishHandler(f)

	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 2, nil); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to server\n")
	}
	data, err := json.Marshal(sentData)
	if err != nil {
		l("error : ", err)
	}
	// fmt.Println("dat is", data, sentData)
	client.Publish("nn/sensors/1", 0, false, data)
	<-c
}
