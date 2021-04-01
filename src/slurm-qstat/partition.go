package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func getNodeInformation() (map[string]nodeData, error) {
	var result = make(map[string]nodeData)

	raw, err := executeCommand("scontrol", "show", "--oneliner", "nodes")
	if err != nil {
		return nil, err
	}

	rawstr := string(raw)

	for _, line := range strings.Split(rawstr, "\n") {
		// Don't even start to process empty lines
		if len(line) == 0 {
			continue
		}

		data := make(map[string]string)

		// Split whitespace separated list of key=value pairs
		kvlist := strings.Split(line, " ")
		for i, kv := range kvlist {

			// Separate key and value
			_kv := strings.SplitN(kv, "=", 2)

			if len(_kv) == 1 {
				// FIXME: This is a crude workaround, because OS= contains white spaces too (e.g. OS=Linux 5.10.0-5-amd64)
				continue
			}

			key := _kv[0]
			value := _kv[1]
			// Reason is always the last part of the string and can contain white spaces!
			if key == "Reason" {
				value = strings.Replace(strings.Join(kvlist[i:len(kvlist)-1], " "), "Reason=", "", 1)
			}

			data[key] = string(value)
		}

		node, found := data["NodeName"]
		// Should never happen!
		if !found {
			panic("NodeName not found")
		}

		result[node] = data
	}

	return result, nil
}

func isPartitionInList(partition string, filter []string) bool {
	for _, flt := range filter {
		if partition == flt {
			return true
		}
	}

	return false
}

func getAllPartitions() ([]string, error) {
	var result []string

	raw, err := executeCommand("sinfo", "--all", "--noheader", "--Format=partitionname")
	if err != nil {
		return nil, err
	}

	rawStr := string(raw)
	for _, line := range strings.Split(rawStr, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		result = append(result, line)
	}

	return result, nil
}

func getPartitionInformation(nodeInfo map[string]nodeData, filter []string) (map[string]partitionInfo, error) {
	var result = make(map[string]partitionInfo)
	var regexpWhiteSpace = regexp.MustCompile("\\s+")
	var err error

	if len(filter) == 0 {
		filter, err = getAllPartitions()
		if err != nil {
			return nil, err
		}
	}

	// Note: partitionname returns the name of the partition, but without an asterisk for the default partition (like partition)
	raw, err := executeCommand("sinfo", "--all", "--noheader", "--Format=partitionname")
	if err != nil {
		return nil, err
	}

	rawstr := string(raw)
	for _, line := range strings.Split(rawstr, "\n") {
		if len(line) == 0 {
			continue
		}

		// Condense white space in output line to a single space
		line = regexpWhiteSpace.ReplaceAllString(line, " ")
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		partition := line

		// Discard information if partition is not in request filter
		if !isPartitionInList(partition, filter) {
			continue
		}

		// Get all nodes for requested partition
		partitionNodes, err := getNodesOfPartiton(partition)
		if err != nil {
			return nil, err
		}

		var pInfo partitionInfo
		for _, node := range partitionNodes {
			nInfo, found := nodeInfo[node]
			if !found {
				log.Panicf("BUG: Node %s in node list of partion %s, but no node information found from scontrol\n", node, partition)
			}

			cpuStr, found := nInfo["CPUTot"]
			if !found {
				log.Panicf("BUG: CPUTot not found for node %s\n", node)
			}
			cpus, err := strconv.ParseUint(cpuStr, 10, 64)
			if err != nil {
				return nil, err
			}

			cpuStr, found = nInfo["CPUAlloc"]
			if !found {
				log.Panicf("BUG: CPUAlloc not found for node %s\n", node)
			}
			used, err := strconv.ParseUint(cpuStr, 10, 64)
			if err != nil {
				return nil, err
			}

			state, found := nInfo["State"]
			if !found {
				log.Panicf("BUG: State not found for node %s\n", node)
			}

			pInfo.CoresTotal += cpus
			if state == "IDLE" {
				pInfo.CoresIdle += cpus
			} else if state == "ALLOCATED" || state == "MIXED" {
				pInfo.CoresIdle += cpus - used
				pInfo.CoresAllocated += used
			} else {
				pInfo.CoresOther += cpus
			}

			pInfo.Name = partition
		}

		result[partition] = pInfo
	}

	return result, nil
}

func getNodesOfPartiton(partition string) ([]string, error) {
	var result []string
	var raw []byte
	var err error

	if partition != "" {
		raw, err = executeCommand("sinfo", "--noheader", "--Format=nodehost", "--partition="+partition)
	} else {
		raw, err = executeCommand("sinfo", "--noheader", "--Format=nodehost")
	}

	if err != nil {
		return result, err
	}

	rawstr := string(raw)
	for _, line := range strings.Split(rawstr, " ") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		result = append(result, line)
	}

	return result, nil
}
