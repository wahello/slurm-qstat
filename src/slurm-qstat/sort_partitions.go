package main

import (
	"sort"
)

func sortPartitions(unsorted map[string]partitionInfo, sortBy uint8) []partitionInfo {
	var sorted []partitionInfo

	for _, value := range unsorted {
		sorted = append(sorted, value)
	}

	// XXX: Correct values are checkecd and defaults are set in previous calls to argument parser
	switch sortBy & maskSortReverse {

	case sortPartitionsByPartition:
		if sortBy&sortReverse == sortReverse {
			sorted = partitionSortedByPartitionName(sorted, true)
		} else {
			sorted = partitionSortedByPartitionName(sorted, false)
		}

	case sortPartitionsByAllocated:
		if sortBy&sortReverse == sortReverse {
			sorted = partitionSortedByAllocated(sorted, true)
		} else {
			sorted = partitionSortedByAllocated(sorted, false)
		}

	case sortPartitionsByAllocatedPercent:
		if sortBy&sortReverse == sortReverse {
			sorted = partitionSortedByAllocatedPercent(sorted, true)
		} else {
			sorted = partitionSortedByAllocatedPercent(sorted, false)
		}

	case sortPartitionsByIdle:
		if sortBy&sortReverse == sortReverse {
			sorted = partitionSortedByIdle(sorted, true)
		} else {
			sorted = partitionSortedByIdle(sorted, false)
		}

	case sortPartitionsByIdlePercent:
		if sortBy&sortReverse == sortReverse {
			sorted = partitionSortedByIdlePercent(sorted, true)
		} else {
			sorted = partitionSortedByIdlePercent(sorted, false)
		}

	case sortPartitionsByOther:
		if sortBy&sortReverse == sortReverse {
			sorted = partitionSortedByOther(sorted, true)
		} else {
			sorted = partitionSortedByOther(sorted, false)
		}

	case sortPartitionsByOtherPercent:
		if sortBy&sortReverse == sortReverse {
			sorted = partitionSortedByOtherPercent(sorted, true)
		} else {
			sorted = partitionSortedByOtherPercent(sorted, false)
		}

	case sortPartitionsByTotal:
		if sortBy&sortReverse == sortReverse {
			sorted = partitionSortedByTotal(sorted, true)
		} else {
			sorted = partitionSortedByTotal(sorted, false)
		}
	}
	return sorted
}

func partitionSortedByPartitionName(s []partitionInfo, reverse bool) []partitionInfo {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].Name > s[j].Name
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].Name < s[j].Name
		})
	}

	return s
}

func partitionSortedByAllocated(s []partitionInfo, reverse bool) []partitionInfo {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].CoresAllocated > s[j].CoresAllocated
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].CoresAllocated < s[j].CoresAllocated
		})
	}

	return s
}

func partitionSortedByAllocatedPercent(s []partitionInfo, reverse bool) []partitionInfo {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			var a float64
			var b float64

			if s[i].CoresTotal != 0 {
				a = float64(s[i].CoresAllocated) / float64(s[i].CoresTotal)
			}

			if s[j].CoresTotal != 0 {
				b = float64(s[j].CoresAllocated) / float64(s[j].CoresTotal)
			}

			return a > b
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			var a float64
			var b float64

			if s[i].CoresTotal != 0 {
				a = float64(s[i].CoresAllocated) / float64(s[i].CoresTotal)
			}

			if s[j].CoresTotal != 0 {
				b = float64(s[j].CoresAllocated) / float64(s[j].CoresTotal)
			}

			return a < b
		})
	}

	return s
}

func partitionSortedByIdle(s []partitionInfo, reverse bool) []partitionInfo {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].CoresIdle > s[j].CoresIdle
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].CoresIdle < s[j].CoresIdle
		})
	}

	return s
}

func partitionSortedByIdlePercent(s []partitionInfo, reverse bool) []partitionInfo {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			var a float64
			var b float64

			if s[i].CoresTotal != 0 {
				a = float64(s[i].CoresIdle) / float64(s[i].CoresTotal)
			}

			if s[j].CoresTotal != 0 {
				b = float64(s[j].CoresIdle) / float64(s[j].CoresTotal)
			}

			return a > b
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			var a float64
			var b float64

			if s[i].CoresTotal != 0 {
				a = float64(s[i].CoresIdle) / float64(s[i].CoresTotal)
			}

			if s[j].CoresTotal != 0 {
				b = float64(s[j].CoresIdle) / float64(s[j].CoresTotal)
			}

			return a < b
		})
	}

	return s
}

func partitionSortedByOther(s []partitionInfo, reverse bool) []partitionInfo {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].CoresOther > s[j].CoresOther
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].CoresOther < s[j].CoresOther
		})
	}

	return s
}

func partitionSortedByOtherPercent(s []partitionInfo, reverse bool) []partitionInfo {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			var a float64
			var b float64

			if s[i].CoresTotal != 0 {
				a = float64(s[i].CoresOther) / float64(s[i].CoresTotal)
			}

			if s[j].CoresTotal != 0 {
				b = float64(s[j].CoresOther) / float64(s[j].CoresTotal)
			}

			return a > b
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			var a float64
			var b float64

			if s[i].CoresTotal != 0 {
				a = float64(s[i].CoresOther) / float64(s[i].CoresTotal)
			}

			if s[j].CoresTotal != 0 {
				b = float64(s[j].CoresOther) / float64(s[j].CoresTotal)
			}

			return a < b
		})
	}

	return s
}

func partitionSortedByTotal(s []partitionInfo, reverse bool) []partitionInfo {
	if reverse {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].CoresTotal > s[j].CoresTotal
		})
	} else {
		sort.SliceStable(s, func(i int, j int) bool {
			return s[i].CoresTotal < s[j].CoresTotal
		})
	}

	return s
}
