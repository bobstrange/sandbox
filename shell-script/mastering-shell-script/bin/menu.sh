#!/bin/bash

while true; do
  clear
  echo "Choose an item: a, b or c"
  echo "a: Backup"
  echo "b: Display Calendar"
  echo "c: Exit"
  read -sn1 -r
  case "$REPLY" in
    a)
      echo "Backing up..."
      sleep 1
    ;;
    b)
      echo "Show calendar"
      sleep 1
    ;;
    c)
      echo "exit"
      sleep 1
      exit 0
    ;;
    *)
      echo "Option should be a, b or c"
    ;;
  esac

  read -n1 -p -r "Press any key to continue..."
done
