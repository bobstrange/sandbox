#include "all.h"

int main() {
  std::cout << "Integer: " << 42 << "\n"s;

  // 変数
  auto answer = 42;
  std::cout << answer << "\n"s;

  // 関数
  auto print = [](auto x) {
    std::cout << x << "\n"s;
  };

  print(100);
  print("something"s);
}
