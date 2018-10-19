package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"go-test/internal/model"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"encoding/json"
	"time"
)

var server *http.Server
var signalChan chan os.Signal

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

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
	infoLogger.Printf("记录当前pid:%d", pid)
}

func LogInit() {

	fileName := strconv.Itoa(time.Now().Year()) + "-" + strconv.Itoa(int(time.Now().Month())) + "-" + strconv.Itoa(time.Now().Day()) + "-" + strconv.Itoa(time.Now().Hour()) + ".log"
	checkFileAndAutoCreate("Log/error")
	checkFileAndAutoCreate("Log/warn")
	checkFileAndAutoCreate("Log/info")
	errFile, err := os.OpenFile("Log/error/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	warnFile, err := os.OpenFile("Log/warn/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	infoFile, err := os.OpenFile("Log/info/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	infoLogger = log.New(io.MultiWriter(os.Stderr, infoFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(io.MultiWriter(os.Stderr, warnFile), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func checkFileAndAutoCreate(filePath string)  {
	_,err := os.Stat(filePath)
	if err != nil{
		os.MkdirAll(filePath,0755)
	}
}
