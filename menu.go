package main

import (
    "os"
    "strings"
    "log"
    "encoding/json"
)

var ranAutos = false
var menu MenuRec

func LoadMenu(file string) MenuRec {
    menuFile, err := os.Open("menudata/" + file)
    defer menuFile.Close()

    log.Println("Loading Menu: " + file)
    if err != nil {
        log.Println(err.Error())
    }

    var loadedMenu MenuRec
    jsonParser := json.NewDecoder(menuFile)
    jsonParser.Decode(&loadedMenu)

    return loadedMenu
}

/**
 * Execute commands tagged with AUTO to run ONCE automatically before the menut prompt
 */
func RunAutos(menuCommands []CommandRec) {
    for _, cmd := range menuCommands {
        if (Match(cmd.Param1, "AUTO") || Match(cmd.Param2, "AUTO")) {
            ranAutos = true
            log.Println("AUTO: " + cmd.Cmd)
            RunCommand(cmd)
        }
    }
}

/**
 * Execute commands tagged with 'EVERY' at every prompt. ]
 */
func RunEvery(menuCommands []CommandRec) {
    for _, cmd := range menuCommands {
        if (Match(cmd.Param1, "EVERY") || Match(cmd.Param2, "EVERY")) {
            log.Println("EVERY: " + cmd.Cmd)
            RunCommand(cmd)
        }
    }
}


/**
 * Read a JSON menu and find it's commands.
 * @param file string menu file name.
 */
func RunMenu(curMenu string) {
    var tmpMenuFile = ""
    // var menu MenuRec
    // Main loop. Repeat running menus forever.
    for {
        log.Println("Got Menu: " + curMenu)

        for tmpMenuFile != curMenu {
            tmpMenuFile = curMenu
            menu = LoadMenu(curMenu)
            // Execute 'AUTO' commands.
            if (false == ranAutos) {
                RunAutos(menu.Commands)
            }
        }

        // Execute 'EVERY' commands.
        RunEvery(menu.Commands)

        ch := Prompt(80, menu.Prompt)

        for _, cmd := range menu.Commands {
            if (Match(cmd.Key, ch)) {
                RunCommand(cmd)
            }
        }
    }
}

/**
 * Execute a nenus command.
 * @param cmd CommandRec Command to execute.
 */
func RunCommand(cmd CommandRec) {
    log.Println("RunCmd: " + cmd.Cmd)
	switch strings.ToUpper(cmd.Cmd) {
        case "CLEAR":
            NetWrite(ClearScr())
        case "@MENU":
            ranAutos = false
            if (Match(cmd.Param1, "AUTO") || Match(cmd.Param1, "EVERY")) {
                menu = LoadMenu(cmd.Param2)
            } else {
                menu = LoadMenu(cmd.Param1)
            }
            log.Println("New Menu: " + menu.Name)
		case "PRINT":
            if ("" != cmd.Param1) {
                PrintFile(cmd.Param1)
            } else if ("" != cmd.Param2) {
                PrintFile(cmd.Param2)
            }
        case "LOGIN":
            MatrixLogin()
        case "GDBYE":
            LogOff(cmd.Param1, cmd.Param2)
		default:
	}
}
