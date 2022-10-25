#!/bin/bash

echo "ファイル作成前のメモリ使用量を表示"

free

echo "100 MiB のファイルを新規作成 -> カーネルは 100 MiB のページキャッシュを確保"
dd if=/dev/zero of=testfile bs=1M count=100

echo "ページキャッシュ確保後のメモリ使用量を表示"

free

echo "ファイルを削除 -> カーネルは 100 MiB のページキャッシュを解放後のメモリ使用量を表示"

rm testfile

free
