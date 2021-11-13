gcc_options = -std=c++17 -Wall --pedantic-error

# Makefile のメモ
# $@ ターゲットは、ターゲットのファイル名
# ↓の場合は a.out となる
# $< ターゲットは、最初の事前条件
# ↓の場合は main.cpp となる
a.out: main.cpp all.h all.h.gch
	g++ $(gcc_options) $< -o $@

all.h.gch: all.h
	g++ $(gcc_options) -x c++-header $< -o $@

run: a.out
	./a.out

clean:
	rm -f ./a.out
	rm -f ./all.h.gch

.PHONY: run clean
