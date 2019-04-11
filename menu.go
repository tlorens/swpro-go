package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var menu MenuRec
var prevMenu MenuRec

func loadMenu(file string) MenuRec {
	menuFile, err := os.Open("menudata/" + file)

	log.Println("Loading Menu: " + file)
	if err != nil {
		log.Println(err.Error())
	}
	defer menuFile.Close()

	byteValue, _ := ioutil.ReadAll(menuFile)
	var loadedMenu MenuRec
	json.Unmarshal(byteValue, &loadedMenu)

	return loadedMenu
}

func (c *Client) RunMenu() {
	for {
		// Load initial menu
		menu := loadMenu(c.curMenu)

		if !c.ranAutos {
			c.RunAutos(menu.Commands)
		}

		// RunAutos could possible set a new menu.
		menu = loadMenu(c.curMenu)

		ch := c.Prompt(80, menu.Prompt)
		for _, cmd := range menu.Commands {
			if Match(ch, cmd.Key) {
				c.RunCommand(cmd)
			}
		}
	}
}

/**
 * Execute commands tagged with AUTO to run ONCE automatically before the menut prompt
 */
func (c *Client) RunAutos(menuCommands []CommandRec) {
	for _, cmd := range menuCommands {
		if Match(cmd.Param1, "AUTO") || Match(cmd.Param2, "AUTO") {
			log.Println("Auto: " + cmd.Name + " :: (" + cmd.Param1 + ") (" + cmd.Param2 + ")")
			c.RunCommand(cmd)
		}
	}
}

/**
 * Execute commands tagged with 'EVERY' at every prompt. ]
 */
func (c *Client) RunEvery(menuCommands []CommandRec) {
	// var curMenu = ""
	// for _, cmd := range menuCommands {
	//     if (Match(cmd.Param1, "EVERY") || Match(cmd.Param2, "EVERY")) {
	//         log.Println("EVERY: " + cmd.Cmd)
	//         curMenu = c.RunCommand(cmd)
	//     }
	// }
}

/**
 * Execute a nenus command.
 * @param cmd CommandRec Command to execute.
 */
func (c *Client) RunCommand(cmd CommandRec) {
	log.Println("Running Command: " + cmd.Cmd)
	switch strings.ToUpper(cmd.Cmd) {
	case "ECHO":
		c.SetEcho()
	case "CLEAR":
		c.NetWrite(ClearScr())
	case "@MENU":
		if Match(cmd.Param1, "AUTO") || Match(cmd.Param1, "EVERY") {
			c.SetMenu(cmd.Param2)
		} else {
			c.SetMenu(cmd.Param1)
		}
	case "PRINT":
		if "" != cmd.Param1 {
			log.Println("::" + cmd.Param1)
			c.PrintFile(cmd.Param1)
		} else if "" != cmd.Param2 {
			log.Println("::" + cmd.Param2)
			c.PrintFile(cmd.Param2)
		}
	case "LOGIN":
		c.MatrixLogin()
	case "GDBYE":
		c.LogOff(cmd.Param1, cmd.Param2)
	}
}
