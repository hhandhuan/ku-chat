$(function () {
    let addr = 'ws://localhost:8080/ws?cid=user_1'
    let ws = new WebSocket(addr)
    ws.onerror = function (event) {
        console.log(event)
    }
    ws.onopen = function (event) {
        let data = JSON.stringify({"id": 1, "data": {"uuid": "token123", "username":"eric", "avatar": "avatar.png"}})
        ws.send(data)
    }
    ws.onclose = function (event) {
        console.log(event)
    }
    ws.onmessage = function (event) {
        console.log(event)
    }
})