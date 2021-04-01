package main

type nodeData map[string]string

type partitionInfo struct {
	CoresAllocated uint64
	CoresIdle      uint64
	CoresOther     uint64
	CoresTotal     uint64
	Name           string
}
