package main

const name = "slurm-qstat"
const version = "1.0.1-20210402"

const versionText = `%s version %s
Copyright (C) 2021 by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

pkidb is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Build with go version: %s
`

const helpText = `Usage: %s [--filter=<part>,...] [--help] --jobs=<filter>|--nodes|--partitions [--version]

    --filter=<part>,...         Limit output to comma separated list of partitions

    --help                      Show this help text

    --jobs=<filter>             Show job information. <filter can be one of:
                                    all         - show all jobs
                                    not-running - show not running only (state other than RUNNING)
                                    running     - show only running jobs (state RUNNING)

    --nodes                     Show node information

    --partitions                Show partition information

    --version                   Show version information
`
