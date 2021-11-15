#include "../all.h"

int input() {
  int x {};
  std::cout << "整数を入力してください\n" << "0 を入力すると終了します\n";
  std::cin >> x;
  return x;
}

void loop_until_zero() {
  if (input() == 0) {
    return;
  } else {
    loop_until_zero();
  }
}

int main() {
  loop_until_zero();
}
