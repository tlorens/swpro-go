package main

import (
	"strings"
)

func Match(str1 string, str2 string) bool {
	if (len(strings.Trim(str1," ")) == len(strings.Trim(str2, " "))) {
		if (0 == strings.Compare(strings.ToUpper(str1), strings.ToUpper(str2))) {
			return true;
		}
	}
	return false;
}

func MatchRaw(str1 string, str2 string) bool {
	if (len(strings.Trim(str1," ")) == len(strings.Trim(str2, " "))) {
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
}

func (c *Client) LogOff(param1 string, param2 string) {
	if ("ASK" == param1 || "ASK" == param2) {

	}
	c.Disconnect()
}
