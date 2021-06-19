package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortNodes(unsorted map[string]nodeData, sortBy uint8) []nodeData {
	var sorted []nodeData

	for _, value := range unsorted {
		sorted = append(sorted, value)
	}

	// XXX: Correct values are checkecd and defaults are set in previous calls to argument parser
	switch sortBy & maskSortReverse {

	case sortNodesByNodeName: // this is the default
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByString(sorted, "NodeName", true)
		} else {
			sorted = nodeSortedByString(sorted, "NodeName", false)
		}

	case sortNodesByHostName:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByString(sorted, "NodeHostName", true)
		} else {
			sorted = nodeSortedByString(sorted, "NodeHostName", false)
		}

	case sortNodesByPartition:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByString(sorted, "Partitions", true)
		} else {
			sorted = nodeSortedByString(sorted, "Partitions", false)
		}

	case sortNodesByState:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByString(sorted, "State", true)
		} else {
			sorted = nodeSortedByString(sorted, "State", false)
		}

	case sortNodesBySlurmVersion:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByString(sorted, "Version", true)
		} else {
			sorted = nodeSortedByString(sorted, "Version", false)
		}

	case sortNodesByTresConfigured:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByString(sorted, "CfgTRES", true)
		} else {
			sorted = nodeSortedByString(sorted, "CfgTRES", false)
		}

	case sortNodesByTresAllocated:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByString(sorted, "AllocTRES", true)
		} else {
			sorted = nodeSortedByString(sorted, "AllocTRES", false)
		}

	case sortNodesBySockets:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByInt(sorted, "Sockets", true)
		} else {
			sorted = nodeSortedByInt(sorted, "Sockets", false)
		}

	case sortNodesByBoards:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByInt(sorted, "Boards", true)
		} else {
			sorted = nodeSortedByInt(sorted, "Boards", false)
		}

	case sortNodesByThreadsPerCore:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByInt(sorted, "ThreadsPerCore", true)
		} else {
			sorted = nodeSortedByInt(sorted, "ThreadsPerCore", false)
		}

	case sortNodesByReason:
		if sortBy&sortReverse == sortReverse {
			sorted = nodeSortedByString(sorted, "Reason", true)
		} else {
			sorted = nodeSortedByString(sorted, "Reason", false)
		}
	}
	return sorted
}

func nodeSortedByString(s []nodeData, field string, reverse bool) []nodeData {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			a := s[i][field]
			b := s[j][field]
			return a > b
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			a := s[i][field]
			b := s[j][field]
			return a < b
		})
	}

	return s
}

func nodeSortedByInt(s []nodeData, field string, reverse bool) []nodeData {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			_a := s[i][field]
			a, err := strconv.ParseInt(_a, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Can't convert data from key %s to a number: %s", field, err))
			}

			_b := s[j][field]
			b, err := strconv.ParseInt(_b, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Can't convert data from key %s to a number: %s", field, err))
			}
			return a > b
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			_a := s[i][field]
			a, err := strconv.ParseInt(_a, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Can't convert data from key %s to a number: %s", field, err))
			}

			_b := s[j][field]
			b, err := strconv.ParseInt(_b, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Can't convert data from key %s to a number: %s", field, err))
			}
			return a < b
		})
	}

	return s
}
