#include "../all.h"
#include <cstddef>

// int main() {
//   std::vector<int> vi;
//   std::vector<double> vd;
//   std::vector<std::string> vs;

//   for (int i = 0; i < 1000; i++) {
//     vi.push_back(i);
//   }

//   std::cout << "vi.size: "s << vi.size() << "\n"s;
//   std::cout << "vi.at(4): "s << vi.at(4) << "\n"s;

//   std::cout << "vi.at(2000): "s << vi.at(2000) << "\n"s;
//   /**
//    * ↑をコメントアウトするとエラーになる terminate called after throwing an instance of 'std::out_of_range'
//   what():  vector::_M_range_check: __n (which is 2000) >= this->size() (which is 1000)
// Makefile:6: recipe for target 'run' failed
//   */
// }

int input() {
  int x{};
  std::cin >> x;
  return x;
}

// int main() {
//   std::vector<int> input_buff;
//   int x{};

//   std::cout << "整数を入力してください\n"s;
//   while ((x = input()) != 0) {
//     input_buff.push_back(x);
//   }

//   std::cout << "input asc\n"s;

//   for (std::size_t index = 0; index < input_buff.size(); index++) {
//     std::cout << input_buff.at(index) << "\n"s;
//   }

//   std::cout << "input desc\n"s;

//   for (std::size_t index = input_buff.size() - 1; index != 0; index--) {
//     std::cout << input_buff.at(index) << "\n"s;
//   }
//   std::cout << input_buff.at(0) << "\n"s;
// }

int main() {
  std::vector<int> v = {8, 3, 7, 4, 2, 9, 3};
  std::size_t size = v.size();

  for (std::size_t head = 0; head < size; head++) {
    std::size_t min = head;

    for (std::size_t index = head + 1; index < size; index++) {
      if (v.at(index) < v.at(min)) {
        min = index;
      }
    }
    auto tmp = v.at(head);
    v.at(head) = v.at(min);
    v.at(min) = tmp;
  }

  for (std::size_t index = 0; index < size; index++) {
    std::cout << v.at(index) << "\n"s;
  }
}
