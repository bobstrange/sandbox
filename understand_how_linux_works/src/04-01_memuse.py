#!/usr/bin/python3

import subprocess

size = 10000000

print("メモリ確保前のシステム全体のメモリ使用量")
subprocess.run("free")

array = [0] * size

print("メモリ確保後のシステム全体のメモリ使用量")
subprocess.run("free")


