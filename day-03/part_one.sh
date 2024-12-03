#!/bin/sh

set -e

input_file="input.txt"

if [ ! -z "$1" ]; then
	input_file=$1
fi

grep -Eo "mul\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\)" $input_file | cut -d "(" -f 2 | cut -d ")" -f 1 | awk -F "," 'BEGIN{result=0} {result+=$1*$2} END{print result}'

