#!/bin/bash

awk '{a[$1]++;} END {for (i in a){print a[i]"  "i;}}' ../visit.log | sort -nr | head -n3
echo "##############################"
awk '{a[$1]++;} END {for (i in a){print a[i]"  "i;}}' ../visit.log | sort -t " " -k 1 -nr | head -n3