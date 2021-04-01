package main

const name = "slurm-qstat"
const version = "1.0.0-20210401"

const versionText = `%s version %s
Copyright (C) 2021 by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

pkidb is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Build with go version: %s
`

const helpText = `Usage: %s [--filter=<part>,...] [--help] --jobs|--partitions [--version]

    --filter=<part>,...         Limit output to comma separated list of partitions

    --help                      Show this help text

    --jobs                      Show only job information

    --partitions                Show only partition information

    --version                   Show version information
`
