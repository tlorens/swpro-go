package main

import (
    "bufio"
)

var (
    rwIO *bufio.ReadWriter
    client *ClientRec
)

func NewClient(cli *ClientRec) {
    client = cli
    rwIO = bufio.NewReadWriter(bufio.NewReader(cli.conn), bufio.NewWriter(cli.conn))
}

func NetWrite(str string) {
    rwIO.WriteString(str)
    rwIO.Flush()
}

func NetWriteln(str string) {
    rwIO.WriteString(str + "\n")
    rwIO.Flush()
}

func NetReadCh() byte {
    ch, _ := rwIO.ReadByte()
    return ch
}

func Disconnect() {
	client.conn.Close()
}
