#include "../all.h"

int main() {

  // 標準入力を受け取るための変数
  double h{};
  std::cout << "身長 (m) を入力してください\n"s;
  std::cin >> h;

  double m{};
  std::cout << "体重 (kg) を入力してください\n"s;
  std::cin >> m;

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
