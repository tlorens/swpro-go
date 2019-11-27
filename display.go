package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

// MCIPrint parses 'message control instructions'
// displays the string to the console.
func (c *Client) MCIPrint(str string) {
	currentTime := time.Now()
	out := str
	for i := 0; i < len(out); i++ {
		if string(out[i]) == "|" {
			cmd := string(out[i+1]) + string(out[i+2])
			switch cmd {
			case "HK":
				out = strings.Replace(out, "|HK", "", -1)
				c.HitKey()
			}
		}
	}
	out = strings.Replace(out, "|01", SetColor(1, 0), -1)
	out = strings.Replace(out, "|02", SetColor(2, 0), -1)
	out = strings.Replace(out, "|03", SetColor(3, 0), -1)
	out = strings.Replace(out, "|04", SetColor(4, 0), -1)
	out = strings.Replace(out, "|05", SetColor(5, 0), -1)
	out = strings.Replace(out, "|06", SetColor(6, 0), -1)
	out = strings.Replace(out, "|07", SetColor(7, 0), -1)
	out = strings.Replace(out, "|08", SetColor(8, 0), -1)
	out = strings.Replace(out, "|09", SetColor(9, 0), -1)
	out = strings.Replace(out, "|10", SetColor(10, 0), -1)
	out = strings.Replace(out, "|11", SetColor(11, 0), -1)
	out = strings.Replace(out, "|12", SetColor(12, 0), -1)
	out = strings.Replace(out, "|13", SetColor(13, 0), -1)
	out = strings.Replace(out, "|14", SetColor(14, 0), -1)
	out = strings.Replace(out, "|15", SetColor(15, 0), -1)
	out = strings.Replace(out, "|CR", "\n", -1)
	out = strings.Replace(out, "|HK", "", -1)
	out = strings.Replace(out, "|DA", currentTime.Format("01-02-2006"), -1)
	out = strings.Replace(out, "|TI", currentTime.Format("3:4:5 pm"), -1)
	c.NetWrite(out)
}

// MCIPrintLn will display a 'message control instruction' to the screen.
func (c *Client) MCIPrintLn(str string) {
	c.MCIPrint(str + "\n")
}

// PrintFile will display a file to the screen.
func (c *Client) PrintFile(filename string) {
	fh, err := os.Open(filename)

	if err != nil {
		log.Printf("Error opening file: %v\n", err)
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		// i := strings.LastIndex(line, string(26))
		// log.Println(i)
		c.MCIPrintLn(line)
	}
	c.NetWrite("\r\n")
}

func (c *Client) GetLine(maxLen int) string {
	var tmp string
	var ch byte

	for ch != 13 {
		ch = c.NetReadCh()
		if len(tmp) <= maxLen {
			switch {
			case ch == 13:
				break
			case ch == 127 && len(tmp) > 0:
				c.NetWrite(CursorLf(1))
				c.NetWrite(" ")
				c.NetWrite(CursorLf(1))
				tmp = tmp[:len(tmp)-1]
			case ch >= 32 && ch <= 126:
				tmp += string(ch)
				c.NetWrite(string(ch))
			}
		}
	}
	c.NetWrite("\r\n")
	return strings.Trim(tmp, "\u0000\r\n")
}

// Prompt Displays an input promt and wait for input
func (c *Client) Prompt(maxLen int, prompt string) string {
	c.MCIPrint(prompt)
	return strings.TrimSpace(c.GetLine(maxLen))
}

func (c *Client) KeyPrompt(prompt string) string {
	var ch byte

	c.MCIPrint(prompt)

	for {
		ch = c.NetReadCh()
		if ch > 31 && ch < 255 {
			break
		}
	}

	c.NetWriteln(string(ch))

	return string(ch)
}
