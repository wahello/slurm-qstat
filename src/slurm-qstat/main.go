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
	var jobs = flag.Bool("jobs", false, "Show job information")
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

	if !*partitions && !*jobs {
		fmt.Fprintln(os.Stderr, "Error: What should be displayed?\n")
		showHelp()
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

		printPartitionStatus(partInfo)
	}
}
