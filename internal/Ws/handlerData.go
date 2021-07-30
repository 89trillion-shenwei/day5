package Ws

import (
	"day5/internal/message1"
	"day5/internal/model"
	"day5/logUtil"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Message struct {
	Msg      message1.Msg //消息
	Conn     *Connection  //连接
	Username string       //用户名
}

//连接信息
type Connection struct {
	Ws   *websocket.Conn   //连接
	Send chan message1.Msg //发送通道
}

//服务端读取数据
func (m Message) ReadPump() {
	c := m.Conn
	defer func() {
		H.Unregister <- m
		c.Ws.Close()
	}()
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error {
		c.Ws.SetReadDeadline(time.Now().Add(pongWait))
		fmt.Println("心跳检测" + time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})
	for {
		//得到发送信息
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		//解析数据

		m.Msg = model.Proto2Struct(message)
		switch m.Msg.MsgType {
		//广播
		case "talk":
			H.Broadcastss <- m
			logUtil.Log.Printf(m.Username + "请求发送广播:" + m.Msg.Msg)
			//得到当前所有在线用户
		case "userlist":
			msg := message1.Msg{
				UserName: m.Username,
				Msg:      List2String(model.UserList),
				MsgType:  "userlist",
			}
			m.Conn.Send <- msg
		//退出
		case "exit":
			H.Broadcastss <- m
			H.Unregister <- m
			//fmt.Println("退出")
			return
		//心跳检测
		case "ping":
			H.Ping <- m
		//登录
		case "login":
			H.Register <- m

		}
	}
}

func (c *Connection) write(mt int, payload []byte) error {
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Ws.WriteMessage(mt, payload)
}

//服务端上传数据
func (s *Message) WritePump() {
	c := s.Conn

	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, model.Struct2proto(message)); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, nil); err != nil {
				logUtil.Log.Printf("pong")
				return
			}
		}
	}
}

//list转字符串
func List2String(strs []string) string {
	var str string
	for _, i := range strs {
		str = str + i + " "
	}
	return str
}

func Strs2List(strs []string) []*message1.List {
	list := make([]*message1.List, 0)
	if len(strs) == 0 {
		return list
	}
	for i := 0; i < len(strs); i++ {
		list[i].Username = strs[i]
	}
	return list
}
