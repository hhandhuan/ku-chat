// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"ku-chat/ws"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

var (
	cid      uint32
	core     = ws.NewCore()
	upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
)

func main() {
	flag.Parse()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		if conn, err := upgrader.Upgrade(w, r, nil); err != nil {
			log.Println(err)
			return
		} else {
			cid++
			connection := ws.NewConnection(cid, conn, core)
			core.Add(connection)
			go connection.Reader()
			go connection.Writer()
		}
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
