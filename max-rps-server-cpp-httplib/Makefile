all:
# g++ -std=c++17 main.cpp -o main
	clang++ -O3 -march=native -o main main.cpp


asm:	
	g++ -std=c++17 main.cpp -S -o main.s && code main.s

init:
	g++ --version
	
	# -std=c++17
	g++ -v --help | code -	
	
	# https://stackoverflow.com/questions/34836775/compiler-standards-support-c11-c14-c17
	g++ -v --help 2> /dev/null | sed -n '/^ *-std=\([^<][^ ]\+\).*/ {s//\1/p}'