#!/bin/bash

false &
wait $!
echo "false コマンドが終了しました: $?"

true &
wait $!
echo "true コマンドが終了しました: $?"
