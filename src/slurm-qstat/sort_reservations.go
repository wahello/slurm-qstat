package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortReservations(unsorted map[string]reservationData, sortBy uint8) []reservationData {
	var sorted []reservationData

	for _, value := range unsorted {
		sorted = append(sorted, value)
	}

	// XXX: Correct values are checkecd and defaults are set in previous calls to argument parser
	switch sortBy & maskSortReverse {

	case sortReservationsByName: // this is the default
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "ReservationName", true)
		} else {
			sorted = reservationSortedByString(sorted, "ReservationName", false)
		}

	case sortReservationsByPartition:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Partition", true)
		} else {
			sorted = reservationSortedByString(sorted, "Partition", false)
		}

	case sortReservationsByState:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "State", true)
		} else {
			sorted = reservationSortedByString(sorted, "State", false)
		}

	case sortReservationsByStartTime:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "StartTime", true)
		} else {
			sorted = reservationSortedByString(sorted, "StartTime", false)
		}

	case sortReservationsByEndTime:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "EndTime", true)
		} else {
			sorted = reservationSortedByString(sorted, "EndTime", false)
		}

	case sortReservationsByDuration:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByInt(sorted, "Duration", true)
		} else {
			sorted = reservationSortedByInt(sorted, "Duration", false)
		}

	case sortReservationsByNodes:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Nodes", true)
		} else {
			sorted = reservationSortedByString(sorted, "Nodes", false)
		}

	case sortReservationsByNodeCount:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByInt(sorted, "NodeCnt", true)
		} else {
			sorted = reservationSortedByInt(sorted, "NodeCnt", false)
		}

	case sortReservationsByCoreCount:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByInt(sorted, "CoreCnt", true)
		} else {
			sorted = reservationSortedByInt(sorted, "CoreCnt", false)
		}

	case sortReservationsByFeatures:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Features", true)
		} else {
			sorted = reservationSortedByString(sorted, "Features", false)
		}

	case sortReservationsByFlags:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Flags", true)
		} else {
			sorted = reservationSortedByString(sorted, "Flags", false)
		}

	case sortReservationsByTres:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Tres", true)
		} else {
			sorted = reservationSortedByString(sorted, "Tres", false)
		}

	case sortReservationsByUsers:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Users", true)
		} else {
			sorted = reservationSortedByString(sorted, "Users", false)
		}

	case sortReservationsByAccounts:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Accounts", true)
		} else {
			sorted = reservationSortedByString(sorted, "Accounts", false)
		}

	case sortReservationsByLicenses:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Licenses", true)
		} else {
			sorted = reservationSortedByString(sorted, "Licenses", false)
		}

	case sortReservationsByBurstBuffer:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "BurstBuffer", true)
		} else {
			sorted = reservationSortedByString(sorted, "BurstBuffer", false)
		}

	case sortReservationsByWatts:
		if sortBy&sortReverse == sortReverse {
			sorted = reservationSortedByString(sorted, "Watts", true)
		} else {
			sorted = reservationSortedByString(sorted, "Watts", false)
		}
	}
	return sorted
}

func reservationSortedByString(s []reservationData, field string, reverse bool) []reservationData {
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

func reservationSortedByInt(s []reservationData, field string, reverse bool) []reservationData {
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
