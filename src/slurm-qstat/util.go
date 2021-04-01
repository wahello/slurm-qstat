package main

import (
	"os"
	"os/exec"
)

func executeCommand(cmd string, args ...string) ([]byte, error) {
	exe := exec.Command(cmd, args...)
	exe.Env = append(os.Environ())
	exe.Env = append(exe.Env, "LANG=C")
	exe.Env = append(exe.Env, "SLURM_TIME_FORMAT=standard")

	return exe.Output()
}
