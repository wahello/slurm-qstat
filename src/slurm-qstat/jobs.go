package main

import (
	"log"
	"regexp"
	"strings"
)

func getJobInformation() (map[string]jobData, error) {
	var result = make(map[string]jobData)
	var regexpWhiteSpace = regexp.MustCompile("\\s+")

	raw, err := executeCommand("scontrol", "--details", "--oneliner", "--quiet", "show", "jobs")
	if err != nil {
		return nil, err
	}

	rawStr := string(raw)
	for _, line := range strings.Split(rawStr, "\n") {
		// XXX: Remove duplicate white space, because some SLURM versions, like 17.11.6, contain extra white space after CoreSpec=
		line = regexpWhiteSpace.ReplaceAllString(line, " ")
		line = strings.TrimSpace(line)

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		data := make(map[string]string)

		// Split whitespace separated list of key=value pairs
		kvlist := strings.Split(line, " ")
		for _, kv := range kvlist {
			// Separate key and value
			_kv := strings.SplitN(kv, "=", 2)

			if len(_kv) == 1 {
				continue
			}

			key := _kv[0]
			value := _kv[1]

			data[key] = string(value)
		}

		job, found := data["JobId"]
		// Should never happen!
		if !found {
			log.Panic("BUG: JobId not found\n")
		}

		result[job] = data
	}

	return result, nil
}

// Weed out COMPLETED jobs, return list of pending and non-pending jobs
func splitByPendState(jobs map[string]jobData) ([]string, []string) {
	var pending []string
	var nonPending []string

	for jobID, jdata := range jobs {
		jstate, found := jdata["JobState"]
		// Should never happen!
		if !found {
			log.Panicf("BUG: Job %s doesn't have JobState field", jobID)
		}

		if jstate == "COMPLETED" {
			continue
		}

		if jstate == "PENDING" {
			pending = append(pending, jobID)
		} else {
			nonPending = append(nonPending, jobID)
		}
	}

	return pending, nonPending
}

func filterJobs(jobs map[string]jobData, filter []string) map[string]jobData {
	var result = make(map[string]jobData)

	if len(filter) == 0 {
		return jobs
	}

	for key, value := range jobs {
		partition, found := value["Partition"]
		if !found {
			log.Panicf("BUG: No Partition found for job %s\n", key)
		}

		if !isPartitionInList(partition, filter) {
			continue
		}

		result[key] = value
	}

	return result
}
