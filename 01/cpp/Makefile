CXXFLAGS = -Wall -Werror --std=c++11

main: main.cpp
	$(CXX) $(CXXFLAGS) -o $@ $^

debug: CXXFLAGS += -g
debug: main

all: main

clean:
	rm -rf main main.dSYM

