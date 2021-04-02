package main

const name = "slurm-qstat"
const version = "1.0.0-20210402"

const versionText = `%s version %s
Copyright (C) 2021 by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

pkidb is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Build with go version: %s
`

const helpText = `Usage: %s [--filter=<part>,...] [--help] --jobs=<filter>|--partitions [--version]

    --filter=<part>,...         Limit output to comma separated list of partitions

    --help                      Show this help text

    --jobs=<filter>             Show job information. <filter can be one of:
                                    all         - show all jobs
                                    not-running - show not running only (state other than RUNNING)
                                    running     - show only running jobs (state RUNNING)

    --partitions                Show partition information

    --version                   Show version information
`

var compactJobState = map[string]string{
	"BOOT_FAIL":     "BF",
	"CANCELLED":     "CA",
	"COMPLETED":     "CD",
	"COMPLETING":    "CG",
	"CONFIGURING":   "CF",
	"DEADLINE":      "DL",
	"FAILED":        "F",
	"NODE_FAIL":     "NF",
	"OUT_OF_MEMORY": "OOM",
	"PENDING":       "PD",
	"PREEMPTED":     "PR",
	"REQUEUED":      "RQ",
	"REQUEUE_FED":   "RF",
	"REQUEUE_HOLD":  "RH",
	"RESIZING":      "RS",
	"RESV_DEL_HOLD": "RD",
	"REVOKED":       "RV",
	"RUNNING":       "R",
	"SPECIAL_EXIT":  "SE",
	"STOPPED":       "ST",
	"SUSPENDED":     "S",
}
