package logger

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/thegorgeouslang/logit"
	"go/build"
	"os"
	"time"
)

var SLog = *logit.Syslog

func init() {
	// appending custom categories
	SLog.AppendCategories(map[string][]string{
		"godotenv": {"Godotenv error:",
			"the godotenv dependency is not working properly"},
	})

	// loading godotenv
	err := godotenv.Load()
	if err != nil {
		// in case of error the path and extension of the log file are being
		// hardcoded
		SLog.Filepath = fmt.Sprintf("%s/%s%s%s", build.Default.GOPATH, "logs/",
			time.Now().Format("2006_01_02"), ".logs")
		// append
		SLog.WriteLog("godotenv",
			"the godotenv dependency is not properly working", SLog.GetTraceMsg())
	} else { // try to get the values of the .env file
		// changing the default log file path
		path := fmt.Sprintf("%s/%s%s%s", build.Default.GOPATH, os.Getenv("logfile_path"),
			time.Now().Format("2006_01_02"), os.Getenv("logfile_ext"),
		)
		SLog.Filepath = path
	}
}
