package init

import (
	"fmt"
	"os"
)

func initLog(logPath string) *os.File {
	fmt.Println("initing log system at: ", logPath)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		os.Create(logPath)
		logFile, _ = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		return logFile
	} else {
		fmt.Println("Log init succeed")
		return logFile
	}
}
