package server

import (
	"context"
	"fmt"
	"log"
	"mydata/internal/model"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"encoding/json"
)

var server *http.Server
var signalChan chan os.Signal

func RegisterSignal() {
	signalChan = make(chan os.Signal, 10)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
}

func Start() {
	go func() {
		http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
			response := []byte("你好\r\n")
			writer.Write(response)
		})
		http.HandleFunc("/select", func(writer http.ResponseWriter, request *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("catch")
					fmt.Println(r)
				}
			}()
			testModel := model.GetTestModelInstance()
			data := testModel.Select()
			jsonByte, _ := json.Marshal(data)
			writer.Write(jsonByte)
		})

		server = &http.Server{
			Addr:    ":9090",
			Handler: nil,
		}
		server.ListenAndServe()
	}()
	<-signalChan
	fmt.Println("退出")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.Shutdown(ctx)
}

func LogPid() {
	pid := os.Getpid()
	log.Printf("记录当前pid:%d", pid)
}
