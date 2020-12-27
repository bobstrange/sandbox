package go_error_test

import (
	"errors"
	"fmt"
	"testing"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by 0")
	}
	return x / y, nil
}

func TestError01(t *testing.T) {
	_, err := divide(3, 0)
	if err != nil {
		fmt.Println("err is", err)
	} else {
		t.Fatal("divide(3, 0) should return error")
	}
}
