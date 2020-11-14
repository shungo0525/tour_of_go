package main

import (
	"fmt"
	"math"
)

// 関数の型は引数の後に記載する
func add(x, y int) int {
	return x + y
}

// 複数の戻り値を返す
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 初期値の設定
var i, j int = 1, 2

func main() {
	fmt.Println(math.Pi)

	fmt.Println(add(1,3))

	// a, b := swap("hello", "world")
	// fmt.Println(a, b)

	fmt.Println(split(17))

	// 初期値がある場合、型を省略
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// 暗黙的な型宣言
	k := 3
	fmt.Println("k:", k)


	var i int
	var f float64
	var b bool
	var s string
	// Printfはフォーマットを指定して出力する %v:値, %T:型
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	const Pi = 3.14
	fmt.Println(Pi)
}
