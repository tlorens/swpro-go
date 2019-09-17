// ShockWave PRO BBS ANSI Output library.
// Author: Timothy Lorens (tlorens@cyberdyne.org)

package main

import (
	"fmt"
	"math"
)

const esc = "\033["

var curFg = 7 // Default forground color
var curBg = 0 // Default background color

/**
 *
 * 	Private method to clear the terminal
 *
 * 	integer i
 * 		0 Clear from cursor to end of screen;
 * 		1 Clear from cursor to beginning of screen
 * 		2 Clear entire screen and home cursor (1,1) (ANSI.SYS)
 */
func clear(i int) string {
	return fmt.Sprintf(esc+"%dJ", i)
}

func clearline(i int) string {
	return fmt.Sprintf(esc+"%dK", i)
}

// DetectANSI will 'Get Cursor postion' to detect ANSI
func DetectANSI() string {
	return esc + "6n"
}

// ColorReset resets text color to terminal default.
func ColorReset() string {
	return esc + "0m"
}

// ClearScr Clears the entire screen and home cursor (1,1)
// Since ^[[2J doesn't work in bash/linux.
func ClearScr() string {
	return fmt.Sprintf("%s %s", clear(2), GotoXy(1, 1))
}

// ClearLine Clears current line.
func ClearLine() string {
	return clearline(2)
}

// GotoXy Puts the cursor at X, Y coordinates
func GotoXy(x int, y int) string {
	return fmt.Sprintf(esc+"%d;%dH", x, y)
}

// CursorUp will move cursor up i number of lins
func CursorUp(i int) string {
	return fmt.Sprintf(esc+"%dA", i)
}

// CursorDn will move cursor down i number of lines.
func CursorDn(i int) string {
	return fmt.Sprintf(esc+"%dB", i)
}

// CursorRt will move cursor right i number of columns
func CursorRt(i int) string {
	return fmt.Sprintf(esc+"%dC", i)
}

// CursorLf will move cursor left i number of columns
func CursorLf(i int) string {
	return fmt.Sprintf(esc+"%dD", i)
}

// CursorSave saves cursors X,Y coordinates
func CursorSave() string {
	return fmt.Sprintf(esc + "s")
}

// CursorRestore restores saved cursor coordinates.
func CursorRestore() string {
	return fmt.Sprintf(esc + "u")
}

// SetColor sets forground and background color.
//
//	integer f Forground color value 0-15
//	integer b background color value 0-7
func SetColor(foreGround int, backGround int) string {
	var output string
	var colors = [8]int{0, 4, 2, 6, 1, 5, 3, 7}

	// Prevent sending color codes if they haven't changed.
	if curFg != foreGround || curBg != backGround {
		curFg = foreGround
		curBg = backGround

		fg := colors[int(math.Mod(float64(foreGround), 8))] + 30
		bg := colors[int(math.Mod(float64(backGround), 8))] + 40

		switch {
		case backGround > 7:
			output += "5;"
		default:
			output += "0;"
		}

		if foreGround > 7 {
			output += "1;"
		}

		return fmt.Sprintf(esc+"%s%d;%dm", output, fg, bg)
	}

	return ""
}
