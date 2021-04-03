package handler

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/manojown/connector/model"
	"github.com/manojown/connector/service"
)

func Ping(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(`{"message":"connected"}`)
}
func StartServices(rw http.ResponseWriter, r *http.Request) {
	var payload model.PayloadReciver
	var payloadResponder model.PayloadResponder
	var conf model.Configuration
	responseReciever := make(chan model.TestResponse, 1)
	log := log.New(os.Stdout, "StartTesting: ", log.LstdFlags)
	err := json.NewDecoder(r.Body).Decode(&payload)
	conf = payload.Conf
	if err != nil {
		log.Println("Error while json parse", err.Error())
	}
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
		payloadResponder.UID = payload.UID
		payloadResponder.Responder = payload.Ip
		log.Println("Ip is", payload.Ip)
		payloadResponder.TestResponse = response
		// responseArray = append(responseArray, response)
		dataToSent, err := json.Marshal(payloadResponder)
		if err != nil {
			log.Println("Something went wrong while marshal in StartServices.")
		}
		service.APICall("https://manoj-api-testing.herokuapp.com/result", "POST", dataToSent)
	}()
	json.NewEncoder(rw).Encode("Test request accepted.")

}

func Connect(rw http.ResponseWriter, r *http.Request) {

}
