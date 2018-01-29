package main

import (
	"strings"
)

func Match(str1 string, str2 string) bool {
	if (len(strings.Trim(str1, " ")) == len(strings.Trim(str2, " "))) {
		if (0 == strings.Compare(strings.ToUpper(str1), strings.ToUpper(str2))) {
			return true;
		}
	}
	return false;
}

func MatchRaw(str1 string, str2 string) bool {
	if (len(strings.Trim(str1, " ")) == len(strings.Trim(str2, " "))) {
		if (0 == strings.Compare(str1, str2)) {
			return true;
		}
	}
	return false;
}


func (c *Client) GetPassword(pword string) bool {
	password := c.Prompt(80, "Password: ")
	if (MatchRaw(pword, password)) {
		return true
	}
	return false
}

func (c *Client) HitKey() {
	c.MCIPrint("|11[|03Hit a key|11]")
	_ = c.NetReadCh()
	// c.Discard()
	c.MCIPrint(ClearLine())
}

func (c *Client) LogOff(param1 string, param2 string) {
	if ("ASK" == param1 || "ASK" == param2) {

	}
	c.Disconnect()
}

func (c *Client) YesNo(prompt string, def bool) bool {
	c.MCIPrint(prompt)
	var input = strings.TrimRight(strings.TrimSpace(c.GetLine(3)), "\n")
	if (strings.Trim(strings.ToUpper(input), "\n") == "Y") {
		return true
	}
	return false
}

func (c *Client) SetEcho() {
	useEcho := c.Prompt(3, "|11Enable echo |03(|15Y|08/|15n|03)|08: ")
	if (strings.ToUpper(useEcho) == "Y") {
		c.MCIPrintLn("|15ECHO set: TRUE|CR|HK")
		c.useEcho = true
	}
	c.MCIPrintLn("|15ECHO set: FALSE|CR|HK")
	c.useEcho = false

}
