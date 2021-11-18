#include "all.h"


// 関数の場合は引数は auto にはできない
auto plus (int x, int y) {
  return x + y;
}

void std_err() {
  std::cerr << "標準エラー出力なので > /dev/null しても表示されるよ\n"s;
  std::cerr << "標準エラー出力なので 2> じゃないとリダイレクト出来ないよ\n"s;
}

void vars_and_funcs() {
  std::cout << "Integer: "s << 42 << "\n"s;

  // 変数
  auto answer = 42;
  std::cout << answer << "\n"s;

  // 関数 (ラムダ式の場合は引数も auto にできる)
  auto print = [](auto x) { std::cout << x << "\n"s; };
  print(100);
  print("something"s);
}

void number() {
  int decimal = 123;

  // 0 をつけると 8 進数になる
  int octal = 0123;

  // 2 進数は 0b or 0B で始まる
  int binary = 0b1010;

  // 16 進数は 0x で始まる
  int hexadecimal = 0x123;

  std::cout << "decimal: "s << decimal << "\n"s;
  std::cout << "binary: "s << binary << "\n"s;
  std::cout << "octal: "s << octal << "\n"s;
  std::cout << "hexadecimal: "s << hexadecimal << "\n"s;
  // 出力は全部 10 進数になる
  // decimal: 123
  // binary: 10
  // octal: 83
  // hexadecimal: 291

  // ruby の　_ みたいな感じで数値を区切れる
  int a = 1000'000;
}

int main() {
  vars_and_funcs();
  std_err();
  number();
}

