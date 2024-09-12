package main 

import (
    "fmt"
    "net/http"
    "github.com/devpsiarch/Chat-app.git/websocket"
) 

//def websocket endpoint
func Serve_ws(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.Host)
    //upgrading to ws
    ws,err := upgrader.Upgrade(w,r,nil)
    if err != nil {
        fmt.Fprintf(w, "%+V\n", err)
    }
    go websocket.Writer(ws)
    websocket.Reader(ws)
}

//the function defined below is a handler 
func setuproutes() {
    http.HandleFunc("/ws",Serve_ws)
}


func main() {
    fmt.Println("Server running ... hopfully")
    setuproutes()
    http.ListenAndServe(":8080",nil)
}
