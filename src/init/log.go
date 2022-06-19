package init

import (
	"fmt"
	"log"
	"os"
)

func initLog(logPath string) {
	fmt.Println("initing log system at: ", logPath)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetOutput(logFile)
	log.SetPrefix("[PS]")
	if err == nil {
		fmt.Println("Log init succeed")
	}
}
