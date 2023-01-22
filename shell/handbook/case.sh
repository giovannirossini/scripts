#!/bin/bash

function get_free_memory() {
  FREE_MEMORY=$(free -m | grep Mem | awk '{print $4}')
  echo $FREE_MEMORY
}

function get_used_memory() {
  USED_MEMORY=$(free -m | grep Mem | awk '{print $3}')
  echo $USED_MEMORY
}

function get_total_memory() {
  TOTAL_MEMORY=$(free -m | grep Mem | awk '{print $2}')
  echo $TOTAL_MEMORY
}

case $1 in
  free)
    get_free_memory
    ;;
  used)
    get_used_memory
    ;;
  total)
    get_total_memory
    ;;
  *)
    echo "Usage: $0 {free|used|total}"
    exit 1
esac