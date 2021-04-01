package main

import (
	"fmt"
	"runtime"
)

func showVersion() {
	fmt.Printf(versionText, name, version, runtime.Version())
}

func showHelp() {
	showVersion()
	fmt.Print(helpText)
}
