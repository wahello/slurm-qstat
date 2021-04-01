package main

import (
	"log"
	"os"
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
	/*
	   var idlePct float64
	   var allocatedPct float64
	   var otherPct float64
	   var totalPct float64
	*/

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
	table.AppendBulk(data)
	table.Render()

}
