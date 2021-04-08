package model

type Configuration struct {
	ID        string              `json:"id,omitempty"`
	UserID    string              `json:"userID,omitempty"`
	URL       string              `json:"url"`
	Requests  int64               `json:"requests"`
	Time      int                 `json:"time"`
	Clients   int                 `json:"clients"`
	Headers   []map[string]string `json:"headers"`
	KeepAlive bool                `json:"keepAlive"`
	Method    string              `json:"method"`
	Ips       []string            `json:"ips"`
	PostData  []byte              `json:"postData"`
	Created   int64               `json:"created"`
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
	UserID    string        `json:"userID"`
	RequestID string        `json:"requestID"`
	Responder string        `json:"responder"`
	Conf      Configuration `json:"conf"`
}
type PayloadResponder struct {
	UserID       string       `json:"userID"`
	RequestID    string       `json:"requestID"`
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
