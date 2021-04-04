package main

import (
	"log"
	"strings"
)

func getReservations() (map[string]reservationData, error) {
	var result = make(map[string]reservationData)

	raw, err := executeCommand("scontrol", "show", "--oneliner", "--quiet", "reservations")
	if err != nil {
		return nil, err
	}

	rawStr := string(raw)
	for _, line := range strings.Split(rawStr, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		data := make(map[string]string)

		// Split whitespace separated list of key=value pairs
		kvlist := strings.Split(line, " ")
		for i, kv := range kvlist {

			// Separate key and value
			_kv := strings.SplitN(kv, "=", 2)

			if len(_kv) == 1 {
				// FIXME: This is a crude workaround, because OS= contains white spaces too (e.g. OS=Linux 5.10.0-5-amd64)
				continue
			}

			key := _kv[0]
			value := _kv[1]
			// Reason is always the last part of the string and can contain white spaces!
			if key == "Reason" {
				value = strings.Replace(strings.Join(kvlist[i:len(kvlist)-1], " "), "Reason=", "", 1)
			}

			data[key] = string(value)
		}

		node, found := data["ReservationName"]
		// Should never happen!
		if !found {
			panic("ReservationName not found")
		}

		result[node] = data
	}

	return result, nil
}

func filterReservations(rsv map[string]reservationData, filter []string) map[string]reservationData {
	var result = make(map[string]reservationData)

	if len(filter) == 0 {
		return rsv
	}

	for _rsv, rsvData := range rsv {
		partition, found := rsvData["PartitionName"]
		if !found {
			log.Panicf("BUG: No PartitionName for reservation %s", _rsv)
		}

		if !isPartitionInList(partition, filter) {
			continue
		}

		result[_rsv] = rsvData
	}

	return result
}
