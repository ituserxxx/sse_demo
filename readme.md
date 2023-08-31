# SSE技术 （使用http将服务端数据推送到客户端）

SSE 是轻量的 WebSocket，SSE 可以不需要用户执行任何操作直接向客户端发送消息，SSE 是单向通道，发送文本信息。

基于包

```
"gopkg.in/antage/eventsource.v1"
```

示例如下

```go

// 可自定义事件类型
sse.SendEventMessage("now is event1 ", "event1", "111")
time.Sleep(2 * time.Second)
es.SendEventMessage("now is event2 ", "event2", "222")
time.Sleep(2 * time.Second)
es.SendEventMessage("now is no event ", "", "333")
time.Sleep(2 * time.Second)
```

客户端接收

```javascript
   <script type="text/javascript">
        window.addEventListener("DOMContentLoaded", function () {
            var source = new EventSource("/events");

            // 连接成功后会触发 open 事件
            source.addEventListener('open', () => {
                console.log('Connected');
            }, false);

            // 服务器发送信息到客户端时，如果没有 event 字段，默认会触发 message 事件
            source.addEventListener('message', e => {
                console.log(`no event data: ${e.data}`);
            }, false);

            // 自定义 EventHandler，在收到 event1 字段为 event1 的消息时触发
            source.addEventListener('event1', e => {
                console.log(`event1 data: ${e.data}`); // => data: 7
            }, false);
            // 自定义 EventHandler，在收到 event2 字段为 event2 的消息时触发
            source.addEventListener('event2', e => {
                console.log(`event2 data: ${e.data}`); // => data: 7
            }, false);
            // 连接异常时会触发 error 事件并自动重连
            source.addEventListener('error', e => {
                if (e.target.readyState === EventSource.CLOSED) {
                    console.log('Disconnected');
                } else if (e.target.readyState === EventSource.CONNECTING) {
                    console.log('Connecting...');
                }
            }, false)
        })
    </script>
```
