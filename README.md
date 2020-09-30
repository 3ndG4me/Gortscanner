# Simple-Gort-Scanner
[Simple port scanner](https://github.com/3ndG4me/Simple-Port-Scanner) rewritten in go


## Purpose
Sometimes firewalls are tough and you need the most basic stupid portscanner there is to just test for open ports. This was rewritten in Golang to both improve performance of the python3 version, and just for fun as an exercise porting python to Golang.

It's called Simple "Gort" Scanner because you're supposed to put "go" in the name of golang programs apparently. 

Instead of "po" in "port" you get "go" for "gort", because that's how that works.

## Features
- Can scan single IPs and single ports Example: `gortscanner 192.168.0.1 22`
- Can parse CIDR range and scan multiple ips Example: `gortscanner 192.168.0.1/24 22`
- Can parse port ranges and scan multiple ports Example: `gortscanner 192.168.0.1 1-1024`
- Any combiniation of the above 3
- Cross platform, makes it easy to drop a binary and not deal with python dependencies
- Significantly faster than the python version. Idk by how much, but it's pretty obvious side by side even tweaking delays.
- Better output than version 1.0. Now displays `Host: <host> Ports: <Port/TCP>` for easier parsing.
    - Try `./gortscanner <host(s)> <port(s)> | grep “Host:” | tee scan.out` for a nice easy to cut up report.

## TODO:
- Add the option to parse a list of ports i.e. `gortscanner <IP> 22, 23, 445`
- Add the option to parse a list of IPs i.e. `gortscanner.py 192.168.0.1, 192.168.0.2, 192.168.0.3 <port(s)>`
- Add timeout flag to change the length of timeouts between scans


## Build:
- `go build -o gortscanner main.go`

## Usage:
- `gortscanner <IP> <port>`
- Example: `gortscanner 192.168.0.1/24 1-1024`



