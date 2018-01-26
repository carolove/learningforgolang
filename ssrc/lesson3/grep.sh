#!/bin/bash

find *.go | xargs grep -o | uniq -c | sort -nr -k1 | head -n3