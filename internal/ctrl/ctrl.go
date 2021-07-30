package ctrl

import (
	"day5/internal/Ws"
	"day5/internal/message1"
	"day5/internal/model"
	"day5/internal/service"
	"day5/logUtil"
	"fmt"
	"log"
	"net/http"
)

//建立连接后，服务器从header里面拿username
func ServerWs(w http.ResponseWriter, r *http.Request, h Ws.Hub) {
	username := r.Header.Get("username")
	if username == "" {
		log.Fatalln("用户名为空")
	}
	fmt.Println(username)
	//判断重名
	if !service.Check(username, model.UserList) {
		return
	}
	//升级为websocket
	ws, err := model.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logUtil.Log.Printf("升级websocket失败，" + err.Error())
		return
	}
	//初始化连接
	msg := message1.Msg{}
	msg.MsgType = "login"
	connect := &Ws.Connection{Send: make(chan message1.Msg, 1024), Ws: ws}
	m := Ws.Message{message1.Msg{}, connect, username}
	h.Register <- m
	//开启读写
	go m.WritePump()
	go m.ReadPump()
}
