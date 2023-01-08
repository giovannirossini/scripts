#!/bin/bash

INPUT=$@

if [[ -z $INPUT ]]; then
  echo "You need to input something. Try again."
else
  echo "Your input is ${INPUT}."
fi
