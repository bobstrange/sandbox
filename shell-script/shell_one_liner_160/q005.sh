#!/bin/bash

# Q. ntp.conf の pool にあるサーバの名前を抽出する

tmp=$PWD

cd shellgei160/qdata/5

cat ntp.conf | grep ^pool | awk '{print $2}'

cd "$tmp"
