package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "log"
    "encoding/json"
)

var CurMenuFile = "start.mnu"
var PrevMenuFile = ""
var StringSet struct{}

const (
    CONN_HOST = "192.168.1.3"
    CONN_PORT = "9000"
    CONN_TYPE = "tcp"
)

func LoadStrings() {
    var strings StringSetsRec
    stringFile, err := os.Open("data/strings.json")
    defer stringFile.Close()

    if err != nil {
        log.Println(err.Error())
    }

    jsonParser := json.NewDecoder(stringFile)
    jsonParser.Decode(&strings)

    log.Println(strings.Sets[0].Strings)
}

func main() {

    LoadStrings()
    // Listen for incoming connections.
    listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer listener.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        fmt.Printf("Connection from: %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
        Connection = conn
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
    New(bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)))
    CurMenuFile = "start.mnu"
    RunMenu()
    conn.Close()
    fmt.Printf("%s Disconnected\n", conn.RemoteAddr())
}
