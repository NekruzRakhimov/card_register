package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"time"
)

// SetLogger Установка Logger-а
func SetLogger() {
	f, err := os.OpenFile(AppSettings.AppParams.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error while creating file. Error is: ", err.Error())
		return
	}

	logger := &lumberjack.Logger{
		Filename:   f.Name(),
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	}
	gin.DefaultWriter = io.MultiWriter(logger, os.Stdout)
	log.SetOutput(logger)
}

// FormatLogs Форматирование логов
func FormatLogs(r *gin.Engine) {
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("[GIN] %s - [%s] \"%s %s %s %d %s \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
}
