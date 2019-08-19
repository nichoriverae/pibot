#!/bin/bash

for i in "$*"
do
	eval `ARM=6 GOARCH=arm GOOS=linux go build $i`
done
