package service

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/manojown/connector/model"
)

func ActionFinder(payload model.Payload) model.Payload {
	var newPayload model.Payload
	action := payload.Action
	switch action {
	case "REQUEST":
		conf := payload.Conf
		newPayload.Action = payload.Action
		newPayload.Uid = payload.Uid
		newPayload.TestResponse = ResponseHandler(conf)
		return newPayload
	default:
		return payload
	}
}

func ResponseHandler(conf model.Configuration) model.TestResponse {
	var response model.TestResponse
	responseReciever := make(chan model.TestResponse, 1)
	log := log.New(os.Stdout, "ResponseHandler: ", log.LstdFlags)
	// err := json.Unmarshal(data, &conf)
	// go createConfig(db, conf)
	log.Println("conf is", conf)
	// if err != nil {
	// 	log.Println("Error while json parse", err.Error())
	// }
	start := time.Now()
	go Initialize(&conf, responseReciever)
	response = <-responseReciever
	fmt.Println("REspond recieved")
	t := time.Now()
	response.TotalTimeTaken = int64(math.Ceil(t.Sub(start).Seconds()))
	response.PerSecond = response.SucessRequests / response.TotalTimeTaken
	response.ReadThroughput = response.ReadThroughput / response.TotalTimeTaken
	response.WriteThroughput = response.WriteThroughput / response.TotalTimeTaken
	response.Url = conf.URL
	// responseArray = append(responseArray, response)
	return response
}
