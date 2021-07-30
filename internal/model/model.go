package model

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
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
