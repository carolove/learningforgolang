#!/bin/bash

awk '{a[$1]+=1;} END {for (i in a) {print a[i]" "i;}}' ../visit.log