#!/bin/bash

# to use: get session cookie from AOC page, export as AOC_SESSION in .bashrc/.zshrc
# then run `./input.sh ${day}` from root directory of advent_2022

mkdir $1
url=https://adventofcode.com/2022/day/$1/input
curl ${url} --cookie "session=$AOC_SESSION" > $1/input.txt

cd $1
filename=day$1.go
cp ../3/day3.go $filename

# now implement and execute any `go run` commands from subfolder