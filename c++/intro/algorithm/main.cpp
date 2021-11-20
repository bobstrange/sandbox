#include "../all.h"
#include <algorithm>
#include <ios>
#include <iterator>
#include <vector>

auto for_each = [](auto first, auto last, auto f) {
  for (auto i = first; i != last; ++i) {
    f(*i);
  }
};

void for_each_example() {
  std::vector<int> v = {1, 2, 3, 4, 5};

  auto print_value = [](auto x) {
    std::cout << x;
  };

  std::cout << "for_each print_value\n"s;
  for_each(std::begin(v), std::end(v), print_value);
  std::cout << "\n"s;

  auto print_twice_value = [](auto x) {
    std::cout << 2 * x;
  };

  std::cout << "for_each print_twice_value\n"s;
  for_each(std::begin(v), std::end(v), print_twice_value);
  std::cout << "\n"s;

  std::cout << "use std::for_each\n"s;
  std::for_each(std::begin(v), std::end(v), print_value);
  std::cout << "\n"s;

  // mutate original vector
  auto twice = [](auto &x) { x = 2 * x; };
  std::for_each(std::begin(v), std::end(v), twice);

  std::cout << "v is mutated\n"s;
  std::for_each(std::begin(v), std::end(v), [](auto x) { std::cout << x << ", "s; });
  std::cout << "\n"s;
}

void all_of_example() {
  std::vector<int> v = {1, 2, 3, 4, 5};

  auto is_all_even = [](auto first, auto last) {
    return std::all_of(first, last, [](auto x) { return x % 2 == 0; });
  };

  std::cout << "is_all_even {1, 2, 3, 4, 5}:"s << std::boolalpha << is_all_even(std::begin(v), std::end(v)) << "\n"s;

  auto is_all_less_than_ten = [](auto first, auto last) {
    return std::all_of(first, last, [](auto x) { return x < 10; });
  };

  std::cout << "is_all_less_than_ten {1, 2, 3, 4, 5}:"s << is_all_less_than_ten(std::begin(v), std::end(v)) << "\n"s;
}

int main() {
  for_each_example();
  all_of_example();
}
