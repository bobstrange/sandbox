#!/bin/bash

echo I am a perfect human  | sed -E 's/(.*)/\U\1/'

echo pen-pineapple-apple-pen | sed -E 's/^(.)/\U\1/' | sed -E 's/-(.)/-\U\1/g'
