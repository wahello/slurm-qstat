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
	var filter []string

	flag.Usage = showHelp
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Fprintf(os.Stderr, "Error: Trailing arguments %s\n", strings.Join(flag.Args(), " "))
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

	if !*partitions && *jobs == "" && !*nodes {
		fmt.Fprint(os.Stderr, "Error: What should be displayed?\n")
		showHelp()
		os.Exit(1)
	}

	if *jobs != "" {
		if *jobs != "running" && *jobs != "not-running" && *jobs != "all" {
			fmt.Fprint(os.Stderr, "Error: Invalid job display filter\n")
			showHelp()
			os.Exit(1)
		}
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

		printPartitionStatus(partInfo)
	}

	if *jobs != "" {
		jobInfo, err := getJobInformation()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get job data from SLURM: %s\n", err)
			os.Exit(1)
		}

		jobInfo = filterJobs(jobInfo, filter)

		if *jobs == "running" {
			_, notPending := splitByPendState(jobInfo)

			notPending, err = sortByNumber(notPending)
			if err != nil {
				// Should never happen!
				panic(err)
			}

			printJobStatus(jobInfo, notPending)
		} else if *jobs == "not-running" {
			pending, _ := splitByPendState(jobInfo)

			pending, err = sortByNumber(pending)
			if err != nil {
				// Should never happen!
				panic(err)
			}

			printJobStatus(jobInfo, pending)
		} else {
			// show all jobs
			var all []string
			for key := range jobInfo {
				all = append(all, key)
			}

			allJobs, err := sortByNumber(all)
			if err != nil {
				// Should never happen!
				panic(err)
			}

			printJobStatus(jobInfo, allJobs)
		}
	}

	if *nodes {
		nodeInfo, err := getNodeInformation()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get node information from SLURM: %s\n", err)
			os.Exit(1)
		}

		nodeInfo = filterNodes(nodeInfo, filter)
		printNodeStatus(nodeInfo)
	}

}
