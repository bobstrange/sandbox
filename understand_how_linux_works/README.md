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

`execve()` -> 別のプログラムを起動する (別のプログラムでメモリを置き換える)

1. `execve()` 関数を呼び出す
2. `execve()` 関数の引数で指定した実行ファイルからプログラムを読み出して、メモリ上に配置するために必要な情報を読み出す
3. 現在のプロセスのメモリを新しいプロセスのデータで上書きする
4. プロセスを新しいプロセスのエントリポイントから実行開始する
  - [fork-and-exec.py](src/006_fork-and-exec.py) の例

実行ファイルが保持する情報
実行ファイルのフォーマット Exectable and Linking Format (ELF) の情報は `readelf` コマンドで得られる

```bash
readelf -h /usr/bin/ruby
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              DYN (Shared object file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x11a0
  Start of program headers:          64 (bytes into file)
  Start of section headers:          12632 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         13
  Size of section headers:           64 (bytes)
  Number of section headers:         29
  Section header string table index: 28
```

`-S` オプションで、コードとデータのオフセット、サイズ、開始アドレスが確認できる
`.text` と `.data`

```bash
readelf -S pause
There are 31 section headers, starting at offset 0x3938:

Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [15] .text             PROGBITS         0000000000401050  00001050
       0000000000000175  0000000000000000  AX       0     0     16
  [25] .data             PROGBITS         0000000000404020  00003020
       0000000000000010  0000000000000000  WA       0     0     8
```

プロセスのメモリマップを見てみる `/proc/<pid>/maps` と、↑で確認したアドレス内にコード(00401050)とデータ(00404020) がおさまっていることが確認できる

```bash
vagrant@vagrant:~$ ./pause &
[1] 3950
vagrant@vagrant:~$ cat /proc/3950/maps
00400000-00401000 r--p 00000000 fd:00 1333442                            /home/vagrant/pause
00401000-00402000 r-xp 00001000 fd:00 1333442                            /home/vagrant/pause
00402000-00403000 r--p 00002000 fd:00 1333442                            /home/vagrant/pause
00403000-00404000 r--p 00002000 fd:00 1333442                            /home/vagrant/pause
00404000-00405000 rw-p 00003000 fd:00 1333442                            /home/vagrant/pause
```

### プロセスの親子関係

`pstree -p` コマンドでプロセスの親子関係を木構造で表示できる ( `-p` で pid を表示)

`ps aux` コマンドの結果

- `START` -> 起動した時刻
- `TIME` -> 使用した CPU 時間
- `STAT`
  - 1 文字目が `S` スリープ
  - 1 文字目が `R` 実行可能状態
  - ゾンビ状態 `Z`

```bash
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  1.1 101992 11380 ?        Ss   13:48   0:01 /sbin/init
root           2  0.0  0.0      0     0 ?        S    13:48   0:00 [kthreadd]
root           3  0.0  0.0      0     0 ?        I<   13:48   0:00 [rcu_gp]
root           4  0.0  0.0      0     0 ?        I<   13:48   0:00 [rcu_par_gp]
```

### プロセスの終了

`exit_group()` システムコール
カーネルがメモリなどの回収をする

プロセスの終了後、親プロセスから `wait()` や `waitpid()` システムコールで、プロセスの状態を確認できる。

[wait-ret.sh](./src/007_wait_ret.sh) を参照

`wait <pid>` で、プロセスの state の変更を待つ

子プロセスが終了してから、親プロセスが子プロセスの状態を確認できる -> 終了したプロセスもシステム上に存在する
終了したが、親が終了状態を確認指定内プロセス -> ゾンビプロセス

親プロセスが wait を実行する前に、終了した場合は orphan (孤児) プロセス

### シグナル

プロセスになにか通知をして、外部から実行の流れを強制的に変える
`SIGINT` シグナルなど

`kill` コマンドで送信できる `kill -INT <pid>` など
他に

`SIGCHILD`: 子プロセス終了時に親プロセスに送信される。一般的にこのシグナルハンドラで `wait()` システムコールを呼ぶ
`SIGSTOP`: プロセスの実行を一時的に停止する `ctrl + z` で止まるやつ -> zsh でうっかりとまるのはこれだったのか、、、
`SIGCONT`: `SIGSTOP` で停止したプロセスの実行を再開する
`SIGKILL`: 死なないプロセスを強制終了する

プロセスはシグナルハンドラを登録して、シグナルを受け取った際の動作を設定できる

```bash
# background job 実行
sleep infinity &
[1] 4070
# もう一つ実行
sleep infinity &
[2] 4071
# jobs で background jobs を確認できる
jobs
[1]-  Running                 sleep infinity &
[2]+  Running                 sleep infinity &
# fg <job_id> で foreground に持ってこられる
fg 1
sleep infinity
# Ctrl-Z で SIGSTOP を送って一時停止する
^Z
[1]+  Stopped                 sleep infinity
# jobs を確認すると job 1 は停止している
jobs
[1]+  Stopped                 sleep infinity
[2]-  Running                 sleep infinity &
# SIGCONT を送信して再開する
kill -CONT 4070
# job_id 1 が再開されている
vagrant@vagrant:~$ jobs
[1]-  Running                 sleep infinity &
[2]+  Running                 sleep infinity &
```
### セッション

セッションは、端末エミュレータや ssh などをとうしてシステムにログインしたときのログインに対応するもの。
すべてのセッションにはセッションを制御するための端末が紐付いている。

通常は pty/<n> という名前の仮想端末がそれぞれのセッションに割り当てられる。
(`pty` は何のことかわかっていなかったが、仮想端末だったのね)

例

- セッション A
  - ログインシェル: bash
  - vim で go を開発しており、go build を事項中
- セッション B
  - ログインシェル: zsh
  - ps aux を実行して less で見ている
- セッション C
  - ログインシェル: zsh
  - calc という計算プログラムを実行

セッションA -> 端末 pty/0
セッションB -> 端末 pty/1
セッションC -> 端末 pty/2

セッションにはセッションリーダープロセスが存在、通常は bash などのシェル
`ps ajx` で確認できる

```bash
ps ajx

PPID     PID    PGID     SID TTY        TPGID STAT   UID   TIME COMMAND

...

3563    3564    3564    3564 pts/0       4094 Ss    1000   0:00 -bash           # <= セッションリーダー SID 3564
3564    4070    4070    3564 pts/0       4094 S     1000   0:00 sleep infinity
3564    4071    4071    3564 pts/0       4094 S     1000   0:00 sleep infinity
3564    4094    4094    3564 pts/0       4094 R+    1000   0:00 ps ajx
```

ここで別で ssh でログインして `ps ajx` を実行すると、別のセッションができていることがわかる。

```bash
ps ajx

PPID     PID    PGID     SID TTY        TPGID STAT   UID   TIME COMMAND

...

4137    4138    4138    4138 pts/1       4147 Ss    1000   0:00 -bash           # <= セッションリーダー SID 4138
4138    4147    4147    4138 pts/1       4147 R+    1000   0:00 ps ajx
```

`TTY` に書いてあるのが端末の名前、 `SID: 3564` の方は、 `pts/0`, `SID: 4138` の方は `pts/1` になっている。

セッションに紐付いている端末がハングアップすると、セッションリーダーに `SIGHUP` が送信される。
シェルは、自分が管理するジョブを終了させてから、自分も終了する。
上の例だと 1 つめのターミナルエミュレーターのウィンドウを閉じた場合に sleep infinity の job なども終了される。

実行に時間がかかるプロセスの実行中に bash が終了したら困る場合は、
`nohup` コマンドを使って `SIGHUP` を無視する設定にした上でプロセスを起動する、bash の `disown` コマンドで実行中のジョブを bash の管理下から外すなどがある。

### プロセスグループ

`kill` コマンドのプロセスID にマイナス値を指定するとプロセスグループにシグナルを投げられる

```bash
ps ajx

PPID     PID    PGID     SID TTY        TPGID STAT   UID   TIME COMMAND
...
   0       1       1       1 ?             -1 Ss       0   0:01 /sbin/init
   0       2       0       0 ?             -1 S        0   0:00 [kthreadd]
...
4253    4300    4253    4253 ?             -1 S     1000   0:00 sshd: vagrant@pts/0
4300    4301    4301    4301 pts/0       4311 Ss    1000   0:00 -bash
4301    4310    4310    4301 pts/0       4311 S     1000   0:00 sleep infinity
4301    4312    4311    4301 pts/0       4311 S+    1000   0:00 less
 723    4313    4313    4313 ?             -1 Ss       0   0:00 sshd: vagrant [priv]
4313    4353    4313    4313 ?             -1 R     1000   0:00 sshd: vagrant@pts/1
4353    4354    4354    4354 pts/1       4363 Ss    1000   0:00 -bash
4354    4363    4363    4354 pts/1       4363 R+    1000   0:00 ps ajx
```

- フォアグラウンドプロセスグループ
  - セッションに 1 つだけ存在し、セッションの端末に直接アクセスできる。
  - STAT フィールドに `+` があるもの (↑の例だと less や ps ajx)
- バックグラウンドプロセスグループ
  - バックグラウンドジョブ
  - バックグラウンドプロセスが端末を操作しようとすると `SIGSTOP` を受けたときのように実行が一時的に中断され、 `fg` などによりフォアグラウンドプロセスグループになるまでこの状態が続く

### daemon

daemon -> 常駐プロセス

特徴

- 端末から入出力する必要がないため、端末が割り当てられていない
- あらゆるログインセッションが終了しても影響を受けないように、独自のセッションを持つ
- daemon を生成したプロセスが、daemon の終了を気にしなくて良いように init が親になっている

`sshd` の例 -> PPID が 1 (init) SID も自分の PID と一緒
daemon は端末を持たないので、慣習として `SIGHUP` を設定ファイルの読み込み直しのシグナルとして使うことが多い

```bash
ps ajx

PPID     PID    PGID     SID TTY        TPGID STAT   UID   TIME COMMAND
...
   1     723     723     723 ?             -1 Ss       0   0:00 sshd: /usr/sbin/sshd -D [listener] 0 of 10-100 startups
```

## プロセススケジューラー

1 つの論理 CPU 上で同時に動くプロセスは 1 つだけ
実行可能な複数のプロセスに、タイムスライスと呼ばれる単位で順番に CPU を使わせる

### 経過時間と使用時間

適度な回数空ループをする[load.py](src/008_load.py) を `time` コマンドつきで実行

```bash
time ./src/008_load.py

real    0m2.455s
user    0m2.435s
sys     0m0.019s
```

real -> 経過時間
user -> プロセスがユーザーランドで動作していた時間
sys -> プロセスによるシステムコールで、カーネルが動作していたときの時間

プロセスの開始時や、終了時に Python インタープリタがシステムコールを実行する分 sys の時間も多少ある

```bash
time sleep 3

real    0m3.001s
user    0m0.001s
sys     0m0.000s
```

CPU をほぼ使わない `sleep` だと↑のようになる

並列実行をするスクリプト [multiload.sh](./src/009_multiload.sh) を実行すると 1 CPU では、並行度を増やすと user の時間はほぼ変わらないが実際に処理にかかる時間 real 2 倍、3 倍になる

```bash
src/009_multiload.sh 3

real    0m7.091s
user    0m2.365s
sys     0m0.000s

real    0m7.121s
user    0m2.379s
sys     0m0.000s

real    0m7.244s
user    0m2.503s
sys     0m0.000s
```

複数 CPU を使う様にする
仮想マシンの VCPU が 2 なので↓のようになる。
3 の場合も CPU を使い切り 2/3 の時間で終わっている

```bash
src/009_multiload.sh -m 1

real    0m2.306s
user    0m2.306s
sys     0m0.000s

src/009_multiload.sh -m 2

real    0m2.450s
user    0m2.442s
sys     0m0.008s

real    0m2.509s
user    0m2.500s
sys     0m0.000s

src/009_multiload.sh -m 3

real    0m3.594s
user    0m2.520s
sys     0m0.000s

real    0m3.614s
user    0m2.397s
sys     0m0.012s

real    0m3.861s
user    0m2.540s
sys     0m0.004s
```

### タイムスライス

### コンテキストスイッチ

プロセスの実行中に、タイムスライスが切れた場合は必ずコンテキストスイッチが発生する
コード中のある行の実行とその次の行の実行が連続して行われるとは限らない

### 性能

- ターンアラウンドタイム: システムに処理を依頼してから個々の処理が終わるまでの時間
- スループット: 単位時間あたりに処理を終えられる数

論理 CPU の数よりプロセス数を多くしても、平均ターンアラウンドタイムが長くなり、スループットは向上しない
(逆に、プロセス数を多くしていくとコンテキストスイッチの影響でスループットが下がる)

論理 CPU の数よりプロセス数が少ない状態では、平均ターンアラウンドタイムは線形に上昇する

## メモリ管理システム

### メモリ管理情報の習得

`free` コマンドでメモリについて確認できる

- total: 全メモリ
- free: 見かけ上の空きメモリ
- buff/cache: バッファキャッシュ、ページキャッシュが利用するメモリ。free が減ってきたらカーネルによって解放される
- available: 実質的な空きメモリ free + 解放可能なメモリを足したもの
- used: システもが使用中のメモリから buff/cache を引いたもの

```bash
free
              total        used        free      shared  buff/cache   available
Mem:        1000068      211420      113240        1000      675408      637488
Swap:       1999868         780     1999088
```

p75 参照

<------------------------------------------- total --------------------------------------------------->
<------ カーネルが使用中 ------>
                <---------------- available ------------------>
<-- 解放不可 --><-- 解放可能 --><---------- free -------------><--------- プロセスが使用中 ----------->
        <- buff/cache ->

used

[memuse.py](src/04-01_memuse.py) で
メモリを確保する前に free を実行し、適当にメモリを確保後に free を再度実行したときの挙動

メモリ確保後は、`used` が  81MiB ((298412 - 215372) / 1024) 程度増加
プログラム終了後に free を実行すると used の値はほぼ元に戻る

```bash
./src/04-01_memuse.py
メモリ確保前のシステム全体のメモリ使用量
              total        used        free      shared  buff/cache   available
Mem:        1000068      215372      105452        1000      679244      633292
Swap:       1999868         780     1999088
メモリ確保後のシステム全体のメモリ使用量
              total        used        free      shared  buff/cache   available
Mem:        1000068      298412       63860        1000      637796      551944
Swap:       1999868         780     1999088

free
              total        used        free      shared  buff/cache   available
Mem:        1000068      211816      182860         972      605392      639040
Swap:       1999868         780     1999088
```

buff/cache

ストレージアクセス >> メモリアクセス なので、ストレージのデータをメモリーに乗せるとメモリ上にデータをキャッシュしておく
[buff-cache.sh](src/04-02_buff-cache.sh) を実行すると
100 MiB のファイルを作成する時点で buff/cache が 100 MiB 増え、ファイルを削除した時点で、buff/cache が 100 MiB 減る

```bash
/src/04-02_buff-cache.sh
ファイル作成前のメモリ使用量を表示
              total        used        free      shared  buff/cache   available
Mem:        1000068      198804      358792        1012      442472      653836
Swap:       1999868         524     1999344
100 MiB のファイルを新規作成 -> カーネルは 100 MiB のページキャッシュを確保
100+0 records in
100+0 records out
104857600 bytes (105 MB, 100 MiB) copied, 0.0532288 s, 2.0 GB/s
ページキャッシュ確保後のメモリ使用量を表示
              total        used        free      shared  buff/cache   available
Mem:        1000068      198904      253564        1012      547600      652312
Swap:       1999868         524     1999344
ファイルを削除 -> カーネルは 100 MiB のページキャッシュを解放後のメモリ使用量を表示
              total        used        free      shared  buff/cache   available
Mem:        1000068      199612      357848        1012      442608      652932
Swap:       1999868         524     1999344
```

`sar` コマンドでメモリの情報を見ることもできる

```bash
sar -r 1 5
Linux 5.4.0-110-generic (vagrant)       10/25/2022      _x86_64_        (2 CPU)

01:05:00 PM kbmemfree   kbavail kbmemused  %memused kbbuffers  kbcached  kbcommit   %commit  kbactive   kbinact   kbdirty
01:05:01 PM    344600    646584    154784     15.48     40176    376696   1188136     39.61    183176    345748       584
01:05:03 PM    344600    646584    154784     15.48     40176    376696   1188136     39.61    183176    345748       584
01:05:04 PM    344600    646584    154784     15.48     40176    376696   1188136     39.61    183176    345748       584
01:05:05 PM    350876    652860    148576     14.86     40176    376696   1173920     39.13    176920    345748       584
01:05:06 PM    350956    652940    148548     14.85     40176    376696   1173920     39.13    176864    345744       584
Average:       347126    649110    152295     15.23     40176    376696   1182450     39.42    180662    345747       584
```

free メモリが少なくなってきた場合に、カーネルは buff/cache の内回収可能なメモリ領域(ディスクからデータを読み出してから変更していないページキャッシュなど)を解放して、 free の値を増やそうとする。

それでもメモリが足りない場合は、OOM (Out Of Memory) 状態になり、OOM Killer によりプロセスが強制終了されてプロセスが使用していたメモリを解放しようとする。
OOM Killer が動作した後は、`dmesg` で見られるカーネルログに、ログが残る。

### 仮想記憶

仮想記憶が無いと

- メモリの断片化
  - 物理メモリが断片化していても、仮想メモリ上では 1 つの領域としてみなせる
- マルチプロセスの実現
  - プロセス毎に独立したアドレス空間を持てる
- 不正な領域へのアクセス
  - 他のプロセスのメモリにはアクセスできない (page fault)

の課題がある。

仮想記憶: プロセスがメモリにアクセスする際に、システムに搭載されているメモリ(物理アドレス)に直接アクセスさせるのではなく、仮想アドレスを用いて間接的にアクセスさせる。

### ページテーブル

仮想アドレスから物理アドレスへの変換表
ページのサイズ x86_64 なら 4KiB

ページテーブルの仮想アドレスに対応する物理アドレスが未割り当ての場合、その仮想アドレスにプロセスがアクセスすると CPU 上で page fault 例外が発生し、CPU で実行中の命令が中断され、カーネルのメモリに配置された page fault handler が実行される。

ページテーブルエントリが存在しないページに対するアクセスには `SIGSEGV`、そうでない場合は、対応する物理メモリを割り当てる。
不正なアドレスにアクセスする [`segv`](./src/04-03_segv.go) と page fault handler は `SIGSEGV` シグナルをプロセスに送信し、プロセスは強制終了される。

```bash
./04-03_segv
不正メモリアクセス前
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x48cf0b]

goroutine 1 [running]:
main.main()
        /home/vagrant/src/04-03_segv.go:8 +0x7b
```

### プロセスへの新規のメモリの割り当て

1. メモリ領域の割り当て: 仮想空間に新規にアクセス可能なメモリ領域をマップ
2. メモリの割り当て: ↑のメモリ領域に物理メモリを割り当てる

`mmap()` システムコール: メモリ領域の割り当て

メモリの領域の確保を確認するプログラム [`mmap`](./src/04-05_mmap.go)

`/proc/<pid>/maps` は、各行が個々のメモリ領域、第一フィールドがメモリ領域

```bash
./04-05_mmap
新規メモリ領域獲得前のメモリマップ
00400000-004ab000 r-xp 00000000 fd:00 1841942                            /home/vagrant/04-05_mmap
004ab000-00584000 r--p 000ab000 fd:00 1841942                            /home/vagrant/04-05_mmap
00584000-0059b000 rw-p 00184000 fd:00 1841942                            /home/vagrant/04-05_mmap
0059b000-005b9000 rw-p 00000000 00:00 0
c000000000-c004000000 rw-p 00000000 00:00 0
7f7f08984000-7f7f0ac35000 rw-p 00000000 00:00 0
7ffcd2342000-7ffcd2363000 rw-p 00000000 00:00 0                          [stack]
7ffcd23ba000-7ffcd23bd000 r--p 00000000 00:00 0                          [vvar]
7ffcd23bd000-7ffcd23be000 r-xp 00000000 00:00 0                          [vdso]
ffffffffff600000-ffffffffff601000 --xp 00000000 00:00 0                  [vsyscall]

=== 新規メモリ領域: アドレス = 0x7f7f02584000, サイズ = 0x6400000 ===

新規メモリ領域獲得後のメモリマップ
00400000-004ab000 r-xp 00000000 fd:00 1841942                            /home/vagrant/04-05_mmap
004ab000-00584000 r--p 000ab000 fd:00 1841942                            /home/vagrant/04-05_mmap
00584000-0059b000 rw-p 00184000 fd:00 1841942                            /home/vagrant/04-05_mmap
0059b000-005b9000 rw-p 00000000 00:00 0
c000000000-c004000000 rw-p 00000000 00:00 0
7f7f02584000-7f7f0ac35000 rw-p 00000000 00:00 0
7ffcd2342000-7ffcd2363000 rw-p 00000000 00:00 0                          [stack]
7ffcd23ba000-7ffcd23bd000 r--p 00000000 00:00 0                          [vvar]
7ffcd23bd000-7ffcd23be000 r-xp 00000000 00:00 0                          [vdso]
ffffffffff600000-ffffffffff601000 --xp 00000000 00:00 0                  [vsyscall]
```

前
7f7f08984000-7f7f0ac35000 rw-p 00000000 00:00 0
後
7f7f02584000-7f7f0ac35000 rw-p 00000000 00:00 0

### デマンドページング

`mmap()` システムコール呼び出し直後は、新規メモリに対応する物理メモリは割り当てられていない。
新規領域の各ページに最初にアクセスしたときに、物理メモリを割り当てる。

100 MiB メモリを確保し、確保した領域に 1 秒ごとに 10 MiB ずつアクセスする [`demand-paging.py`](./src/04-06_demand-paging.py)
を動作させながら `sar -r` でメモリの変化量を見てみる

```bash
14:54:25: 新規メモリ領域獲得前. Enter キーを押すと 100 MiB の新規メモリ領域を獲得する.

14:54:26: 新規メモリ領域獲得. Enter キーを押すと 1 秒に 10 MiB づつ、合計 100 MiB の新規メモリ領域にアクセスする.

14:54:28: メモリ領域アクセス中. 0.0 MiB アクセス済み.
14:54:29: メモリ領域アクセス中. 10.0 MiB アクセス済み.
14:54:30: メモリ領域アクセス中. 20.0 MiB アクセス済み.
14:54:31: メモリ領域アクセス中. 30.0 MiB アクセス済み.
14:54:32: メモリ領域アクセス中. 40.0 MiB アクセス済み.
14:54:33: メモリ領域アクセス中. 50.0 MiB アクセス済み.
14:54:34: メモリ領域アクセス中. 60.0 MiB アクセス済み.
14:54:35: メモリ領域アクセス中. 70.0 MiB アクセス済み.
14:54:36: メモリ領域アクセス中. 80.0 MiB アクセス済み.
14:54:37: メモリ領域アクセス中. 90.0 MiB アクセス済み.
14:54:38: 新規獲得したメモリ領域にすべてアクセスしました
```

```bash
sar -r 1
Linux 5.4.0-110-generic (vagrant)       10/26/2022      _x86_64_        (2 CPU)

02:54:21 PM kbmemfree   kbavail kbmemused  %memused kbbuffers  kbcached  kbcommit   %commit  kbactive   kbinact   kbdirty
02:54:22 PM    209760    641044    152464     15.25     50240    489308   1176936     39.23    241040    406680         0
02:54:23 PM    209760    641044    152464     15.25     50240    489308   1176936     39.23    241096    406680         0
02:54:24 PM    209760    641044    152464     15.25     50240    489308   1176936     39.23    241096    406680         0
02:54:25 PM    209760    641044    152464     15.25     50240    489308   1176936     39.23    241096    406680         0
02:54:26 PM    206224    637508    156000     15.60     50240    489308   1283500     42.78    244664    406680         0
02:54:27 PM    206224    637508    156000     15.60     50240    489308   1283500     42.78    244664    406680         0
02:54:28 PM    206224    637508    156000     15.60     50240    489308   1283500     42.78    244664    406680         0
02:54:29 PM    196648    627932    165576     16.56     50240    489308   1283500     42.78    254864    406680         0
02:54:30 PM    186568    617852    175656     17.56     50240    489308   1283500     42.78    265064    406680         0
02:54:31 PM    176236    607520    185988     18.60     50240    489308   1283500     42.78    275384    406680         0
02:54:32 PM    165904    597188    196320     19.63     50240    489308   1283500     42.78    285584    406680         0
02:54:33 PM    155824    587108    206400     20.64     50240    489308   1283500     42.78    295844    406680         0
02:54:34 PM    145492    576776    216732     21.67     50240    489308   1283500     42.78    306044    406680         0
02:54:35 PM    135412    566696    226812     22.68     50240    489308   1283500     42.78    316304    406680         0
02:54:36 PM    125080    556364    237144     23.71     50240    489308   1283500     42.78    326504    406680         0
02:54:37 PM    115504    546788    246720     24.67     50240    489308   1283500     42.78    336764    406680         0
02:54:38 PM    210484    641768    151740     15.17     50240    489308   1177184     39.24    241156    406680         0
02:54:39 PM    210516    641800    151708     15.17     50240    489308   1177184     39.24    241096    406680         0
```

- 100 MiB 確保 (14:54:26) 時点でも、確保した領域にアクセスされるまでは `kbmemused` は変化しない
- 14:54:28 ~ メモリを 10 MiB ずつ確保するので `kbmused` は秒間 10 MiB ずつ増える
- 14:54:38 に終了すると `kbmemused` は元の値にもどる