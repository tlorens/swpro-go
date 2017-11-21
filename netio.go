package main

import (
    "bufio"
    "net"
)

type Client struct {
    conn net.Conn
    rw *bufio.ReadWriter
}

func NewClient(c net.Conn) (*Client, error) {
    rwIO := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
    return &Client{
        conn: c,
        rw: rwIO,
    }, nil
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
    ch, _ := c.rw.ReadByte()
    return ch
}

func (c *Client) Disconnect() {
    c.conn.Close()
}
