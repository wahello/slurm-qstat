package main

import (
	"fmt"
	"runtime"
)

func showVersion() {
	fmt.Printf(versionText, name, version, name, runtime.Version())
}

func showHelp() {
	showVersion()
	fmt.Printf(helpText, name)
}
