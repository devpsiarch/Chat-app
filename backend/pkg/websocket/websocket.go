package websocket



import (
    "fmt"
    "io"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)



func Check(err error){
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

//upgrader function
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn,error) {
    ws,err := upgrader.Upgrade(w,r,nil);
    if err != nil {
        log.Println(err)
        return ws,err
    }
    //succsses upgrading to websocket
    return ws,nil
}

//this bad boy will listen for messages sent to out websocket
//needs a websocket connections
func Reader(conn *websocket.Conn){
    for{
        //receives the message
        message_type,p,err := conn.ReadMessage()
        Check(err)
        //lets the messgae
        fmt.Println(string(p))
        err = conn.WriteMessage(message_type,p)
        Check(err)
    }
}

func Writer(conn *websocket.Conn){
    for {
        fmt.Println("Sending")
        //getting the error
        messageType, r, err := conn.NextReader()
        Check(err)
        //setting up the writer
        w, err := conn.NextWriter(messageType)
        Check(err)
        if _, err := io.Copy(w, r); err != nil {
            fmt.Println(err)
            return
        }
        if err := w.Close(); err != nil {
            fmt.Println(err)
            return
        }
    }
} 
