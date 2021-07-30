package Ws

import (
	"day5/internal/message1"
	"day5/internal/model"
	"day5/logUtil"
	"time"
)

const (
	writeWait      = 360 * time.Second
	pongWait       = 360 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

//服务端中心
type Hub struct {
	Conns       map[*Connection]bool //用户连接
	Broadcastss chan Message         //广播通道
	Register    chan Message         //登录通道
	Unregister  chan Message         //退出通道
	Ping        chan Message         //心跳检测通道
}

//初始化服务器对象
var H = Hub{
	Conns:       make(map[*Connection]bool), //连接
	Broadcastss: make(chan Message),         //广播
	Register:    make(chan Message),         //登录
	Unregister:  make(chan Message),         //退出
	Ping:        make(chan Message),         //心跳
}

func (h *Hub) Run() {

	for {
		select {
		case m := <-H.Register: //登录连接
			//在线
			h.Conns[m.Conn] = true
			list := message1.List{}
			list.Username = m.Username
			model.UserList = append(model.UserList, m.Username)
			data := m.Msg
			m.Conn.Send <- data
			logUtil.Log.Printf(m.Username + "已上线")
			//fmt.Println(m.username+"已上线")

		case m := <-h.Unregister: //注销连接
			//fmt.Println("服务器收到退出请求")
			conns := h.Conns
			if conns != nil {
				if _, ok := conns[m.Conn]; ok {
					logUtil.Log.Printf(m.Username + "已下线")
					//下线
					model.UserList = DeleteSlice1(model.UserList, m.Username)
					h.Conns[m.Conn] = false
					delete(conns, m.Conn) //删除链接
					close(m.Conn.Send)
				}
			}

		case m := <-h.Broadcastss: //传输全员广播信息
			for con := range h.Conns {
				select {
				case con.Send <- m.Msg:
				default:
					close(con.Send)
					delete(h.Conns, con)
				}
			}

		case m := <-h.Ping: //心跳检测
			data := m.Msg
			m.Conn.Send <- data
		}
	}
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

/*func main1() {

	go h.run()
	http.HandleFunc("/",serverWs )
	//监听
	err := http.ListenAndServe(":8811", nil)
	if err != nil {
		logUtil.Fatal("ListenAndServe: ", err)
	}
}*/
