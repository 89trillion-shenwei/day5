package model

import (
	"day5/internal/message1"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var UserList []string

const (
	writeWait = 10 * time.Second

	pongWait = 360 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

//升级器参数
var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Proto2Struct(byts []byte) message1.Msg {
	var msg message1.Msg
	err := proto.Unmarshal(byts, &msg)
	if err != nil {
		fmt.Println(err.Error())
		return msg
	}

	return msg
}

// Struct2proto 解析为proto格式
func Struct2proto(msg message1.Msg) []byte {
	byts, err := proto.Marshal(&msg)
	if err != nil {
		fmt.Println("解析为proto格式失败")
	}
	return byts
}

//删除切片中的值
func DeleteSlice1(a []string, s string) []string {
	ret := make([]string, 0, len(a))
	for _, val := range a {
		if val != s {
			ret = append(ret, val)
		}
	}
	return ret
}
