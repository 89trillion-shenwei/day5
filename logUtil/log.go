package logUtil

import (
	"log"
	"os"
)

var Log *log.Logger

func Init() {
	file, err := os.OpenFile("./logUtil/log.logUtil", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0766)
	if err != nil {
		log.Fatalln("日志打开失败," + err.Error())
	}
	Log = log.New(file, "service:", log.Ldate|log.Ltime|log.Lshortfile)
}
