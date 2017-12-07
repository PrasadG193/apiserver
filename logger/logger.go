package logger

import (
	"os"
	"fmt"
	"github.com/Sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	f, err := os.OpenFile("/var/log/apiserver.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	Log.Out = f
}
