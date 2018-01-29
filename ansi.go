/**
 *
 *	ShockWave PRO BBS ANSI Output library.
 *
 * 	Author: Timothy Lorens (tlorens@cyberdyne.org)
 *
 */
package main

import (
	"fmt"
	"math"
)

const ESC = "\033["

// Storage variables for preventing echoing of color esc
// sequences multiple times if the colors haven't changed.
var curFg int = 7  // Default forground color
var curBg int = 0  // Default background color

/**
 *
 *	Used 'Get Cursor postion' to detect ANSI
 *
 */
func DetectANSI() string {
	return ESC + "6n"
}

/**
 *
 *	Reset text color to terminal default.
 *
 */
func ColorReset() string {
	return ESC + "0m"
}

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
	return fmt.Sprintf(ESC + "%dJ", i)
}


func clearline(i int) string {
	return fmt.Sprintf(ESC + "%dK", i)
}

/**
 *
 *	Clear entire screen and home cursor (1,1)
 *
 * 	Since ^[[2J doesn't work in bash/linux.
 *
 */
func ClearScr() string {
	return fmt.Sprintf("%s %s", clear(2), GotoXY(1,1))
}

/**
 *
 *	Clears entire line.
 *
 */
func ClearLine() string {
	return clearline(2)
}

/**
 *
 *	Put the cursor at X, Y coordinates
 *
 */
func GotoXY(x int, y int) string {
	return fmt.Sprintf(ESC + "%d;%dH", x, y)
}

/**
 *
 *	Move cursor up i number of lins
 *
 */
func CursorUp(i int) string {
	return fmt.Sprintf(ESC + "%dA", i)
}

/**
 *
 *	Move cursor down i number of lines.
 *
 */
func CursorDn(i int) string {
	return fmt.Sprintf(ESC + "%dB", i)
}

/**
 *
 *	Move cursor right i number of columns
 *
 */
func CursorRt(i int) string {
	return fmt.Sprintf(ESC + "%dC", i)
}

/**
 *
 *	Move cursor left i number of columns
 *
 */
func CursorLf(i int) string {
	return fmt.Sprintf(ESC + "%dD", i)
}

/**
 *
 *	Save cursors X,Y coordinates
 *
 */
func CursorSave() string {
	return fmt.Sprintf(ESC + "s")
}

/**
 *
 *	Restore / Goto saved cursor coordinates.
 *
 */
func CursorRestore() string {
	return fmt.Sprintf(ESC + "u")
}

/**
 *
 *	Set forground and background color.
 *
 *	integer f Forground color value 0-15
 *	integer b background color value 0-7
 *
 */
func SetColor(f int, b int) string {
	var tmp string
	var colors = [8]int {0,4,2,6,1,5,3,7}

	// Prevent sending color codes if they haven't changed.
	if curFg != f || curBg != b {
		curFg = f
		curBg = b

		fg := colors[int(math.Mod(float64(f), 8))] + 30
		bg := colors[int(math.Mod(float64(b), 8))] + 40

		if b > 7 {
			tmp += "5;"
		} else {
			tmp += "0;"
		}

		if f > 7 {
			tmp += "1;"
		}

		return fmt.Sprintf(ESC + "%s%d;%dm", tmp, fg, bg)
	}
	return ""
}
