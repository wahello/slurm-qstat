**_Note:_** Because I'm running my own servers for several years, main development is done at at https://git.ypbind.de/cgit/slurm-qstat/

----

# Preface
`slurm-qstat` tries to display information about jobs, nodes and partitions of the SLURM queueing system

# Build requirements
This tool is implemented in Go so, obviously, a Go compiler is required.
Additionally the [tablewriter](https://github.com/olekukonko/tablewriter) is required for the build.

# Command line parameters

| *Option* | *Description* | *Note* |
|:---------|:--------------|:-------|
| `--filter=<part>,...` | Limit output of jobs, nodes and partitions to the comma separated list of partitions | If not specified, information about jobs, nodes and partitions of all partitions will be displayed |
| `--help` | Show the help text | |
| `--jobs=<filter>` | List information about jobs | Filter is mandatory and can be one of |
|                   |           |   `all` - show all jobs |
|                   |           |   `not-running` - show not running jobs only |
|                   |           |   `running` - show only running jobs |
| `--nodes` | List information about nodes | |
| `--partitions` | List information about partitions | |
| `--version` | Show version information |

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

