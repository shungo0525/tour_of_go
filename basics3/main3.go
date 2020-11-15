package main

import (
	"fmt"
	"strings"
)

// Person は人間を表す構造体。
type Person struct {
    Name string
    Age  int
}

type Vertex struct {
	X int
	Y int
}

// 関数はクロージャーである
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func main() {

	// map: rubyでいうhash
	// p := map[string]string{
	// 		"Name": "太郎",
	// 		"Age":  "20",
	// }
	// fmt.Println(p)

	// p := Person{
	// 	Name: "太郎",
	// 	Age:  20,
	// }
	// fmt.Println(p)

	// p2 := &p
	// p2.Name = "次郎"
	// p2.Age = 19
	// fmt.Println(p)

	var p *int // ポインタは，アドレスを格納する変数
	// & を頭に付けると、新しく割り当てられたstructへのポインタを戻します。
	i := 10
	p = &i
	fmt.Println(" p is ", p)
	fmt.Println("*p is ", *p)
	fmt.Println("&p is ", &p)


	// fmt.Println(Vertex{X: 1, Y: 2})
	fmt.Println(Vertex{1, 2})

	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v)

	// 配列の定義
	var a[2]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	// primes[4]を除く -> [3, 5, 7]
	var s = primes[1:4]
	fmt.Println(s)

	// スライスの要素を変更すると元の配列も変更される　
	s[0] = 1
	fmt.Println(primes) // -> [2 1 5 7 11 13]

	fmt.Println(len(primes))
	fmt.Println(cap(primes))
	fmt.Println(primes)

	// 配列を作成
	slice := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf"}
	fmt.Println(slice)
	// ","で区切って文字列結合する
	fmt.Println(strings.Join(slice, ","))

	var b []int
	fmt.Println(b)

	// 配列に要素と追加
	b = append(b, 0)
	fmt.Println(b)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		// %d: 10進数
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for v := range pow {
		fmt.Println(v)
	}


	// mapを生成  make(map[keyの型]valueの型)
	// make 関数は指定された型のマップを初期化して、使用可能な状態で返す。初期値は入れられない。
	m := make(map[string]string)
	m["test1"] = "test1"
	m["test2"] = "test2"
	m["test3"] = "test3"
	fmt.Println(m)
	fmt.Println(m["test1"])

	l := map[string]string{"go":"golang", "rb":"ruby", "js":"javascript"}
	fmt.Println(l)

	m2 := make(map[string]int)
	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])

	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}
