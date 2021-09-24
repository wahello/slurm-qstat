package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortClusters(unsorted map[string]clusterData, sortBy uint8) []clusterData {
	var sorted []clusterData

	for _, value := range unsorted {
		sorted = append(sorted, value)
	}

	// XXX: Correct values are checkecd and defaults are set in previous calls to argument parser
	switch sortBy & maskSortReverse {
	case sortClusterByName: // this is the default
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByString(sorted, "Cluster", true)
		} else {
			sorted = clusterSortedByString(sorted, "Cluster", false)
		}

	case sortClusterByControlHost:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByString(sorted, "ControlHost", true)
		} else {
			sorted = clusterSortedByString(sorted, "ControlHost", false)
		}

	case sortClusterByControlPort:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByString(sorted, "ControlPort", true)
		} else {
			sorted = clusterSortedByString(sorted, "ControlPort", false)
		}

	case sortClusterByNodeCount:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByInt(sorted, "NodeCount", true)
		} else {
			sorted = clusterSortedByInt(sorted, "NodeCount", false)
		}

	case sortClusterByDefaultQos:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByString(sorted, "DefaultQoS", true)
		} else {
			sorted = clusterSortedByString(sorted, "DefaultQoS", false)
		}

	case sortClusterByFairShare:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByString(sorted, "FairShare", true)
		} else {
			sorted = clusterSortedByString(sorted, "FairShare", false)
		}

	case sortClusterByMaxJobs:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByInt(sorted, "MaxJobs", true)
		} else {
			sorted = clusterSortedByInt(sorted, "MaxJobs", false)
		}

	case sortClusterByMaxNodes:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByInt(sorted, "MaxNodes", true)
		} else {
			sorted = clusterSortedByInt(sorted, "MaxNodes", false)
		}

	case sortClusterByMaxSubmitJobs:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByInt(sorted, "MaxSubmitJobs", true)
		} else {
			sorted = clusterSortedByInt(sorted, "MaxSubmitJobs", false)
		}

	case sortClusterByMaxWall:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByString(sorted, "MaxWall", true)
		} else {
			sorted = clusterSortedByString(sorted, "MaxWall", false)
		}

	case sortClusterByTres:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByString(sorted, "TRES", true)
		} else {
			sorted = clusterSortedByString(sorted, "TRES", false)
		}

	case sortClusterByClusterNodes:
		if sortBy&sortReverse == sortReverse {
			sorted = clusterSortedByString(sorted, "ClusterNodes", true)
		} else {
			sorted = clusterSortedByString(sorted, "ClusterNodes", false)
		}

	}

	return sorted
}

func clusterSortedByString(c []clusterData, field string, reverse bool) []clusterData {
	if reverse {
		sort.SliceStable(c, func(i int, j int) bool {
			a := c[i][field]
			b := c[j][field]
			return a > b
		})
	} else {
		sort.SliceStable(c, func(i int, j int) bool {
			a := c[i][field]
			b := c[j][field]
			return a < b
		})
	}

	return c
}

func clusterSortedByInt(c []clusterData, field string, reverse bool) []clusterData {
	if reverse {
		sort.SliceStable(c, func(i int, j int) bool {
			_a := c[i][field]
			a, err := strconv.ParseInt(_a, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Can't convert data from key %s to a number: %s", field, err))
			}

			_b := c[j][field]
			b, err := strconv.ParseInt(_b, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Can't convert data from key %s to a number: %s", field, err))
			}
			return a > b
		})
	} else {
		sort.SliceStable(c, func(i int, j int) bool {
			_a := c[i][field]
			a, err := strconv.ParseInt(_a, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Can't convert data from key %s to a number: %s", field, err))
			}

			_b := c[j][field]
			b, err := strconv.ParseInt(_b, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Can't convert data from key %s to a number: %s", field, err))
			}
			return a < b
		})
	}

	return c
}
