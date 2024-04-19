package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 将指定HTTP 升级为 WebSocket连接的参数
var upgrader = websocket.Upgrader{ReadBufferSize: 1024,
	WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool {
		return true
	}}
var webConn *websocket.Conn

func handler(w http.ResponseWriter, r *http.Request) {
	// 调用将http 升级为websockt
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	webConn = conn
	// 开一个线程处理 websocket 中的数据
	go func() {
		for {
			// 阻塞读取
			messageType, p, err := webConn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			if err := webConn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}()
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func main() {
	http.HandleFunc("/ws", handler)
	http.HandleFunc("/hello", helloWorld)
	http.ListenAndServe(":9999", nil)
}
