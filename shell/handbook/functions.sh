#!/bin/bash

INPUT=$@

function success_log() {
  message=$1
  date=$(date -u +"%Y-%m-%d %H:%M:%S")
  echo "${date} ${message}."
}

error_log() {
  message=$1
  date=$(date -u +"%Y-%m-%d %H:%M:%S")
  echo "${date} [ERROR] ${message}."
}

if [ ! -z $INPUT ]; then
  success_log "Your input was ${INPUT}";
else
  error_log "You need to input something.";
fi
