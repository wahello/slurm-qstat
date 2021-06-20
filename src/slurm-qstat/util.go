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

func buildSortFlag(s string) (uint32, error) {
	var fl uint32
	var n uint8
	var j uint8
	var p uint8
	var r uint8
	var err error

	if len(s) == 0 {
		return fl, nil
	}

	for _, splitted := range strings.Split(s, ",") {
		whatby := strings.SplitN(splitted, ":", 2)
		if len(whatby) != 2 {
			return fl, fmt.Errorf("Invalid sorting string %s", s)
		}
		what := strings.ToLower(whatby[0])
		by := strings.ToLower(whatby[1])
		_rev := false
		if what[0] == '-' {
			_rev = true
			what = strings.Replace(what, "-", "", 1)
		}

		switch what {
		case "nodes":
			n, err = buildSortFlagNodes(by)
			if err != nil {
				return fl, err
			}
			if _rev {
				n |= sortReverse
			}

		case "jobs":
			j, err = buildSortFlagJobs(by)
			if err != nil {
				return fl, err
			}
			if _rev {
				j |= sortReverse
			}

		case "partitions":
			p, err = buildSortFlagPartitions(by)
			if err != nil {
				return fl, err
			}
			if _rev {
				p |= sortReverse
			}

		case "reservations":

		default:
			return fl, fmt.Errorf("Invalid sorting object to sort %s", s)
		}
	}

	fl = uint32(r)<<24 + uint32(p)<<16 + uint32(j)<<8 + uint32(n)
	return fl, nil
}

func buildSortFlagPartitions(s string) (uint8, error) {
	var n uint8
	switch s {
	case "partition":
		n = sortPartitionsByPartition
	case "allocated":
		n = sortPartitionsByAllocated
	case "allocatedpercent":
		n = sortPartitionsByAllocatedPercent
	case "idle":
		n = sortPartitionsByIdle
	case "idlepercent":
		n = sortPartitionsByIdlePercent
	case "other":
		n = sortPartitionsByOther
	case "otherpercent":
		n = sortPartitionsByOtherPercent
	case "total":
		n = sortPartitionsByTotal
	default:
		return n, fmt.Errorf("Invalid sort field %s for partitions", s)
	}
	return n, nil
}

func buildSortFlagJobs(s string) (uint8, error) {
	var n uint8

	switch s {
	case "batchhost":
		n = sortJobsByBatchHost
	case "cpus":
		n = sortJobsByCPUs
	case "gres":
		n = sortJobsByGres
	case "jobid":
		n = sortJobsByJobID
	case "licenses":
		n = sortJobsByLicenses
	case "name":
		n = sortJobsByName
	case "nodes":
		n = sortJobsByNodes
	case "partition":
		n = sortJobsByPartition
	case "reason":
		n = sortJobsByReason
	case "starttime":
		n = sortJobsByStartTime
	case "state":
		n = sortJobsByState
	case "tres":
		n = sortJobsByTres
	case "user":
		n = sortJobsByUser
	default:
		return n, fmt.Errorf("Invalid sort field %s for jobs", s)
	}
	return n, nil
}

func buildSortFlagNodes(s string) (uint8, error) {
	var n uint8

	switch s {
	case "nodename":
		n = sortNodesByNodeName
	case "hostname":
		n = sortNodesByHostName
	case "partition":
		n = sortNodesByPartition
	case "state":
		n = sortNodesByState
	case "slurmversion":
		n = sortNodesBySlurmVersion
	case "tresconfigured":
		n = sortNodesByTresConfigured
	case "tresallocated":
		n = sortNodesByTresAllocated
	case "sockets":
		n = sortNodesBySockets
	case "boards":
		n = sortNodesByBoards
	case "threadsbycore":
		n = sortNodesByThreadsPerCore
	case "reason":
		n = sortNodesByReason
	default:
		return n, fmt.Errorf("Invalid sort field %s for nodes", s)
	}
	return n, nil
}
