#!/bin/bash

cd -- ${0%/*}
tmp=$(pwd)

cd shellgei160/qdata/10


cat headings.md | sed -E 's/^# (.*)/\1\n===/' | sed -E 's/^## (.*)/\1\n---/'

cd "$tmp"

