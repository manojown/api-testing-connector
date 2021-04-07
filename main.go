package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/manojown/connector/app"
	"github.com/manojown/connector/model"
)

func main() {
	// getIP()
	var config model.Config

	flag.StringVar(&config.Token, "token", "", "a string")
	flag.StringVar(&config.Port, "port", "3004", "a string")
	flag.StringVar(&config.URL, "url", "http://localhost:8080/connector", "a string")

	flag.Parse()
	if config.Token == "" {
		log.Fatal("Please pass the token by -token=YOUR-TOKEN")
	}

	app.Initialize(config)
	// var sentData model.Payload
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// sentData.Server.ServerIP, _ = os.Hostname()

}

func getIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}
