#!/usr/bin/env python3

import sys

count = len(sys.argv)

if count == 1:
  name = input("Enter a name: ")
else:
  name = sys.argv[1]

print(f"Hello {name}")
print(f"Exiting {sys.argv[0]}")

log = open("./tmp/script.log", "a")
log.write(f"Hello {name}")
log.close

