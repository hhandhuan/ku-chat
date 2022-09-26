$(function () {
    if (window.user !== undefined) {
        let addr = 'ws://127.0.0.1:8080/ws?cid='+user.cid
        let ws = new WebSocket(addr)
        ws.onerror = function (event) {
            console.log(event)
        }
        ws.onopen = function (event) {
            let data = JSON.stringify({"id": 1, "data": user})
            ws.send(data)
        }
        ws.onclose = function (event) {
            console.log(event)
        }
        ws.onmessage = function (event) {
            console.log(event)
        }
    }
})