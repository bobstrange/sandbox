gcc_options = -std=c++17 -Wall --pedantic-error

a.out: main.cpp all.h all.h.gch
	g++ $(gcc_options) -include all.h $< -o $@

all.h.gch: all.h
	g++ $(gcc_options) -x c++-header $< -o $@

run: a.out
	./a.out

clean:
	rm -f ./a.out
	rm -f ./all.h.gch

.PHONY: run clean
