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
	var clusters = flag.Bool("clusters", false, "Show clusters")
	var filter []string
	var cluster = flag.String("cluster", "", "Show data for cluster <cluster>")
	var _clusterInfo map[string]clusterData
	var _cluster []string

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

	if !*partitions && *jobs == "" && !*nodes && !*reservations && !*clusters {
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
		fmt.Fprintf(os.Stderr, "Error: Can't parse sort string: %s\n\n", err)
		showHelp()
		os.Exit(1)
	}

	if *cluster != "" {
		_clusterInfo, err = getClusterInformation()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get cluster list from SLURM: %s\n", err)
			os.Exit(1)
		}
		if *cluster == "all" {
			for key := range _clusterInfo {
				_cluster = append(_cluster, key)
			}
		} else {
			_cluster = strings.Split(*cluster, ",")
			// Sanity check, all always include all clusters
			for _, v := range _cluster {
				if v == "all" && len(_cluster) > 1 {
					fmt.Fprintf(os.Stderr, "Error: Keyword 'all' found in cluster list. Don't specify any additional clusters because they will be already included")
					os.Exit(1)
				}
			}
		}

		err = checkClusterlist(_cluster, _clusterInfo)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if *clusters {
		clustInfo, err := getClusterInformation()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get cluster list from SLURM: %s\n", err)
			os.Exit(1)
		}
		//		_cInfo := filterCluster(clustInfo, _cluster)
		cInfo := sortClusters(clustInfo, uint8((sortFlag&sortClusterMask)>>32))
		printClusterStatus(cInfo, *brief)
	}

	if *partitions {
		nodeInfo, err := getNodeInformation()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get node data from SLURM: %s\n", err)
			os.Exit(1)
		}

		_partInfo, err := getPartitionInformation(nodeInfo, filter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get partition information from SLURM: %s\n", err)
			os.Exit(1)
		}

		partInfo := sortPartitions(_partInfo, uint8((sortFlag&sortPartitionsMask)>>16))
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

		jobInfoSorted := sortJobs(jobInfo, uint8((sortFlag&sortJobsMask)>>8))
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
		_rsvInfo, err := getReservations()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can'T get reservations from SLURM: %s\n", err)
			os.Exit(1)
		}

		rsvInfo := sortReservations(filterReservations(_rsvInfo, filter), uint8((sortFlag&sortReservationsMask)>>24))
		printReservationStatus(rsvInfo, *brief)
	}

}
