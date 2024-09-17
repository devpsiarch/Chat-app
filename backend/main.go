package main 

import (
    "fmt"
    "net/http"
    "github.com/devpsiarch/Chat-app/pkg/websocket"
)



//def websocket endpoint
func Serve_ws(pool *websocket.Pool,w http.ResponseWriter, r *http.Request){
	fmt.Println("Websocket endpoint !!!!")
    //upgrading to ws
    conn,err := websocket.Upgrade(w,r)
    if err != nil {
        fmt.Fprintf(w, "%+V\n", err)
    }
	
	client := &websocket.Client{
		Conn:conn,
		Pool:pool,
	}
	pool.Register <- client
	client.Read()
}

//the function defined below is a handler 
func setuproutes() {
	pool := websocket.NewPool();
	go pool.Start()
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        Serve_ws(pool, w, r)
    })	
}


func main() {
    fmt.Println("Server running ... hopfully")
    setuproutes()
    http.ListenAndServe(":8080",nil)
}
