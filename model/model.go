package model

type Configuration struct {
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
	URL             string `json:"url"`
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

type PayloadReciver struct {
	UID  int64         `json:"uid"`
	Ip   string        `json:"ip"`
	Conf Configuration `json:"conf"`
}
type PayloadResponder struct {
	UID          int64        `json:"uid"`
	Responder    string       `json:"responder"`
	TestResponse TestResponse `json:"testResponse"`
}

type Server struct {
	Token         string     `json:"token" bson:"token"`
	Port          string     `json:"port" bson:"port"`
	RAM           MemStatus  `json:"ram" bson:"ram"`
	DiskSpace     DiskStatus `json:"DiskSpace" bson:"DiskSpace"`
	CPU           CPUStatus  `json:"cpu" bson:"cpu"`
	LastConnected int64      `json:"lastConnected" bson:"lastConnected"`
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

type CPUStatus struct {
	Cores int     `json:"cores"`
	Usage float64 `json:"usage"`
}
type MemStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	Self uint64 `json:"self"`
}

type Config struct {
	Token string
	URL   string
	Port  string
}
