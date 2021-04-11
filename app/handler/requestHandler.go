package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/manojown/connector/model"
	"github.com/manojown/connector/service"
)

func Ping(config model.Config, rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(`{"message":"connected"}`)
}
func StartServices(config model.Config, rw http.ResponseWriter, r *http.Request) {
	var payload model.PayloadReciver
	var payloadResponder model.PayloadResponder
	var conf model.Configuration
	d, _ := ioutil.ReadAll(r.Body)

	responseReciever := make(chan model.TestResponse, 1)
	log := log.New(os.Stdout, "StartTesting: ", log.LstdFlags)
	err := json.Unmarshal(d, &payload)
	// log.Println("Called start service:payload", payload)
	if err != nil {
		log.Println(err)
	}
	conf = payload.Conf

	go func() {
		start := time.Now()
		go service.Initialize(&conf, responseReciever)
		response := <-responseReciever
		t := time.Now()
		response.TotalTimeTaken = int64(math.Ceil(t.Sub(start).Seconds()))
		response.PerSecond = response.SucessRequests / response.TotalTimeTaken
		response.ReadThroughput = response.ReadThroughput / response.TotalTimeTaken
		response.WriteThroughput = response.WriteThroughput / response.TotalTimeTaken
		response.URL = conf.URL
		payloadResponder.UserID = payload.UserID
		payloadResponder.RequestID = payload.RequestID
		payloadResponder.Responder = payload.Responder
		// log.Println("Ip is", payload.Ip)
		payloadResponder.TestResponse = response
		// responseArray = append(responseArray, response)
		dataToSent, err := json.Marshal(payloadResponder)
		fmt.Println("payloadResponder", payloadResponder)
		if err != nil {
			log.Println("Something went wrong while marshal in StartServices.", err.Error(), dataToSent)
		}
		service.APICall(config.URL+"/result", "POST", dataToSent)
	}()
	json.NewEncoder(rw).Encode("Test request accepted.")

}

func Connect(config model.Config, rw http.ResponseWriter, r *http.Request) {
	service.Polling(config)
}
