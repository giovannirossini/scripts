#!/bin/bash

# How to run: ./git_stats.sh "Jan 01 2022" will return the data from all repositories in the path since January 1st. 
# This data will only works on MacOS

BEGIN_DATE=$1

START=$(date -j -f "%b %d %Y %H:%M:%S" "$BEGIN_DATE 00:00:00" +%s)
END=$(date +%s)
DAYS=$(( (END-START)/86400 ))

TOTAL_ADDED_LINES=0
TOTAL_DELETED_LINES=0
TOTAL_COMMITS=0

for DIR in $(find . -type d -depth 1);
do
  cd $DIR
  ADDED=$(git log --all --numstat --pretty="%H" --since=$DAYS.days | awk 'NF==3 {added+=$1} END {printf(added)}')
  DELETED=$(git log --all --numstat --pretty="%H" --since=$DAYS.days | awk 'NF==3 {deleted+=$2} END {printf(deleted)}')
  COMMITS=$(git log --all --numstat --pretty="%H" --since=$DAYS.days | awk 'NF==1 {total++} END {printf(total)}')
  TOTAL_ADDED_LINES=$((TOTAL_ADDED_LINES + ADDED))
  TOTAL_DELETED_LINES=$((TOTAL_DELETED_LINES + DELETED))
  TOTAL_COMMITS=$((TOTAL_COMMITS + COMMITS))
  cd ..
done

echo "Total added lines: $TOTAL_ADDED_LINES
Total deleted lines: $TOTAL_DELETED_LINES
Total commits: $TOTAL_COMMITS
Average added lines per commit: $((TOTAL_ADDED_LINES / TOTAL_COMMITS))
Average deleted lines per commit: $((TOTAL_DELETED_LINES / TOTAL_COMMITS))
Average lines per commit: $(((TOTAL_ADDED_LINES + TOTAL_DELETED_LINES) / TOTAL_COMMITS))"
