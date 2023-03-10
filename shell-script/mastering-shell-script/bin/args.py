#!/usr/bin/env python3

import sys

count = len(sys.argv)
if count > 1:
    print(f"Arguments suppried: {len(sys.argv)}")
    print(f"Hello {sys.argv[1]}")
print(f"Exiting {sys.argv[0]}")
