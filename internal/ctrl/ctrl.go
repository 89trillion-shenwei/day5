package ctrl

import (
	"day5/internal/Ws"
	"day5/internal/message1"
	"day5/internal/model"
	"day5/logUtil"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//建立连接后，服务器从header里面拿username
func ServerWs(w http.ResponseWriter, r *http.Request, h Ws.Hub) {
	/*r.ParseForm()
	username:=r.Form["username"][0]*/
	//升级为websocket
	ws, err := model.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logUtil.Log.Printf("升级websocket失败，" + err.Error())
		return
	}
	//初始化连接
	msg := message1.Msg{}
	username := r.Header.Get("username")
	if username == "" {
		log.Fatalln("用户名为空")
	}
	fmt.Println(username)
	msg.MsgType = "login"
	connect := &Ws.Connection{Send: make(chan message1.Msg, 1024), Ws: ws}
	m := Ws.Message{message1.Msg{}, connect, username}
	h.Register <- m
	//开启读写
	go m.WritePump()
	go m.ReadPump()
}

//字节切片转结构体
func Byte2Struct(byts []byte, msg message1.Msg) message1.Msg {
	err := json.Unmarshal(byts, &msg)
	if err != nil {
		fmt.Println("正在解析数据")
		return message1.Msg{}
	}

	return msg
}

//结构体转字节切片
func Struct2byte(msg message1.Msg) ([]byte, error) {
	byts, err := json.Marshal(&msg)
	if err != nil {
		return nil, err
	}
	return byts, nil
}
