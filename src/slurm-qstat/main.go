package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var help = flag.Bool("help", false, "Show help text")
	var version = flag.Bool("version", false, "Show version information")
	var filterStr = flag.String("filter", "", "Limit output to filter list")
	var partitions = flag.Bool("partitions", false, "Show partition information")
	var nodes = flag.Bool("nodes", false, "Show node information")
	var jobs = flag.String("jobs", "", "Show jobs")
	var reservations = flag.Bool("reservations", false, "Show reservations")
	var brief = flag.Bool("brief", false, "Show brief output")
	var sortby = flag.String("sort", "", "Sort output by fields")
	var filter []string

	flag.Usage = showHelp
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Fprintf(os.Stderr, "Error: Trailing arguments %s\n\n", strings.Join(flag.Args(), " "))
		showHelp()
		os.Exit(1)
	}

	if *help {
		showHelp()
		os.Exit(0)
	}

	if *version {
		showVersion()
		os.Exit(0)
	}

	if *filterStr != "" {
		filter = strings.Split(*filterStr, ",")
	}

	if !*partitions && *jobs == "" && !*nodes && !*reservations {
		fmt.Fprint(os.Stderr, "Error: What should be displayed?\n\n")
		showHelp()
		os.Exit(1)
	}

	if *jobs != "" {
		if *jobs != "running" && *jobs != "not-running" && *jobs != "all" {
			fmt.Fprint(os.Stderr, "Error: Invalid job display filter\n\n")
			showHelp()
			os.Exit(1)
		}
	}

	sortFlag, err := buildSortFlag(*sortby)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Can't parse sort string: %s", err)
		os.Exit(1)
	}

	if *partitions {
		nodeInfo, err := getNodeInformation()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get node data from SLURM: %s\n", err)
			os.Exit(1)
		}

		partInfo, err := getPartitionInformation(nodeInfo, filter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get partition information from SLURM: %s\n", err)
			os.Exit(1)
		}

		printPartitionStatus(partInfo, *brief)
	}

	if *jobs != "" {
		_jobInfo, err := getJobInformation()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get job data from SLURM: %s\n", err)
			os.Exit(1)
		}

		_jobInfo = filterJobs(_jobInfo, filter)
		_jobInfo = massageJobs(_jobInfo)

		var displayJobs []string

		if *jobs == "running" {
			_, displayJobs = splitByPendState(_jobInfo)
		} else if *jobs == "not-running" {
			displayJobs, _ = splitByPendState(_jobInfo)
		} else {
			for k := range _jobInfo {
				displayJobs = append(displayJobs, k)
			}
		}

		// only keep job data for jobs to be displayed
		var jobInfo = make(map[string]jobData)
		for _, j := range displayJobs {
			jobInfo[j] = _jobInfo[j]
		}

		jobInfoSorted := sortJobs(jobInfo, uint8(sortFlag&sortJobsMask))
		printJobStatus(jobInfoSorted, *brief)
	}

	if *nodes {
		nodeInfo, err := getNodeInformation()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get node information from SLURM: %s\n", err)
			os.Exit(1)
		}

		nodeInfoSorted := sortNodes(filterNodes(nodeInfo, filter), uint8(sortFlag&sortNodesMask))
		printNodeStatus(nodeInfoSorted, *brief)
	}

	if *reservations {
		rsvInfo, err := getReservations()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can'T get reservations from SLURM: %s\n", err)
			os.Exit(1)
		}

		rsvInfo = filterReservations(rsvInfo, filter)
		printReservationStatus(rsvInfo, *brief)
	}

}
