#include <gtest/gtest.h>
#include "calc.h"

// Demonstrate some basic assertions.
TEST(HelloTest, BasicAssertions) {
  // Expect two strings not to be equal.
  EXPECT_STRNE("hello", "world");
  // Expect equality.
  EXPECT_EQ(7 * 6, 42);
}

TEST(CalcTest, OnePlusTwoEqualThree) {
  EXPECT_EQ(sum(1, 2), 3);
}

TEST(CalcTest, ThreeByFourEqualTwelve) {
  EXPECT_EQ(calc::multiply(3, 4), 12);
}
