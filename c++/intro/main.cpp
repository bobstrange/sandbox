#include "all.h"
#include <cstdlib>
#include <iostream>
#include <iterator>

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


namespace example_namespace {
  int x{};

  auto sum = [](auto x, auto y) { return x + y; };
}

void use_namespace() {
  example_namespace::x = 50;

  std::cout << "example_namespace::x : "s << example_namespace::x << "\n"s;
  std::cout << "example_namespace::sum(10,20) : "s << example_namespace::sum(10, 20) << "\n"s;

  // namespace の alias
  namespace ns = example_namespace;
  std::cout << "ns::x : "s << ns::x << "\n"s;
  std::cout << "ns::sum(1,2) : "s << ns::sum(1, 2) << "\n"s;

  // mixin 的な
  using namespace std;
  using namespace example_namespace;

  cout << "x: "s << x << "\n"s;
  cout << "sum(10,30): "s << sum(10,30) << "\n"s;

  // typealias
  using Number = int;
  Number a = 123;
}

void use_iterator() {
  std::vector<int> v = { 1, 2, 3, 4, 5 };
  auto itr = std::begin(v);

  int x = *itr; // 先頭の要素 1
  *itr = 100; // 先頭の要素を 100 に変更 v は { 100, 2, 3, 4, 5 } になる

  // iterator の next は ++
  ++itr;
  *itr; // 2
  ++itr;
  *itr; // 3

  // iterator の prev は --
  --itr;
  *itr; // 2

  // 最後の次の要素への iterator
  auto last = std::end(v);

  // iterator は比較できるので、最後の要素の次を指す iterator を利用して loop が書ける
  for (auto iter = std::begin(v); iter != last; ++iter) {
    std::cout << *iter << "\n"s;
  }

  // iterator の interface を使えば、いろいろなクラスのループに対応できる
  auto output_all = [](auto first, auto last) {
    for (auto iter = first; iter != last; ++iter) {
      std::cout << *iter << "\n"s;
    }
  };

  output_all(std::begin(v), std::end(v));

  // std::cout << "数字を入力してください (Ctrl+D で終了)";
  // std::istream_iterator<int> fp( std::cin ), lp;
  // output_all(fp, lp);
}

void lvalue_ref() {
  int a = 1;
  int b = 2;

  b = a;
  std::cout << "a: "s << a <<  " b: "s << b << "\n"s; // a: 1 b: 1

  b = 3;
  std::cout << "a: "s << a <<  " b: "s << b << "\n"s; // a: 1 b: 3

  int x = 1;
  // & をつけると参照渡しにできる
  int & y = x;

  std::cout << "x: "s << x <<  " y: "s << y << "\n"s; // x: 1 y: 1

  y = 3;
  std::cout << "x: "s << x <<  " y: "s << y << "\n"s; // x: 3 y: 3

  x = 5;
  std::cout << "x: "s << x <<  " y: "s << y << "\n"s; // x: 5 y: 5

  auto f = [](int & input) {
    input = 10;
  };

  f(x);

  std::cout << "x: "s << x <<  " y: "s << y << "\n"s; // x: 10 y: 10
}

void lambda() {
  auto message = "hello lambda"s;

  std::cout << "copy capture\n"s;

  [=]() {
    std::cout << message;
  }();

  std::cout << "\n"s;

  std::cout << "ref capture\n"s;

  int x = 0;

  auto increment = [&]() {
    ++x;
  };

  increment();
  std::cout << "x: "s << x << "\n"s; // x: 1

  increment();
  std::cout << "x: "s << x << "\n"s; // x: 2
}

void struct_example() {
  struct Point {
    int x;
    int y;
  };

  Point p1;
  p1.x = 1;
  p1.y = 2;

  Point p2 = p1;
  p2.x = 10;
  p2.y = 20;

  // 参照わたしではなく値わたし
  std::cout << "p1.x "s << p1.x << " p1.y "s << p1.y << "\n"s;
  std::cout << "p2.x "s << p2.x << " p2.y "s << p2.y << "\n"s;

  struct fractional {
    int numerator;
    int denominator;
    double value() {
      return static_cast<double>(numerator) / denominator;
    }
    void set(int numerator_) {
      fractional::numerator = numerator_;
    }
    void set(int numerator_, int denominator_) {
      fractional::numerator = numerator_;
      fractional::denominator = denominator_;
    }
  };

  fractional f1{3, 7};
  std::cout << "3/7 is "s << f1.value() << "\n"s; // 3/7 is 0.42857142857142855

  f1.set(5, 8);
  std::cout << "5/8 is "s << f1.value() << "\n"s; // 5/8 is 0.625
}

void my_array() {
  struct array_int_3 {
    int m0;
    int m1;
    int m2;

    int & operator [](std::size_t i) {
      switch (i) {
        case 0: return m0;
        case 1: return m1;
        case 2: return m2;
        default: std::abort();
      }
    }
  };

  struct array_int_10 {
    int data[10];
    int & operator [](std::size_t i) {
      return data[i];
    }
  };
}

template <typename T>
T twice(T x) {
  return x * 2;
}

template<typename T, std::size_t N>
struct array {
  using value_type = T;
  using reference = T &;
  using size_type = std::size_t;

  value_type storage[N];
  value_type & operator [](size_type i) {
    return storage[i];
  }
};

void template_example() {
  std::cout << "twice(100) "s << twice(100) << "\n"s; // 200
  std::cout << "twice(12.3) "s << twice(12.3) << "\n"s; // 24.6
  std::cout << "twice<int>(12.3) "s << twice<int>(12.3) << "\n"s; // 24

  array<int, 3> a{1, 2, 3};
  std::cout << "a[1] "s << a[1] << "\n"s; // 2
}

void vector_example() {
  // ref: vector の初期化 https://zenn.dev/yohhoy/articles/quiz-init-vector
  std::vector<int> v1(2);
  std::vector<int> v2{2};
  std::vector<int> v3(3, 4);
  std::vector<int> v4{3, 4};

  auto print = [](auto name, auto first, auto last) {
    std::cout << "vector "s << name << "\n"s;
     for (auto iter = first; iter != last; ++iter) {
      std::cout << *iter << "\n"s;
    }
  };

  print("v1"s, std::begin(v1), std::end(v1));
  print("v2"s, std::begin(v2), std::end(v2));
  print("v3"s, std::begin(v3), std::end(v3));
  print("v4"s, std::begin(v4), std::end(v4));


}

int main() {
  vars_and_funcs();
  std_err();
  number();
  floating_point_number();
  use_namespace();
  use_iterator();
  lvalue_ref();
  lambda();
  struct_example();
  my_array();
  template_example();
  vector_example();
}

