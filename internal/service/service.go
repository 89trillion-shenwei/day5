package service

import (
	"day5/internal/message1"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
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

//判断重名
func Check(s string, strs []string) bool {
	for i := 0; i < len(strs); i++ {
		if s == strs[i] {
			return false
		}
	}
	return true
}
