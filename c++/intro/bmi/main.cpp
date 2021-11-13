#include "../all.h"

int main() {
  double h = 1.63;
  double m = 72.0;

  auto bmi = [](double height, double mass) {
    return mass / (height * height);
  };

  std::cout << "BMI is " << bmi(h, m) << "\n"s;
}
