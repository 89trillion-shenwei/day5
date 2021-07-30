package httpserver

import (
	"day5/internal/Ws"
	"day5/internal/router"
	"log"
	"net/http"
)

//服务器初始化
func Init() {
	go Ws.H.Run()
	router.WebsokectRouter()
	//监听
	err := http.ListenAndServe(":8811", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
