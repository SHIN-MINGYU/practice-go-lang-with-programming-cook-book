package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0.0"
const usage = `Usage:
%s [command]
Commands : 
	Greet
	Version
`

const greetUsage = `Usage:
%s greet name [flag]
Positional Arguments:
	  name - the name to greet
Flags :
`

// MenuConf struct save all level about nested command line
type MenuConf struct {
	Goodbye bool
}

// SetupMenu function set basic flag
func (m *MenuConf) SetupMenu() *flag.FlagSet{
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	menu.Usage = func(){
		fmt.Println(usage,os.Args[0])
		menu.PrintDefaults()
	}
	return menu
}

// GetSubMenu function return collection of flag about inferiority menu
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	submenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	submenu.BoolVar(&m.Goodbye,"goodbye",false,"Say goodbye instead of hello")
	submenu.Usage = func(){
		fmt.Println(greetUsage,os.Args[0])
        submenu.PrintDefaults()
	}
	return submenu
}

// Greet function is called by greet command
func (m *MenuConf) Greet(name string) {
	if(m.Goodbye) {
		fmt.Println("Goodbye " + name + "!")
	}else{
		fmt.Println("Hello " + name + "!")
	}
}

// Version function is to print saved const of version
func (m *MenuConf) Version() {
	fmt.Println("Version : " + version)
}
