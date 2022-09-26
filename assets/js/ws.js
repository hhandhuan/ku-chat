$(function () {
    if (window.user !== undefined) {
        let addr = 'ws://127.0.0.1:8080/ws?cid=' + user.cid
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

    $('.add-friend-btn').on('click', function () {
        // let self = $(this)
        // self.attr("disabled", true)
        // $.ajax({
        //     url: '/search',
        //     type: 'get',
        //     data: {user: $('.add-friend-input').val()},
        //     dataType: 'json',
        //     success: function (res) {
        //         self.attr("disabled", false)
        //         if (res.code !== 0) {
        //             alert(res.msg)
        //         } else {
        //
        //         }
        //     },
        //     error: function (res) {
        //         console.log(res)
        //     }
        // })
        $.ajax({
            url: '/add-friend',
            type: 'POST',
            data: {target_id: 7, remark: "你好呀"},
            dataType: 'json',
            success: function (res) {
                self.attr("disabled", false)
                if (res.code !== 0) {
                    alert(res.msg)
                } else {

                }
            },
            error: function (res) {
                console.log(res)
            }
        })
    })
})