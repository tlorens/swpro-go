package main

// ConfigRec defines basic setup proeperties
type ConfigRec struct {
	BindIP     string `json:"bindIP"`
	BindPort   int16  `json:"bindPort"`
	SystemName string `json:"systemName"`
	SysOp      string `json:"sysOp"`
	StartMenu  string `json:"startMenu"`
}

type MenuRec struct {
	Name     string
	Prompt   string
	Flags    string
	Password string
	Commands []CommandRec
}

type CommandRec struct {
	Key      string
	Name     string
	Cmd      string
	Flags    string
	Param1   string
	Param2   string
	Priority int
}

type StringSetRec struct {
	Index   int
	Name    string
	Strings struct {
		GetHandle   string
		GetPassord  string
		NewHandle   string
		NewPassword string
	}
}

type StringSetsRec struct {
	Sets []StringSetRec
}
