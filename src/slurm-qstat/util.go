package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func executeCommand(cmd string, args ...string) ([]byte, error) {
	exe := exec.Command(cmd, args...)
	exe.Env = append(os.Environ())
	exe.Env = append(exe.Env, "LANG=C")
	exe.Env = append(exe.Env, "SLURM_TIME_FORMAT=standard")

	return exe.Output()
}

func sortByNumber(u []string) ([]string, error) {
	var temp []int64
	var sorted []string

	for _, n := range u {
		i, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			return nil, err
		}

		temp = append(temp, i)
	}

	sort.SliceStable(temp, func(i int, j int) bool {
		return temp[i] < temp[j]
	})

	for _, i := range temp {
		sorted = append(sorted, strconv.FormatInt(i, 10))
	}

	return sorted, nil
}

func getCpusFromTresString(tres string) (uint64, error) {
	var cpu uint64
	var err error

	_tres := strings.Split(tres, ",")
	for _, tresKv := range _tres {
		kv := strings.Split(tresKv, "=")
		if len(kv) != 2 {
			return 0, fmt.Errorf("BUG: Not a key-value pair in TRES: %s", tresKv)
		}

		if kv[0] == "cpu" {
			cpu, err = strconv.ParseUint(kv[1], 10, 64)
			if err != nil {
				return 0, err
			}

		}
	}
	return cpu, nil
}
