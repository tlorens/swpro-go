package main

import "fmt"

// ClrScr will clear local screen and home cursor.
func ClrScr() {
	fmt.Print(ClearScr())
}

// Writeln will write a string to the local console with a CR/LF
func Writeln(str string) {
	fmt.Println(str)
}

// Write will write a string to the local console.
func Write(str string) {
	fmt.Print(str)
}

// WriteXy will place the cursor in x, y coordinates
// then write the string to the console
func WriteXy(str string, x int, y int) {
	fmt.Print(GotoXy(x, y))
	fmt.Print(str)
}
