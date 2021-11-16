#include "../all.h"

int main() {
  std::vector<int> vi;
  std::vector<double> vd;
  std::vector<std::string> vs;

  for (int i = 0; i < 1000; i++) {
    vi.push_back(i);
  }

  std::cout << "vi.size: "s << vi.size() << "\n"s;
  std::cout << "vi.at(4): "s << vi.at(4) << "\n"s;

  std::cout << "vi.at(2000): "s << vi.at(2000) << "\n"s;
  /**
   * ↑をコメントアウトするとエラーになる terminate called after throwing an instance of 'std::out_of_range'
  what():  vector::_M_range_check: __n (which is 2000) >= this->size() (which is 1000)
Makefile:6: recipe for target 'run' failed
  */
}
