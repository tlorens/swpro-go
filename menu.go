package main

import (
    "os"
    "strings"
    "log"
    "encoding/json"
)

var ranAutos = false

func LoadMenu(file string) MenuRec {
    menuFile, err := os.Open("menudata/" + file)
    defer menuFile.Close()

    log.Println("LoadMenu: " + file)
    if err != nil {
        log.Println(err.Error())
    }

    var menu MenuRec
    jsonParser := json.NewDecoder(menuFile)
    jsonParser.Decode(&menu)

    return menu
}

/**
 * Execute commands tagged with AUTO to run ONCE automatically before the menut prompt
 */
func RunAutos(menuCommands []CommandRec) {
    for _, cmd := range menuCommands {
        if (cmd.Param1 == "AUTO" || cmd.Param2 == "AUTO") {
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
        if (cmd.Param1 == "EVERY" || cmd.Param2 == "EVERY") {
            log.Println("EVERY: " + cmd.Cmd)
            RunCommand(cmd)
        }
    }
}


/**
 * Read a JSON menu and find it's commands.
 * @param file string menu file name.
 */
func RunMenu() {
    var tmpMenuFile = ""
    var menu MenuRec
    // Main loop. Repeat running menus forever.
    for {
        log.Println("RunMenu: " + CurMenuFile)

        for tmpMenuFile != CurMenuFile {
            tmpMenuFile = CurMenuFile
            menu = LoadMenu(CurMenuFile)
            // Execute 'AUTO' commands.
            if (false == ranAutos) {
                RunAutos(menu.Commands)
            }
        }

        // Execute 'EVERY' commands.
        RunEvery(menu.Commands)

        ch := Prompt(80, menu.Prompt)

        for _, cmd := range menu.Commands {
            if (strings.ToUpper(cmd.Key) == strings.ToUpper(ch)) {
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
            Write(ClearScr())
        case "@MENU":
            PrevMenuFile = CurMenuFile
            ranAutos = false
            if ("AUTO" != strings.ToUpper(cmd.Param1) && "EVERY" != strings.ToUpper(cmd.Param1)) {
                CurMenuFile = cmd.Param1
            } else {
                CurMenuFile = cmd.Param2
            }
            log.Println("New Menu: " + CurMenuFile)
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
