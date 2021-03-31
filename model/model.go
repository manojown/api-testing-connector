package model

type Configuration struct {
	Uid       int                 `json:"uid"`
	URL       string              `json:"url"`
	Requests  int64               `json:"requests"`
	Time      int                 `json:"time"`
	Clients   int                 `json:"clients"`
	Headers   []map[string]string `json:"headers"`
	KeepAlive bool                `json:"keepAlive"`
	Method    string              `json:"method"`
	PostData  string              `json:"postData"`
	Created   string              `json:"created"`
	Action    string              `json:"action"`
	Topic     string              `json:"topic"`
	ServerIP  string              `json:"serverIP"`
}

type TestResponse struct {
	Url             string `json:"url"`
	TotalTimeTaken  int64  `json:"totalTime"`
	TotalRequests   int64  `json:"totalRequests"`
	SucessRequests  int64  `json:"sucessRequests"`
	FailedRequests  int64  `json:"failedRequest"`
	NetworkFailed   int64  `json:"networkFailed"`
	WriteThroughput int64  `json:"writeThroughput"`
	ReadThroughput  int64  `json:"readThroughput"`
	PerSecond       int64  `json:"perSecond"`
	Action          string `json:"action"`
	Topic           string `json:"topic"`
	ServerIP        string `json:"serverIP"`
}

type Payload struct {
	Uid          int64               `json:"uid"`
	Responder    string              `json:"responder"`
	Action       string              `json:"action"`
	Server       ConnectConfirmation `json:"server"`
	Conf         Configuration       `json:"conf"`
	TestResponse TestResponse        `json:"testResponse"`
}
type ConnectConfirmation struct {
	Topic    string `json:"topic"`
	ServerIP string `json:"serverIP"`
}
