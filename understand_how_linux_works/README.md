# 試して理解 Linux のしくみ

## Linux の概要

### CPU のモード

- カーネルモード -> 制限なし
  - Linux の場合カーネルのみ
- ユーザーモード -> 制限あり

### システムコール

プロセスがカーネルに処理を依頼するための方法

```
プロセス     =================                 =================
          システムコール発行 ↓                ↑ システムコールから復帰
カーネル                       ================

CPU のモード | ユーザーモード | カーネルモード | ユーザーモード |
```

### システムコール発行の可視化

`strace` コマンドで、発行されるシステムコールを見ることができる

`-T` で発行されたシステムコールの所要時間
`-tt` でシステムコールが発行された日時

```bash
strace -o hello.log ./hello
```

`sar` コマンドで CPU の使用状況を見ることができる
sar Collect, report or save system activity information なので、system activity report の略？

```bash
sar -P 0 1 1
```

`-P 0` -> CPU0
`1` 1 秒ごと
`1` 1 回

結果

```bash
vagrant ssh -c "sar -P 0 1 5"
Linux 5.4.0-110-generic (vagrant)       10/21/2022      _x86_64_        (2 CPU)

02:03:43 PM     CPU     %user     %nice   %system   %iowait    %steal     %idle
02:03:44 PM       0      0.00      0.00      0.00      0.00      0.00    100.00
02:03:45 PM       0      0.00      0.00      0.00      0.00      0.00    100.00
02:03:46 PM       0      0.00      0.00      0.00      0.00      0.00    100.00
```

%user %nice がユーザーモードでプロセスを実行している時間の割合
%system がカーネルがシステムコールを処理している時間の割合

`taskset` コマンドで指定した CPU 上でコマンドを実行できるので、無限ループのプログラムを実行させて
`sar` コマンドで確認する。

```bash
chmod +x ./src/inf-loop.py
taskset -c 0 ./src/inf-loop.py &
```

```
vagrant@vagrant:~$ sar -P 0 1 10
Linux 5.4.0-110-generic (vagrant)       10/21/2022      _x86_64_        (2 CPU)

02:09:26 PM     CPU     %user     %nice   %system   %iowait    %steal     %idle
02:09:27 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:28 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:30 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:31 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:32 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:33 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:34 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:37 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:38 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
02:09:39 PM       0    100.00      0.00      0.00      0.00      0.00      0.00
Average:          0    100.00      0.00      0.00      0.00      0.00      0.00
vagrant@vagrant:~$ sar -P 1 1 10
Linux 5.4.0-110-generic (vagrant)       10/21/2022      _x86_64_        (2 CPU)

02:09:47 PM     CPU     %user     %nice   %system   %iowait    %steal     %idle
02:09:48 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
02:09:50 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
02:09:51 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
02:09:52 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
02:09:53 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
02:09:54 PM       1      0.64      0.00      0.00      0.00      0.00     99.36
02:09:56 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
02:09:57 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
02:09:58 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
02:10:00 PM       1      0.00      0.00      0.00      0.00      0.00    100.00
Average:          1      0.08      0.00      0.00      0.00      0.00     99.92
```

%user が 100 になっている
CPU1 は何も動いていないので %idle 100 のまま

今度は、ループ内で、systemcall を発行する無限ループ

```bash
taskset -c ./src/003_syscall-inf-loop.py &
sar -P 0 1 5

Linux 5.4.0-110-generic (vagrant)       10/21/2022      _x86_64_        (2 CPU)

02:16:13 PM     CPU     %user     %nice   %system   %iowait    %steal     %idle
02:16:16 PM       0     69.70      0.00     30.30      0.00      0.00      0.00
02:16:18 PM       0     65.62      0.00     34.38      0.00      0.00      0.00
02:16:19 PM       0     47.83      0.00     52.17      0.00      0.00      0.00
02:16:20 PM       0     62.26      0.00     37.74      0.00      0.00      0.00
02:16:21 PM       0     64.10      0.00     35.90      0.00      0.00      0.00
Average:          0     63.85      0.00     36.15      0.00      0.00      0.00
```

すると %system の割合が増える

### ライブラリ

`ldd` コマンドで、shared object の依存を表示できる

```bash
ldd /usr/bin/ruby
        linux-vdso.so.1 (0x00007ffca9bfe000)
        libruby-2.7.so.2.7 => /lib/x86_64-linux-gnu/libruby-2.7.so.2.7 (0x00007f65638d1000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f65636df000)
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f65636bc000)
        librt.so.1 => /lib/x86_64-linux-gnu/librt.so.1 (0x00007f65636b2000)
        libgmp.so.10 => /lib/x86_64-linux-gnu/libgmp.so.10 (0x00007f656362e000)
        libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f6563628000)
        libcrypt.so.1 => /lib/x86_64-linux-gnu/libcrypt.so.1 (0x00007f65635eb000)
        libm.so.6 => /lib/x86_64-linux-gnu/libm.so.6 (0x00007f656349c000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f6563c31000)
```

`libc` 標準 C ライブラリや、システムコールのラッパー関数を提供

### 静的ライブラリと共有ライブラリ

- プログラムの生成
  - ソースコードをコンパイル -> オブジェクトファイル
  - オブジェクトファイルが使うライブラリをリンク -> 実行ファイル

静的ライブラリのリンク -> リンク時にプログラムに組み込まれる
共有ライブラリのリンク -> リンク時にプログラムに I/F の情報だけ組み込まれる

例: [pause.c](./src/004_pause.c)

libc の静的ライブラリ `libc.a` を使う場合
コンパイルしたあと `ls -l` でサイズを `ldd` で共有オブジェクトの依存を表示

```bash
cc -static -o pause src/004_pause.c
ls -l pause
-rwxrwxr-x 1 vagrant vagrant 871832 Oct 21 14:39 pause
ldd pause
        not a dynamic executable
```

`libc.so` を使う場合

```bash
cc -o pause src/004_pause.c
ls -l pause
-rwxrwxr-x 1 vagrant vagrant 16704 Oct 21 14:44 pause
ldd pause
        linux-vdso.so.1 (0x00007fff3f5dc000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fb02f3a8000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fb02f5ac000)
```

`ls -l` でサイズを確認すると小さくなっており、 `ldd` で共有ライブラリが使われていることがわかる。

## プロセス管理

### プロセスの生成

`fork()` 関数 と `execve()` 関数

`fork()` -> 発行したプロセスのコピーを作る (メモリコピー (※ copy on write により低コスト))

1. 親プロセスが `fork()` 関数を呼ぶ
2. カーネルが子プロセス用のメモリ領域を確保して、親プロセスのメモリをコピーする
3. 親プロセス、子プロセスが `fork()` 関数から復帰する
  - 親プロセスは戻り値として子プロセスの PID 子プロセスは 0 を得る
  - [fork.py](./src/005_fork.py) の例を見ると良い
