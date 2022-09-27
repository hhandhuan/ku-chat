$(function () {
    if (window.user !== undefined) {
        let addr = 'ws://127.0.0.1:8080/ws?cid=' + user.cid
        let ws = new WebSocket(addr)
        ws.onerror = function (event) {
            console.log(event)
        }

        ws.onopen = function (event) {
        }

        ws.onclose = function (event) {
            console.log(event)
        }

        ws.onmessage = function (event) {
            let data = JSON.parse(event.data)
            chat.handlers[data.id](data.data)
        }
    }

    var chat = {
        handlers: {
            // 处理申请添加好友通知
            200: function (data) {
                let msgBox = $('.msg-box')
                let h = '<span class="badge bg-danger rounded-pill position-absolute">'
                if (msgBox.find('span').length <= 0) {
                    h += '1'
                    h += '</span>'
                } else {
                    h += parseInt(msgBox.find('span').html()) + 1
                    h += '</span>'
                }
                msgBox.html(h)
            }
        },
    }

    // 搜索好友
    $('.search-btn').on('click', function () {
        let self = $(this)
        self.attr("disabled", true)
        $.ajax({
            url: '/search',
            type: 'get',
            data: {user: $('.search-input').val()},
            dataType: 'json',
            success: function (res) {
                self.attr("disabled", false)
                if (res.code !== 0) {
                    return alert(res.msg)
                }
                if (res.data.id > 0) {
                    let h = '<li class="list-group-item d-flex justify-content-between align-items-center">'
                    h += res.data.name
                    h += '<a href="javascript:;" class="text-decoration-none add-friend-btn" data-id="' + res.data.id + '">申请</a>'
                    h += '</li>'
                    $('.friend-box').html(h)
                } else {
                    let h = '<li class="list-group-item d-flex justify-content-between align-items-center text-secondary">'
                    h += ' 该用户不存在'
                    h += '</li>'
                    $('.friend-box').html(h)
                }
            },
            error: function (res) {
                console.log(res)
            }
        })
    })
    // 点击申请
    $('.friend-box').on('click', '.add-friend-btn', function () {
        $.ajax({
            url: '/record-add',
            type: 'POST',
            data: {target_id: $(this).attr('data-id'), remark: "你好呀"},
            dataType: 'json',
            success: function (res) {
                console.log(res)
            },
            error: function (res) {
                console.log(res)
            }
        })
    })

    $('')
    $.ajax({
        url: '/record-logs',
        type: 'GET',
        dataType: 'json',
        success: function (res) {
            console.log(res)

            if (res.data.length > 0) {
                for (let i = 0; i < res.data.length; i++) {
                    let h = '                        <li class="p-2">\n' +
                        '                            <a href="javascript:;" class="d-flex justify-content-between text-decoration-none">\n' +
                        '                                <div class="d-flex flex-row">\n' +
                        '                                    <div>\n' +
                        '                                        <img src="'+res.data[i].user.avatar+'" alt="avatar" class="d-flex align-self-center me-3 rounded-circle" width="50">\n' +
                        '                                    </div>\n' +
                        '                                    <div class="pt-1">\n' +
                        '                                        <p class="mb-0 text-secondary">'+res.data[i].user.name+'</p>\n' +
                        '                                        <p class="small text-secondary">我是何欢</p>\n' +
                        '                                    </div>\n' +
                        '                                </div>\n' +
                        '                                <div class="pt-1">\n' +
                        '                                    <span class="text-secondary">同意</span>\n' +
                        '                                    <span class="text-secondary">拒绝</span>\n' +
                        '                                </div>\n' +
                        '                            </a>\n' +
                        '                        </li>'
                    console.log(h)
                    $('.record-log-list').append(h)
                }
            }
        },
        error: function (res) {
            console.log(res)
        }
    })
})