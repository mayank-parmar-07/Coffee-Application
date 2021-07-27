package config

import (
	"fmt"
	"log"
	"os"

	"github.com/natefinch/lumberjack"
)

/*logging setup*/
func SetupLogging() *log.Logger {
	e, err := os.OpenFile("./coffee-application.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	errLog := log.New(e, "", log.Ldate|log.Ltime)
	errLog.SetOutput(&lumberjack.Logger{
		Filename:   "./coffee-application.log",
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})
	return errLog
}
