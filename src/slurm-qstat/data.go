package main

type nodeData map[string]string
type jobData map[string]string
type reservationData map[string]string
type clusterData map[string]string

type partitionInfo struct {
	CoresAllocated uint64
	CoresIdle      uint64
	CoresOther     uint64
	CoresTotal     uint64
	Name           string
}
