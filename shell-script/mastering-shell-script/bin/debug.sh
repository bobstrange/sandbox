#!/bin/bash


# %w   週の何日目かを表す値 (0..6); 0 を日曜日とする
day_of_the_week=$(date +%w)
date=$(date --iso-8601='seconds')

if [ "${day_of_the_week}" -eq 0 ]; then
  echo "Today is Sunday: ${date}"
else
  echo "Today is not Sunday: ${date}"
fi



