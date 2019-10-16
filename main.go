// ShockWave PRO BBS.
// Author: Timothy Lorens (tlorens@cyberdyne.org)

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
)

const (
	CR = byte('\r')
	LF = byte('\n')
)

const (
	cmdMode     = 1
	cmdLineMode = 34
	cmdSE       = 240
	cmdNOP      = 241
	cmdData     = 242

	cmdBreak = 243
	cmdGA    = 249
	cmdSB    = 250

	cmdWill = 251
	cmdWont = 252
	cmdDo   = 253
	cmdDont = 254

	cmdIAC = 255
)

const (
	optEcho            = 1
	optSuppressGoAhead = 3
)

var config ConfigRec

func main() {
	ClrScr()
	config := Init()
	Writeln("|01-|09=|11] |15ShockWavE:PRO BBS |11[|09=|01-|07")
	Writeln("System: " + config.SystemName)

	// Command line argument -port
	portPtr := flag.Int("port", config.BindPort, "Port number")
	flag.Parse()

	// Only accept so many connections.
	maxConns := int(float64(runtime.NumCPU()) * 1.25)
	runtime.GOMAXPROCS(maxConns)

	// Listen for incoming connections.
	listener, err := net.Listen("tcp", config.BindIP+":"+strconv.Itoa(*portPtr))
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer listener.Close()

	fmt.Printf("Listening on %s:%s (max: %d)\n", config.BindIP, strconv.Itoa(*portPtr), maxConns)
	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		fmt.Printf("Connection from: %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleRequest(conn)
	}
}

func hangup(conn net.Conn) {
	conn.Close()
	fmt.Printf("%s Disconnected\n", conn.RemoteAddr())
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// defer hangup(conn)
	c, _ := NewClient(conn)

	c.InitClient()

	c.SetMenu("start.mnu")
	c.RunMenu()
	c.conn.Close()
	fmt.Printf("%s Disconnected\n", c.conn.RemoteAddr())
}
