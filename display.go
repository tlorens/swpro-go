package main

import (
	"bufio"
	"os"
	"log"
	"strings"
	"time"
)

func (c *Client) MCIPrint(str string) {
	currentTime := time.Now()
	out := str
	for i := 0; i < len(out); i++ {
		if (string(out[i]) == "|") {
			cmd := string(out[i+1]) + string(out[i+2])
			switch cmd {
				case "HK":
					out = strings.Replace(out, "|HK", "", -1)
					c.HitKey()
			}
		}
	}
    out = strings.Replace(out, "|01", SetColor(1,0), -1)
    out = strings.Replace(out, "|02", SetColor(2,0), -1)
    out = strings.Replace(out, "|03", SetColor(3,0), -1)
    out = strings.Replace(out, "|04", SetColor(4,0), -1)
    out = strings.Replace(out, "|05", SetColor(5,0), -1)
    out = strings.Replace(out, "|06", SetColor(6,0), -1)
    out = strings.Replace(out, "|07", SetColor(7,0), -1)
    out = strings.Replace(out, "|08", SetColor(8,0), -1)
    out = strings.Replace(out, "|09", SetColor(9,0), -1)
    out = strings.Replace(out, "|10", SetColor(10,0), -1)
    out = strings.Replace(out, "|11", SetColor(11,0), -1)
    out = strings.Replace(out, "|12", SetColor(12,0), -1)
    out = strings.Replace(out, "|13", SetColor(13,0), -1)
    out = strings.Replace(out, "|14", SetColor(14,0), -1)
    out = strings.Replace(out, "|15", SetColor(15,0), -1)
    out = strings.Replace(out, "|CR", "\n", -1)
    out = strings.Replace(out, "|HK", "", -1)
    out = strings.Replace(out, "|DA", currentTime.Format("01-02-2006"), -1)
    out = strings.Replace(out, "|TI", currentTime.Format("3:4:5 pm"), -1)
    c.NetWrite(out)
}

func (c *Client) MCIPrintLn(str string) {
	c.MCIPrint(str + "\n")
}

/**
 *
 *  Private method to read strings from files.
 *
 */
func freadln(r *bufio.Reader) (string, error) {
    var (isPrefix bool = true
        err error = nil
        line, ln []byte
    )

  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }

  return string(ln), err
}

/**
 *
 *  Print a file to the screen.
 *
 */
func (c *Client) PrintFile(filename string) int {
    fh, err := os.Open(filename)

    if err != nil {
        log.Printf("error opening file: %v\n",err)
        return -1
    }

    reader := bufio.NewReader(fh)

    line, err := freadln(reader)

    for err == nil {
        c.MCIPrint(line)
        line, err = freadln(reader)
    }

    c.NetWrite("\n")
    return 1
}

func (c *Client) GetLine(maxLen int) string {
    var tmp string
    var ch byte

    for ch != '\n' {
        ch = c.NetReadCh()
		if (len(tmp) <= maxLen) {
			tmp = tmp + string(ch)
		}
	}

	return tmp
}

func (c *Client) Prompt(maxLen int, prompt string) string {
    log.Printf("Menu Prompt" + prompt)
	c.MCIPrint(prompt)
	return strings.TrimRight(strings.TrimSpace(c.GetLine(maxLen)), "\n")
}
