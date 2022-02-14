package Middlewares

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//log to file
func LoggerToFile() gin.HandlerFunc {

	logFilePath := "logs"
	logFileName := "quik-assessment-log"
	//config.LOG_FILE_NAME

	//Log file
	fileName := path.Join(logFilePath, logFileName)

	//Write to file
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//Instantiate
	logger := logrus.New()

	//Set output
	logger.Out = src

	//Set the log level
	logger.SetLevel(logrus.DebugLevel)

	//Set the log format
	logger.SetFormatter(&logrus.TextFormatter{})

	return func(c *gin.Context) {
		//Starting time
		startTime := time.Now()

		//Process the request
		c.Next()

		//End Time
		endTime := time.Now()

		//execution time
		latencyTime := endTime.Sub(startTime)

		//request method
		reqMethod := c.Request.Method

		//request routing
		reqUri := c.Request.RequestURI

		//status code
		statusCode := c.Writer.Status()

		//Request IP
		clientIP := c.ClientIP()

		//log format
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
