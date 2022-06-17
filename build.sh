#!/bin/bash
PACKAGE=portsrv
ARCH=amd64
declare -A TARGETS=(
    ["mac"]="darwin"
    ["linux"]="linux"
    ["win"]="windows"
)
for name in ${!TARGETS[@]}
do
    OS=${TARGETS[$name]}
    outfile="${PACKAGE}-${name}"
    [ $OS = "windows" ] && outfile="${outfile}.exe"
    echo "Building for $name OS=$OS ARCH=$ARCH outfile=$outfile"
    env GOOS=$OS GOARCH=$ARCH go build -o $outfile 
done