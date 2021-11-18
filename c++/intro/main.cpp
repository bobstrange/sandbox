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

  // int
  auto x = 123;

  // uint
  unsigned int uint_x = 123;
  auto uint_x2 = 123u;

  // long
  long int long_x = 123;
  unsigned long int ulong_x = 123;
  auto long_x2 = 123l;
  auto ulong_x2 = 123ul;

  // size
  std::cout << "sizeof(int): "s << sizeof(int) << "\n"s;
  std::cout << "sizeof(x): "s << sizeof(x) << "\n"s;
  std::cout << "sizeof(char): "s << sizeof(char) << "\n"s;
  std::cout << "sizeof(short): "s << sizeof(short) << "\n"s;
  std::cout << "sizeof(long): "s << sizeof(long) << "\n"s;
  std::cout << "sizeof(long long): "s << sizeof(long long) << "\n"s;

  // max and min
  std::cout << "std::numeric_limits<int>::min(): " << std::numeric_limits<int>::min() << "\n"s;
  std::cout << "std::numeric_limits<int>::max(): " << std::numeric_limits<int>::max() << "\n"s;

  std::cout << "std::numeric_limits<unsigned int>::min(): " << std::numeric_limits<unsigned int>::min() << "\n"s;
  std::cout << "std::numeric_limits<unsigned int>::max(): " << std::numeric_limits<unsigned int>::max() << "\n"s;
  std::cout << "std::numeric_limits<short>::min(): " << std::numeric_limits<short>::min() << "\n"s;
  std::cout << "std::numeric_limits<short>::max(): " << std::numeric_limits<short>::max() << "\n"s;
  std::cout << "std::numeric_limits<long>::min(): " << std::numeric_limits<long>::min() << "\n"s;
  std::cout << "std::numeric_limits<long>::max(): " << std::numeric_limits<long>::max() << "\n"s;
}

void floating_point_number() {
  auto print = [](std::basic_string<char> type, std::size_t val) {
    std::cout << "sizeof("s << type << "): "s << val << "\n"s;
  };

  print("float"s, sizeof(float));
  print("double"s, sizeof(double));  print("long double"s, sizeof(long double));

  // 誤差
  float a = 10000.0;
  float b = 0.0001;
  float c = a + b;

  std::cout << "a + b = " << c << "\n"s;

  // リテラル (デフォルトはダブル)
  auto float_x = 123.45f;
  auto double_x = 123.45;
  auto long_double_x = 123.45l;

  // 123.456 の表現
  // 仮数 (fractional part) (coefficient, significand, mantissa)
  // 指数 (exponent part)
  auto val1 = 1.23456e2;
  auto val2 = 123456e-3;
  auto val3 = 123.456e0;
}

int main() {
  vars_and_funcs();
  std_err();
  number();
  floating_point_number();
}

