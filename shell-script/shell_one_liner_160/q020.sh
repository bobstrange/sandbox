#!/bin/bash

for file in $(cd /usr; echo *); do
  echo "$file"
done
