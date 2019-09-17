package main

import (
	"bufio"
	"log"
	"net"
	"runtime"
	//    "fmt"
)

// Client properties
type Client struct {
	conn     net.Conn
	rw       *bufio.ReadWriter
	curMenu  string
	preMenu  string
	ranAutos bool
	useEcho  bool
}

// NewClient constructs a new client socket.
func NewClient(c net.Conn) (*Client, error) {
	rwIO := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	return &Client{
		conn:     c,
		rw:       rwIO,
		curMenu:  "start.mnu",
		ranAutos: false,
		useEcho:  true,
	}, nil
}

// InitClient initializes the terminal environment.
func (c *Client) InitClient() {
	c.rw.Write([]byte{cmdIAC, cmdWill, optEcho, cmdIAC, cmdWill, optSuppressGoAhead, cmdIAC, cmdWont, cmdLineMode, cmdIAC, cmdWill, optEcho})
	c.rw.Flush()
	c.rw.Discard(5)
	//c.rw.WriteString("\e[8;25;80t")
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
	c.rw.WriteString(str + "\n\r")
	c.rw.Flush()
}

func (c *Client) NetReadCh() byte {
	ch, err := c.rw.ReadByte()
	if err != nil {
		log.Println("Error:", err.Error())
		c.conn.Close()
		runtime.Goexit()
	}

	return ch
}

func (c *Client) HotKey() byte {
	return c.NetReadCh()
}

func (c *Client) Discard() {
	c.rw.Discard(1)
}

func (c *Client) Disconnect() {
	c.conn.Close()
}
