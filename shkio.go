package main

import (
	"fmt"
	"strings"
	"time"
)

// ClrScr will clear local screen and home cursor.
func ClrScr() {
	fmt.Print(ClearScr())
}

// Writeln will write a string to the local console with a CR/LF
func Writeln(str string) {
	Write(str)
	fmt.Println()
}

// Write will write a string to the local console.
func Write(str string) {
	fmt.Print(MCIPrint(str))
}

// WriteXy will place the cursor in x, y coordinates
// then write the string to the console
func WriteXy(str string, x int, y int) {
	fmt.Print(GotoXy(x, y))
	Write(str)
}

// MCIPrint parses messace control instructions.
func MCIPrint(str string) string {
	config := Init()
	currentTime := time.Now()
	output := str

	for i := 0; i < len(output); i++ {
		if string(output[i]) == "|" {
			cmd := string(output[i+1]) + string(output[i+2])
			switch cmd {
			case "CL":
				output = strings.Replace(output, "|CL", "", -1)
				ClrScr()
			}
		}
	}

	output = strings.Replace(output, "|01", SetColor(1, 0), -1)
	output = strings.Replace(output, "|02", SetColor(2, 0), -1)
	output = strings.Replace(output, "|03", SetColor(3, 0), -1)
	output = strings.Replace(output, "|04", SetColor(4, 0), -1)
	output = strings.Replace(output, "|05", SetColor(5, 0), -1)
	output = strings.Replace(output, "|06", SetColor(6, 0), -1)
	output = strings.Replace(output, "|07", SetColor(7, 0), -1)
	output = strings.Replace(output, "|08", SetColor(8, 0), -1)
	output = strings.Replace(output, "|09", SetColor(9, 0), -1)
	output = strings.Replace(output, "|10", SetColor(10, 0), -1)
	output = strings.Replace(output, "|11", SetColor(11, 0), -1)
	output = strings.Replace(output, "|12", SetColor(12, 0), -1)
	output = strings.Replace(output, "|13", SetColor(13, 0), -1)
	output = strings.Replace(output, "|14", SetColor(14, 0), -1)
	output = strings.Replace(output, "|15", SetColor(15, 0), -1)
	output = strings.Replace(output, "|CR", "\n", -1)
	output = strings.Replace(output, "|BE", "\007", -1)
	output = strings.Replace(output, "|ES", esc, -1)
	output = strings.Replace(output, "|DA", currentTime.Format("01-02-2006"), -1)
	output = strings.Replace(output, "|TI", currentTime.Format("3:4:5 pm"), -1)
	output = strings.Replace(output, "|SY", config.SystemName, -1)
	output = strings.Replace(output, "|SN", config.SysOp, -1)

	return output
}
