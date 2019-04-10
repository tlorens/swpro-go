package main

import (
    "bufio"
    "net"
    "log"
    "runtime"
    "fmt"
)

type Client struct {
    conn net.Conn
    rw *bufio.ReadWriter
    curMenu string
    preMenu string
    ranAutos bool
    useEcho bool
}

func NewClient(c net.Conn) (*Client, error) {
    rwIO := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
    return &Client{
        conn: c,
        rw: rwIO,
        curMenu: "start.mnu",
        ranAutos: false,
        useEcho: true,
    }, nil
}

func (c *Client) InitClient() {
    c.rw.Write([]byte{cmdIAC, cmdWill, optEcho, cmdIAC, cmdWill, optSuppressGoAhead, cmdIAC, cmdWont, cmdLineMode, cmdIAC, cmdWill, optEcho})
    //c.rw.Write([]byte{cmdIAC, cmdDo, cmdLineMode, cmdIAC, cmdSB, cmdLineMode, cmdMode, 0, cmdIAC, cmdSE, cmdIAC, cmdWill, optEcho})
}

func (c *Client) SetMenu(str string) {
    c.curMenu = str
}

func (c *Client) GetMenu() string {
    return c.curMenu
}

func (c *Client) NetWrite(str string) {
    c.rw.WriteString(str)
    c.rw.Flush()
}

func (c *Client) NetWriteln(str string) {
    c.rw.WriteString(str + "\n")
    c.rw.Flush()
}

func (c *Client) NetReadCh() byte {
    ch, err := c.rw.ReadByte()
    if err != nil {
        log.Println("Error:", err.Error())
        c.conn.Close()
        runtime.Goexit()
    }

    fmt.Printf("%c", ch)
    return ch
}

// func (c *Client) HotKey() byte {
//     ch, _ := c.rw.ReadString("\w")
//     return ch
// }

func (c *Client) Discard() {
    c.rw.Discard(1)
}


func (c *Client) Disconnect() {
    c.conn.Close()
}
