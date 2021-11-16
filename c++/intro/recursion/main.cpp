#include "../all.h"

int input() {
  int x {};
  std::cout << "整数を入力してください\n"s;
  std::cin >> x;
  return x;
}

void output(int binary) {
  std::cout << binary << "\n"s;
}

void loop_until_zero() {
  if (input() == 0) {
    return;
  } else {
    loop_until_zero();
  }
}

void until_ten(int x) {
  if (x > 10) {
    return;
  }
  std::cout << x << "\n";
  return until_ten(x + 1);
}

int solve(int n) {
  if (n <= 1) {
    return n;
  }
  return n%10 + 2 * solve(n/10);
}

int convert(int n) {
  if (n > 0) {
    return solve(n);
  }
  return - solve(-n);
}


int main() {
  // loop_until_zero();
  // until_ten(-1);

  // while (true) {
    // auto decimal = input();
    // auto binary = convert(decimal);
    // output(binary);
  // }

  std::cout << "10: "s << solve(10) << "\n"s;
  std::cout << "10110: "s << solve(10110) << "\n"s;
}
