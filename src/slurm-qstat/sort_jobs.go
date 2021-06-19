package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortJobs(unsorted map[string]jobData, sortBy uint8) []jobData {
	var sorted []jobData

	for _, value := range unsorted {
		sorted = append(sorted, value)
	}

	// XXX: Correct values are checkecd and defaults are set in previous calls to argument parser
	switch sortBy & maskSortReverse {

	case sortJobsByJobID: // this is the default
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByInt(sorted, "JobId", true)
		} else {
			sorted = jobSortedByInt(sorted, "JobId", false)
		}

	case sortJobsByPartition:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "Partition", true)
		} else {
			sorted = jobSortedByString(sorted, "Partition", false)
		}

	case sortJobsByUser:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "UserId", true)
		} else {
			sorted = jobSortedByString(sorted, "UserId", false)
		}

	case sortJobsByState:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "State", true)
		} else {
			sorted = jobSortedByString(sorted, "State", false)
		}

	case sortJobsByReason:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "Reason", true)
		} else {
			sorted = jobSortedByString(sorted, "Reason", false)
		}

	case sortJobsByBatchHost:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "BatchHost", true)
		} else {
			sorted = jobSortedByString(sorted, "BatchHost", false)
		}

	case sortJobsByNodes:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "NodeList", true)
		} else {
			sorted = jobSortedByString(sorted, "NodeList", false)
		}

	case sortJobsByCPUs:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByInt(sorted, "NumCPUs", true)
		} else {
			sorted = jobSortedByInt(sorted, "NumCPUs", false)
		}

	case sortJobsByLicenses:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "Licenses", true)
		} else {
			sorted = jobSortedByString(sorted, "Licenses", false)
		}

	case sortJobsByGres:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "Gres", true)
		} else {
			sorted = jobSortedByString(sorted, "Gres", false)
		}

	case sortJobsByTres:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "TRES", true)
		} else {
			sorted = jobSortedByString(sorted, "TRES", false)
		}

	case sortJobsByName:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "JobName", true)
		} else {
			sorted = jobSortedByString(sorted, "JobName", false)
		}

	case sortJobsByStartTime:
		if sortBy&sortReverse == sortReverse {
			sorted = jobSortedByString(sorted, "StartTime", true)
		} else {
			sorted = jobSortedByString(sorted, "StartTime", false)
		}
	}
	return sorted
}

func jobSortedByString(s []jobData, field string, reverse bool) []jobData {
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

func jobSortedByInt(s []jobData, field string, reverse bool) []jobData {
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
