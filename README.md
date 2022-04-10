# Preface
`slurm-qstat` tries to display information about jobs, nodes, partitions and reservations of the SLURM queueing system

# Repositories
* Primary development repository: https://git.ypbind.de/cgit/slurm-qstat/
* Backup repository: https://github.com/Bobobo-bo-Bo-bobo/slurm-qstat

# Build requirements
This tool is implemented in Go so, obviously, a Go compiler is required.
Additionally the [tablewriter](https://github.com/olekukonko/tablewriter) is required for the build.

# Command line parameters

| *Option* | *Description* | *Note* |
|:---------|:--------------|:-------|
| `--brief` | Show brief output | - |
| `--filter=<part>,...` | Limit output of to the comma separated list of partitions | If not specified, information about jobs, nodes and partitions of all partitions will be displayed |
| `--help` | Show the help text | |
| `--jobs=<filter>` | List information about jobs | Filter is mandatory and can be one of |
|                   |           |   `all` - show all jobs |
|                   |           |   `not-running` - show not running jobs only |
|                   |           |   `running` - show only running jobs |
| `--nodes` | List information about nodes | |
| `--partitions` | List information about partitions | |
| `--reservations` | List information about reservations | |
| `--sort=<sort>` | Sort output by field <sort> in ascending order | see documentation below |
| `--version` | Show version information | |

## Sorting output
Output can be sorted using the `--sort=<sort>` option. `<sort>` is a comma separated list of `<object>:<field>`
`<object>` can be prefixed by a minus sign to reverse the sort order of the field and can be one of:

| *Object* | *Description* |
|:---------|:--------------|
| `jobs` | sort jobs |
| `nodes` | sort nodes |
| `partitions` | sort partitions |
| `reservations` | sort reservations |

The value of `<field>` depends of the `<object>` type used and are described below

### Sorting jobs

| *Field* | *Description* |
|:--------|:--------------|
| `batchhost` | sort by batch host |
| `cpus` | sort by cpus |
| `gres` | sort by GRES |
| `jobid` | sort by job id (*this is the default*) |
| `licenses` | sort by licenses |
| `name` | sort by name |
| `nodes` | sort by nodes |
| `partition` | sort by partitions |
| `reason` | sort by state reason |
| `starttime` | sort by starttime |
| `state` | sort by state |
| `tres` | sort by TRES |
| `user` | sort by user |

### Sorting nodes

| *Field* | *Description* |
|:--------|:--------------|
| `boards` | sort by number of boards |
| `hostname` | sort by hostname |
| `nodename` | sort by node name (*this is the default*) |
| `partition` | sort by partitions |
| `reason` | sort by state reason |
| `slurmversion` | sort by reported SLURM version |
| `sockets` | sort by number of sockets |
| `state` | sort by state |
| `threadsbycore` | sort by threads per core |
| `tresallocated` | sort by allocated TRES |
| `tresconfigured` | sort by configured TRES |

### Sorting partitions

| *Field* | *Description* |
|:--------|:--------------|
| `allocated` | sort by allocated nodes |
| `allocatedpercent` | sort by allocation percentage |
| `idle` | sort by idle nodes |
| `idlepercent` | sort by idle percentage |
| `other` | sort by other nodes |
| `otherpercent` | sort by percentage of other nodes |
| `partition` | sort by partition name (*this is the default*) |
| `total` | sort by total nodes |

### Sorting reservations

| *Field* | *Description* |
|:--------|:--------------|
| `accounts` | sort by accounts |
| `burstbuffers` | sort by burst buffers |
| `corecount` | sort by core count |
| `duration` | sort by duration |
| `end time` | sort by end time |
| `features` | sort by features |
| `flags` | sort by flags |
| `licenses` | sort by licenses |
| `name` | sort by reservation name (*this is the default*) |
| `nodecount` | sort by node count |
| `nodes` | sort by nodes |
| `partition` | sort by partition |
| `starttime` | sort by start time |
| `state` | sort by state |
| `tres` | sort by TRES |
| `users` | sort by users |
| `watts` | sort by watts |

# Licenses
## slurm-qstat

Copyright (C) 2021 by Andreas Maus

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

## tablewriter (https://github.com/olekukonko/tablewriter)

Copyright (C) 2014 by Oleku Konko

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

