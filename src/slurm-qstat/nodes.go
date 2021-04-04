package main

import (
	"log"
	"strings"
)

func getNodeInformation() (map[string]nodeData, error) {
	var result = make(map[string]nodeData)

	raw, err := executeCommand("scontrol", "show", "--oneliner", "--quiet", "nodes")
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

		node, found := data["NodeHostName"]
		// Should never happen!
		if !found {
			panic("NodeHostName not found")
		}

		result[node] = data
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

func splitNodePartitionString(p string) []string {
	var result []string

	for _, s := range strings.Split(p, ",") {
		s := strings.TrimSpace(s)
		result = append(result, s)
	}

	return result
}

func filterNodes(nodes map[string]nodeData, filter []string) map[string]nodeData {
	var result = make(map[string]nodeData)

	if len(filter) == 0 {
		return nodes
	}

	for key, value := range nodes {
		partitions, found := value["Partitions"]
		if !found {
			log.Panicf("BUG: No Partitions found for node %s\n", key)
		}

		partList := splitNodePartitionString(partitions)
		for _, part := range partList {
			if isPartitionInList(part, filter) {
				result[key] = value
			}
		}
	}

	return result
}
