#include "../all.h"

int main() {
  double h = 1.63;
  double m = 72.0;

  auto calcbmi = [](double height, double mass) {
    return mass / (height * height);
  };

  double bmi = calcbmi(h, m);

  std::cout << "BMI is " << bmi << "\n"s;

  auto status = [](double bmi) {
    if (bmi < 18.5) {
      return "やせすぎ"s;
    } else if (bmi < 25.0) {
      return "普通"s;
    } else if (bmi < 30.0) {
      return "太り気味"s;
    } else {
      return "肥満"s;
    }
  };

  std::cout << "Status is " << status(bmi) << "\n"s;
}
