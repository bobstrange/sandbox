gcc_options = -std=c++17 -Wall --pedantic-error

a.out: main.cpp
	g++ $(gcc_options) -o $@ $<
run: a.out
	./$<
clean:
	rm -f ./a.out
.PHONY: run clean
