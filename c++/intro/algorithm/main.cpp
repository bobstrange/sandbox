#include "../all.h"
#include <vector>

auto print_all = [](auto first, auto last, auto f) {
  for (auto i = first; i != last; ++i) {
    f(*i);
  }
};

void for_each() {
  std::vector<int> v = {1, 2, 3, 4, 5};

  auto print_value = [](auto x) {
    std::cout << x;
  };

  print_all(std::begin(v), std::end(v), print_value);

  auto print_twice_value = [](auto x) {
    std::cout << 2 * x;
  };
}

int main() {
  for_each();
}
