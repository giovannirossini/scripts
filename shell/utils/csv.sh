#!/bin/env bash

# Print a CSV file in a table format
# Usage:
# csv <file.csv>
# cat example.csv | csv
# csv <(echo "a,b,c"; echo "1,2,3")
function csv() {
  perl -pe 's/((?<=,)|(?<=^)),/ ,/g;' "$@" | column -t -s, 2>/dev/null | less  -F -S -X -K
}