package main 

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
) 

func check(err error){
    if err != nil {
        log.Println(err)
        return
    }
}

//this is our upgrader it will upgarade the connection to websocket (ws)
//needs reader and a readbuffer
var upgrader =  websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    //impliment checking to allow requests from react frontend
    //for this allows all connections 
    CheckOrigin: func(r *http.Request) bool {return true},
}

//this bad boy will listen for messages sent to out websocket
//needs a websocket connections
func reader(conn *websocket.Conn){
    for{
        //receives the message 
        message_type,p,err := conn.ReadMessage()
        check(err)
        //lets the messgae
        fmt.Println(string(p))
        err = conn.WriteMessage(message_type,p)
        check(err)
    }
}

//the function defined below is a handler 
func setuproutes() {
    http.HandleFunc("/",
    func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "hello from simple server")
    })
    http.HandleFunc("/ws",ws_endpoint)
}

//def websocket endpoint
func ws_endpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.Host)
    //upgrading to ws
    ws,err := upgrader.Upgrade(w,r,nil)
    check(err)
    //read from ws
    reader(ws)
}


func main() {
    fmt.Println("Server running ... hopfully")
    setuproutes()
    http.ListenAndServe(":8080",nil)
}
