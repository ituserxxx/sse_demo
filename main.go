package main

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/antage/eventsource.v1"
)

func main() {
	// 开启一个 sse
	sse := eventsource.New(nil, nil)
	defer sse.Close()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/events", sse)

	//每2秒发送一条当前时间消息，
	go func() {
		for {
			// 自定义2个事件类型
			sse.SendEventMessage("now is event1 ", "event1", "111")
			time.Sleep(2 * time.Second)
			sse.SendEventMessage("now is event2 ", "event2", "222")
			time.Sleep(2 * time.Second)
			sse.SendEventMessage("now is no event ", "", "333")
			time.Sleep(2 * time.Second)
		}
	}()
	log.Println("Open URL http://localhost:8999/ in your browser.")
	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
