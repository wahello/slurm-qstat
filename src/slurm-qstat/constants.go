package main

const name = "slurm-qstat"
const version = "1.4.2-20210926"

const versionText = `%s version %s
Copyright (C) 2021-2023 by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

%s is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Build with go version: %s
`

const helpText = `Usage: %s [--brief] [--filter=<part>,...] [--help] --clusters|--jobs=<filter>|--nodes|--partitions|--reservations [--sort=<sort>] [--version]

    --brief                     Show brief output

    --clusters                  Show cluster information

    --filter=<part>,...         Limit output to comma separated list of partitions

    --help                      Show this help text

    --jobs=<filter>             Show job information. <filter> can be one of:
                                    all         - show all jobs
                                    not-running - show not running only (state other than RUNNING)
                                    running     - show only running jobs (state RUNNING)

    --nodes                     Show node information

    --partitions                Show partition information

    --reservations              Show reservation information

    --sort=<sort>               Sort output by field <sort> in ascending order
                                    <sort> is a comma separated list of <object>:<field>
                                    <object> can be prefixed by a minus sign to reverse the sort order of the field
                                    <object> can be one of:
										clusters - sort clusters
                                        jobs - sort jobs
                                        nodes - sort nodes
                                        partitions - sort partitions
                                        reservations - sort reservations

                                    <field> depends of the <object> type:
										clusters:
											name - sort by cluster name
											controlhost - sort by control host
											controlport - sort by control port
											nodecount - sort by node count
											defaultqos - sort by default QoS
											fairshare - sort by fairshare
											maxjobs - sort my max. jobs
											maxnodes - sort by max. nodes
											maxsubmitjobs - sort by max. submitted jobs
											maxwall - sort by max. wall time
											tres - sort by TRES
											clusternodes - sort by cluster nodes

										jobs:
                                            batchhost - sort by batch host
                                            cpus - sort by cpus
                                            gres - sort by GRES
                                            jobid - sort by job id (this is the default)
                                            licenses - sort by licenses
                                            name - sort by name
                                            nodes - sort by nodes
                                            partition - sort by partitions
                                            reason - sort by state reason
                                            starttime - sort by starttime
                                            state - sort by state
                                            tres - sort by TRES
                                            user - sort by user

                                        nodes:
                                            boards - sort by number of boards
                                            hostname - sort by hostname
                                            nodename - sort by node name (this is the default)
                                            partition - sort by partitions
                                            reason - sort by state reason
                                            slurmversion - sort by reported SLURM version
                                            sockets - sort by number of sockets
                                            state - sort by state
                                            threadsbycore - sort by threads per core
                                            tresallocated - sort by allocated TRES
                                            tresconfigured - sort by configured TRES

                                        partitions:
                                            allocated - sort by allocated nodes
                                            allocatedpercent - sort by allocation percentage
                                            idle - sort by idle nodes
                                            idlepercent - sort by idle percentage
                                            other - sort by other nodes
                                            otherpercent - sort by percentage of other nodes
                                            partition - sort by partition name (this is the default)
                                            total - sort by total nodes

                                        reservations:
                                            accounts - sort by accounts
                                            burstbuffers - sort by burst buffers
                                            corecount - sort by core count
                                            duration - sort by duration
                                            end time - sort by end time
                                            features - sort by features
                                            flags - sort by flags
                                            licenses - sort by licenses
                                            name - sort by reservation name (this is the default)
                                            nodecount - sort by node count
                                            nodes - sort by nodes
                                            partition - sort by partition
                                            starttime - sort by start time
                                            state - sort by state
                                            tres - sort by TRES
                                            users - sort by users
                                            watts - sort by watts

    --version                   Show version information
`

const sortReverse uint8 = 0x80
const maskSortReverse uint8 = 0x7f

const sortReservationsMask uint64 = 0x00000000ff000000
const (
	sortReservationsByName uint8 = iota
	sortReservationsByPartition
	sortReservationsByState
	sortReservationsByStartTime
	sortReservationsByEndTime
	sortReservationsByDuration
	sortReservationsByNodes
	sortReservationsByNodeCount
	sortReservationsByCoreCount
	sortReservationsByFeatures
	sortReservationsByFlags
	sortReservationsByTres
	sortReservationsByUsers
	sortReservationsByAccounts
	sortReservationsByLicenses
	sortReservationsByBurstBuffer
	sortReservationsByWatts
)

const sortPartitionsMask uint64 = 0x0000000000ff0000
const (
	sortPartitionsByPartition uint8 = iota
	sortPartitionsByAllocated
	sortPartitionsByAllocatedPercent
	sortPartitionsByIdle
	sortPartitionsByIdlePercent
	sortPartitionsByOther
	sortPartitionsByOtherPercent
	sortPartitionsByTotal
)

const sortNodesMask uint64 = 0x00000000000000ff
const (
	sortNodesByNodeName uint8 = iota
	sortNodesByHostName
	sortNodesByPartition
	sortNodesByState
	sortNodesBySlurmVersion
	sortNodesByTresConfigured
	sortNodesByTresAllocated
	sortNodesBySockets
	sortNodesByBoards
	sortNodesByThreadsPerCore
	sortNodesByReason
)

const sortJobsMask uint64 = 0x000000000000ff00
const (
	sortJobsByJobID uint8 = iota
	sortJobsByPartition
	sortJobsByUser
	sortJobsByState
	sortJobsByReason
	sortJobsByBatchHost
	sortJobsByNodes
	sortJobsByCPUs
	sortJobsByLicenses
	sortJobsByGres
	sortJobsByTres
	sortJobsByName
	sortJobsByStartTime
)

const sortClusterMask uint64 = 0x000000ff00000000
const (
	sortClusterByName uint8 = iota
	sortClusterByControlHost
	sortClusterByControlPort
	sortClusterByNodeCount
	sortClusterByDefaultQos
	sortClusterByFairShare
	sortClusterByMaxJobs
	sortClusterByMaxNodes
	sortClusterByMaxSubmitJobs
	sortClusterByMaxWall
	sortClusterByTres
	sortClusterByClusterNodes
)

var clusterFields = []string{
	"Classification",
	"Cluster",
	"ClusterNodes",
	"ControlHost",
	"ControlPort",
	"DefaultQOS",
	"Fairshare",
	"Flags",
	"GrpTRESMins",
	"GrpTRES",
	"GrpJobs",
	"GrpMemory",
	"GrpNodes",
	"GrpSubmitJob",
	"MaxTRESMins",
	"MaxTRES",
	"MaxJobs",
	"MaxNodes",
	"MaxSubmitJobs",
	"MaxWall",
	"NodeCount",
	"PluginIDSelect",
	"RPC",
	"TRES",
}
