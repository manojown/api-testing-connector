package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/manojown/connector/app/handler"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func Initialize() {
	app := new(App)
	app.Router = mux.NewRouter()
	app.setRouter()
	app.run()
}

func (app *App) run() {
	signalChennal := make(chan os.Signal, 1)
	signal.Notify(signalChennal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	go func() {
		handler := cors.Default().Handler(app.Router)
		log.Println("Server started on:3004")
		http.ListenAndServe(":3004", handler)
	}()
	sig := <-signalChennal
	fmt.Println("Signal recieved", sig)
}

func (app *App) setRouter() {

	app.apiHandler("/test", "POST", handler.StartServices)
	app.apiHandler("/ping", "GET", handler.Ping)

	// app.apiHandler("/history/{id}", "GET", handler.GetHistory)
	// app.apiHandler("/history", "DELETE", handler.DeteteHistory)
	// app.apiHandler("/getCount", "GET", handler.GetHistoryTotal)
	// app.apiHandler("/config", "GET", handler.GetConfiguration)
	// app.apiHandler("/server", "Post", handler.CreateServerHandler)
	// app.apiHandler("/server", "Get", handler.GetServerHandler)
	// // app.apiHandler("/request/{id}", "Get", handler.GetAllRequest)

	// app.apiHandler("/config", "POST", handler.CreateConfiguration)
	// app.apiHandler("/request", "POST", handler.NewSessionRequest)
	// app.apiHandler("/", "GET", handler.CheckServer)
}

func (app *App) apiHandler(path string, method string, handler handlerFunction) {
	app.Router.HandleFunc(path, app.funcHandler(handler)).Methods(method)
}

type handlerFunction func(w http.ResponseWriter, r *http.Request)

func (app *App) funcHandler(handler handlerFunction) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		handler(rw, r)
	}
}
