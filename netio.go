package main

import (
    "bufio"
    "net"
)

var (
    Conn *bufio.ReadWriter
    Connection net.Conn
)

func New(conn *bufio.ReadWriter) (error) {
    Conn = conn
    return nil
}

func Write(str string) {
    Conn.WriteString(str)
    Conn.Flush()
}

func WriteLn(str string) {
    Conn.WriteString(str + "\n")
    Conn.Flush()
}

func Disconnect() {
	Connection.Close()
}
