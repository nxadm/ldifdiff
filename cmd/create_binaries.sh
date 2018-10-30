#!/bin/bash -e
mkdir -p ../binaries
GOOS=linux GOARCH=amd64 go build -o ../binaries/ldifdiff-linuxamd64 ldifdiff.go 
GOOS=linux GOARCH=386 go build -o ../binaries/ldifdiff-linux386 ldifdiff.go 
GOOS=windows GOARCH=386 go build -o ../binaries/ldifdiff-windows386 ldifdiff.go 
GOOS=windows GOARCH=amd64 go build -o ../binaries/ldifdiff-windowsamd64 ldifdiff.go 
GOOS=darwin GOARCH=amd64 go build -o ../binaries/ldifdiff-darwinamd64 ldifdiff.go 
GOOS=darwin GOARCH=386 go build -o ../binaries/ldifdiff-darwin386 ldifdiff.go 
cd ../binaries
for i in ldifdiff-*; do sha512sum $i > $i.sha512; done
