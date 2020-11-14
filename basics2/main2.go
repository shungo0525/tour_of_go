package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Pow(x, n) -> xのn乗
func pow(x, n, lim float64) float64 {
	// ifのスコープ内で変数を定義可能
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// OSを取得
	fmt.Println(runtime.GOOS)

	fmt.Print("Go runs on ")
	//  switch内で変数を定義
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}


	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}


	// defer: returnするまで遅延させる
	// 複数ある場合、returnにしてからLIFO(last in first out)の順で実行される
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		// fmt.Println(i)
	}
	fmt.Println("done")
}
