#!/bin/sh

# Script for copying VHDL files.
# Input arguments are paths where to look for VHDL files.

for var in "$@"
do
	find "$var" -name '*.vhd' -exec cp {} . \;
done
