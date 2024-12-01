#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <sstream>
#include <cstdlib>

#define ABS(X)	(((X) < 0) ? -1 * (X) : (X))

static bool cmp(int a, int b) {
	return a < b;
}

bool ReadInts(std::string path, std::vector<int>& left, std::vector<int>& right) {
	std::ifstream in(path, std::ios_base::in);
	std::string line, word;
	while (getline(in, line)) {
		std::stringstream ss(line);
		ss >> word;
		left.push_back(atoi(word.c_str()));
		ss >> word;
		right.push_back(atoi(word.c_str()));
	}
	std::sort(left.begin(), left.end(), cmp);
	std::sort(right.begin(), right.end(), cmp);
	return true;
}

int Part1(std::vector<int>& left, std::vector<int>& right) {
	int total = 0;
	if (left.size() != right.size()) {
		std::cout << "Lengths are not the same" << std::endl;
		exit(-1);
	}
	for (size_t i = 0; i < left.size(); i++) {
		total += ABS(left[i] - right[i]);
	}
	return total;
}

int Part2(std::vector<int>& left, std::vector<int>& right) {
	int total = 0;
	if (left.size() != right.size()) {
		std::cout << "Lengths are not the same" << std::endl;
		exit(-1);
	}
	std::vector<int> counts;
	counts.resize(*right.end() + 1);
	for (auto i = right.begin(); i != right.end(); i++)
		counts[*i]++;
	for (auto i = left.begin(); i != left.end(); i++) {
		total += *i * counts[*i];
	}
	return total;
}

int main(int argc, char *argv[]) {
	if (--argc != 2) {
		std::cout << "Usage: ./main [1|2] file" << std::endl;
		return -1;
	}
	std::string path = argv[2];
	bool part = strcmp(argv[1], "1") == 0;
	std::vector<int> left, right;
	if (!ReadInts(path, left, right))
		return -1;
	int score;
	if (part) {
		score = Part1(left, right);
		std::cout << "Total Difference: " << score << std::endl;
	} else {
		score = Part2(left, right);
		std::cout << "Similarity Score: " << score << std::endl;
	}
	return 0;
}
