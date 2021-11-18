#include "all.h"


// 関数の場合は引数は auto にはできない
auto plus (int x, int y) {
  return x + y;
}

int main() {
  std::cout << "Integer: " << 42 << "\n"s;

  // 変数
  auto answer = 42;
  std::cout << answer << "\n"s;

  // 関数 (ラムダ式の場合は引数も auto にできる)
  auto print = [](auto x) { std::cout << x << "\n"s; };

  print(100);
  print("something"s);

  // 標準エラー出力は、std::

  std::cerr << "標準エラー出力なので > /dev/null しても表示されるよ\n"s;
  std::cerr << "標準エラー出力なので 2> じゃないとリダイレクト出来ないよ\n"s;
}
