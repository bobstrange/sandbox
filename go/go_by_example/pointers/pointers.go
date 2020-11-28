package main

import (
	"fmt"
)

// ival は値なので、呼び出し元からコピーされる
func zeroval(ival int) {
	ival = 0
}

// iptr はポインタなので、 *iptr に新しい値を設定すると、呼び出し元の参照も変更される
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // &i は i のメモリ上のアドレスということ
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer", &i)
}
