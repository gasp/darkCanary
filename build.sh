#!/bin/bash
# ./setup.sh darkCanary web-browser.svg
./setup.sh darkCanary web-browser.svg
go build darkCanary.go
mv darkCanary darkCanary.app/Contents/MacOS/darkCanary
