#!/bin/bash

# How to run:
# $ ./git_stats.sh "Jan 01 2022" > stats.csv
# will return the data from all repositories in the path since January 1st and create a CSV.

BEGIN_DATE=$1

START=$(date -j -f "%b %d %Y %H:%M:%S" "$BEGIN_DATE 00:00:00" +%s)
END=$(date +%s)
SINCE_DATE=$(( (END-START)/86400 ))

LAST_CELL=1

echo "Project;Lines Added;Lines Deleted;Commits"

for DIR in $(find . -type d -depth 1);
do
  cd $DIR
  ADDED=$(git log --all --numstat --pretty="%H" --since=$SINCE_DATE.days | awk 'NF==3 {added+=$1} END {printf(added)}')
  DELETED=$(git log --all --numstat --pretty="%H" --since=$SINCE_DATE.days | awk 'NF==3 {deleted+=$2} END {printf(deleted)}')
  COMMITS=$(git log --all --numstat --pretty="%H" --since=$SINCE_DATE.days | awk 'NF==1 {total++} END {printf(total)}')

  echo "${DIR/\.\//};$ADDED;$DELETED;$COMMITS"

  cd ..
  LAST_CELL=$((LAST_CELL+1))
done

echo "Total;=SUM(B2:B$LAST_CELL);=SUM(C2:C$LAST_CELL);=SUM(D2:D$LAST_CELL)"
