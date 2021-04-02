package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func printPartitionStatus(p map[string]partitionInfo) {
	var data [][]string
	var keys []string
	var idleSum uint64
	var allocatedSum uint64
	var otherSum uint64
	var totalSum uint64

	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value, found := p[key]
		// Will never happen
		if !found {
			log.Panicf("BUG: No entry found for partition %s\n", key)
		}

		idleSum += value.CoresIdle
		allocatedSum += value.CoresAllocated
		otherSum += value.CoresOther
		totalSum += value.CoresTotal

		data = append(data, []string{
			key,

			strconv.FormatUint(value.CoresIdle, 10),
			strconv.FormatUint(value.CoresAllocated, 10),
			strconv.FormatUint(value.CoresOther, 10),
			strconv.FormatUint(value.CoresTotal, 10),

			strconv.FormatFloat(float64(value.CoresIdle)/float64(value.CoresTotal)*100.0, 'f', 3, 64),
			strconv.FormatFloat(float64(value.CoresAllocated)/float64(value.CoresTotal)*100.0, 'f', 3, 64),
			strconv.FormatFloat(float64(value.CoresOther)/float64(value.CoresTotal)*100.0, 'f', 3, 64),
			strconv.FormatFloat(100.0, 'f', 3, 64),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Partition", "Idle", "Allocated", "Other", "Total", "Idle%", "Allocated%", "Other%", "Total%"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetFooter([]string{
		"Sum",
		strconv.FormatUint(idleSum, 10),
		strconv.FormatUint(allocatedSum, 10),
		strconv.FormatUint(otherSum, 10),
		strconv.FormatUint(totalSum, 10),

		strconv.FormatFloat(float64(idleSum)/float64(totalSum)*100.0, 'f', 3, 64),
		strconv.FormatFloat(float64(allocatedSum)/float64(totalSum)*100.0, 'f', 3, 64),
		strconv.FormatFloat(float64(otherSum)/float64(totalSum)*100.0, 'f', 3, 64),
		strconv.FormatFloat(100.0, 'f', 3, 64),
	})
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.AppendBulk(data)
	table.Render()

}

func printJobStatus(j map[string]jobData, jidList []string) {
	var reUser = regexp.MustCompile(`\(\d+\)`)
	var data [][]string
	var runCount uint64
	var pendCount uint64
	var otherCount uint64
	var totalCount uint64
	var failCount uint64
	var preeemptCount uint64
	var stopCount uint64
	var suspendCount uint64

	for _, job := range jidList {
		var host string
		var startTime string
		var pendingReason string

		jData, found := j[job]
		if !found {
			log.Panicf("BUG: No job data found for job %s\n", job)
		}

		user, found := jData["UserId"]
		if !found {
			log.Panicf("BUG: No user found for job %s\n", job)
		}

		user = reUser.ReplaceAllString(user, "")

		state, found := jData["JobState"]
		if !found {
			log.Panicf("BUG: No JobState found for job %s\n", job)
		}

		switch state {
		case "FAILED":
			failCount++
		case "PENDING":
			pendCount++
		case "PREEMPTED":
			preeemptCount++
		case "STOPPED":
			stopCount++
		case "SUSPENDED":
			suspendCount++
		case "RUNNING":
			runCount++
		default:
			otherCount++
		}
		totalCount++

		partition, found := jData["Partition"]
		if !found {
			log.Panicf("BUG: No partition found for job %s\n", job)
		}

		tres := jData["TRES"]

		_numCpus, found := jData["NumCPUs"]
		if !found {
			log.Panicf("BUG: NumCPUs not found for job %s\n", job)
		}
		numCpus, err := strconv.ParseUint(_numCpus, 10, 64)
		if err != nil {
			log.Panicf("BUG: Can't convert NumCpus to an integer for job %s: %s\n", job, err)
		}

		name, found := jData["JobName"]
		if !found {
			log.Panicf("BUG: JobName not set for job %s\n", job)
		}

		if state == "PENDING" {
			// Jobs can also be submitted, requesting a number of Nodes instead of CPUs
			// Therefore we will check TRES first
			tresCpus, err := getCpusFromTresString(tres)
			if err != nil {
				log.Panicf("BUG: Can't get number of CPUs from TRES as integer for job %s: %s\n", job, err)
			}

			if tresCpus != 0 {
				numCpus = tresCpus
			}

			// PENDING jobs never scheduled at all don't have BatchHost set (for obvious reasons)
			// Rescheduled and now PENDING jobs do have a BatchHost
			host, found = jData["BatchHost"]
			if !found {
				host = "<not_scheduled_yet>"
			}

			// The same applies for StartTime
			startTime, found = jData["StartTime"]
			if !found {
				startTime = "<not_scheduled_yet>"
			}

			// Obviously, PENDING jobs _always_ have a Reason
			pendingReason, found = jData["Reason"]
			if !found {
				log.Panicf("BUG: No Reason for pending job %s\n", job)
			}

		} else {
			host, found = jData["BatchHost"]
			if !found {
				log.Panicf("BUG: No BatchHost set for job %s\n", job)
			}

			startTime, found = jData["StartTime"]
			if !found {
				log.Panicf("BUG: No StartTime set for job %s\n", job)
			}
		}

		data = append(data, []string{
			job,
			partition,
			user,
			state,
			pendingReason,
			host,
			strconv.FormatUint(numCpus, 10),
			startTime,
			name,
		})

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"JobID", "Partition", "User", "State", "Reason", "Host", "CPUs", "Starttime", "Name"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetFooter([]string{
		"Sum",
		fmt.Sprintf("Failed: %d", failCount),
		fmt.Sprintf("Pending: %d", pendCount),
		fmt.Sprintf("Preempted: %d", preeemptCount),
		fmt.Sprintf("Stoped: %d", stopCount),
		fmt.Sprintf("Suspended: %d", suspendCount),
		fmt.Sprintf("Running: %d", runCount),
		fmt.Sprintf("Other: %d", otherCount),
		fmt.Sprintf("Total: %d", totalCount),
	})
	table.SetFooterAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()
}
