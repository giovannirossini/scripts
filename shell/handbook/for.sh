#!/bin/bash

words=(
  "Simple "
  "for "
  "using "
  "bash."
)

for word in ${words[@]};
do
  echo -n "${word} ";
done

for i in {1..10};
do
  echo -n "${i}, "
done