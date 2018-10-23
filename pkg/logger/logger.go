package logger

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

//pointer default is nil
var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
)

const (
	errorType = 1
	warnType  = 2
	infoType  = 3
	debugType = 4
)

func getLogger(logType int) *log.Logger {
	switch logType {
	case errorType:
		if errorLogger == nil {
			errorLogger = initLogger("error")
		}
		return errorLogger
	case warnType:
		if warnLogger == nil {
			warnLogger = initLogger("warn")
		}
		return warnLogger
	case infoType:
		if infoLogger == nil {
			infoLogger = initLogger("info")
		}
		return infoLogger
	case debugType:
		if debugLogger == nil {
			debugLogger = initLogger("info")
		}
		return debugLogger
	}
	return nil
}

func checkFileAndAutoCreate(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		error := os.MkdirAll(filePath, 0755)
		return error
	}
	return nil
}

func initLogger(logName string) *log.Logger {
	mkdirFileError := checkFileAndAutoCreate("Log")
	if mkdirFileError != nil {
		return nil
	}
	fileName := strconv.Itoa(time.Now().Year()) + "-" + strconv.Itoa(int(time.Now().Month())) + "-" + strconv.Itoa(time.Now().Day()) + "-" + strconv.Itoa(time.Now().Hour()) + ".log"

	logPath := fmt.Sprintf("Log/%s", logName)
	mkdirFileError = checkFileAndAutoCreate(logPath)
	if mkdirFileError != nil {
		return nil
	}
	file, openFileErr := os.OpenFile(logPath+"/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if openFileErr != nil {
		errorInfo := fmt.Sprintf("打开%s日志文件失败：%s", logName, openFileErr)
		log.Println(errorInfo)
	}
	prefix := fmt.Sprintf("%s:", logName)
	Logger := log.New(io.MultiWriter(os.Stderr, file), prefix, log.Ldate|log.Ltime|log.Lshortfile)
	return Logger
}

func LogInfo(string string, v ... interface{}) error{
	log := getLogger(infoType)
	if log != nil {
		log.Printf(string, v...)
	} else {
		fmt.Println(fmt.Sprintf(string, v...))
	}
}

func LogErr(string string, v ... interface{}) {
	log := getLogger(errorType)
	if log != nil {
		log.Printf(string, v...)
	} else {
		fmt.Println(fmt.Sprintf(string, v...))
	}
}

func LogWarn(string string, v ... interface{}) {
	log := getLogger(warnType)
	if log != nil {
		log.Printf(string, v...)
	} else {
		fmt.Println(fmt.Sprintf(string, v...))
	}
}

func LogDebug(string string, v ... interface{}) {
	log := getLogger(debugType)
	if log != nil {
		log.Printf(string, v...)
	} else {
		fmt.Println(fmt.Sprintf(string, v...))

	}
}
