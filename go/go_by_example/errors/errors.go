package main

import (
	"errors"
	"fmt"
)

// Go の規約で、エラーは戻り値の最後かつ error 型
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New() でエラー値を作れる
		return -1, errors.New("Can't work with 42")
	}
	// nil 値を返すことで、エラーではないことを表現する
	return arg + 3, nil
}

type ArgError struct {
	Arg  int
	Prob string
}

// Error() メソッドを実装する型を作ることで、カスタムエラーを実装できる
func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.Arg, e.Prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &ArgError{Arg: arg, Prob: "Can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, num := range []int{7, 42} {
		if r, e := f1(num); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, num := range []int{7, 42} {
		if r, e := f2(num); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)

	// カスタムエラーのデータを使う場合は、TypeAssertion で型を絞り込む必要がある
	if ae, ok := e.(*ArgError); ok {
		fmt.Println(ae.Arg)
		fmt.Println(ae.Prob)
	}
}
