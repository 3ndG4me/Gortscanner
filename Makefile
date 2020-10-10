# !!!MAKE SURE YOUR GOPATH ENVIRONMENT VARIABLE IS SET FIRST!!!

# Variables
DIR=builds/
GORT=gortscanner
WINGORTLDFLAGS=-ldflags "-H=windowsgui"
W=Windows-x64
L=Linux-x64
A=Linux-arm
M=Linux-mips
D=Darwin-x64

# Make Directory to store executables
$(shell mkdir -p ${DIR})

# Change default to just make for the host OS and add MAKE ALL to do this
default: gort-windows gort-linux gort-darwin

all: default

# Compile Windows binaries
windows: gort-windows

# Compile Linux binaries
linux: gort-linux

# Compile Arm binaries
arm: gort-arm

# Compile mips binaries
mips: gort-mips

# Compile Darwin binaries
darwin: gort-darwin

# Compile gort - Windows x64
gort-windows:
	export GOOS=windows GOARCH=amd64;go build ${WINGORTLDFLAGS} -o ${DIR}/${GORT}-${W}.exe main.go

# Compile gort - Linux mips
gort-mips:
	export GOOS=linux;export GOARCH=mips;go build -o ${DIR}/${GORT}-${M} main.go

# Compile gort - Linux arm
gort-arm:
	export GOOS=linux;export GOARCH=arm;export GOARM=7;go build -o ${DIR}/${GORT}-${A} main.go

# Compile gort - Linux x64
gort-linux:
	export GOOS=linux;export GOARCH=amd64;go build -o ${DIR}/${GORT}-${L} main.go

# Compile gort - Darwin x64
gort-darwin:
	export GOOS=darwin;export GOARCH=amd64;go build -o ${DIR}/${GORT}-${D} main.go

clean:
	rm -rf ${DIR}*

