package service

import (
	"log"

	"github.com/valyala/fasthttp"
)

func APICall(url string, method string, data []byte) {

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethodBytes([]byte(method))
	req.SetBody(data)
	resp := fasthttp.AcquireResponse()
	err := myClient.Do(req, resp)
	statusCode := resp.StatusCode()
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)
	if err != nil {
		log.Println("Network error.")
	}

	if statusCode == fasthttp.StatusOK || statusCode == fasthttp.StatusMovedPermanently {
		log.Println("Request sent successfully.")

	} else {
		log.Println("Something went wrong from server please check server config.")
	}

}

// func ActionFinder(payload model.Payload) model.Payload {
// 	var newPayload model.PayloadReciver
// 	action := payload.Action
// 	switch action {
// 	case "REQUEST":
// 		conf := payload.Conf
// 		newPayload.Action = payload.Action
// 		newPayload.Uid = payload.Uid
// 		newPayload.TestResponse = ResponseHandler(conf)
// 		return newPayload
// 	default:
// 		return payload
// 	}
// }

// func ResponseHandler(conf model.Configuration) model.TestResponse {
// 	var response model.TestResponse
// 	responseReciever := make(chan model.TestResponse, 1)
// 	log := log.New(os.Stdout, "ResponseHandler: ", log.LstdFlags)
// 	// err := json.Unmarshal(data, &conf)
// 	// go createConfig(db, conf)
// 	log.Println("conf is", conf)
// 	// if err != nil {
// 	// 	log.Println("Error while json parse", err.Error())
// 	// }
// 	start := time.Now()
// 	go Initialize(&conf, responseReciever)
// 	response = <-responseReciever
// 	fmt.Println("REspond recieved")
// 	t := time.Now()
// 	response.TotalTimeTaken = int64(math.Ceil(t.Sub(start).Seconds()))
// 	response.PerSecond = response.SucessRequests / response.TotalTimeTaken
// 	response.ReadThroughput = response.ReadThroughput / response.TotalTimeTaken
// 	response.WriteThroughput = response.WriteThroughput / response.TotalTimeTaken
// 	response.Url = conf.URL
// 	// responseArray = append(responseArray, response)
// 	return response
// }
