package main

import (
	"fmt"
	"unsafe"
)

// Hoge is struct
type Hoge struct {
	A []int
	B int
}

func main() {
	/*
		structの初期化
		new(T)は、T型のフィールド要素を"ゼロ値"で初期化し、*T(T型へのポインタ)をreturnする。
		なので、x2とx3は同じ。違いがあるとすれば、1行で初期化できるか否か。
		structは、32bitか64bitでメモリの取り方が異なり、64bit環境の場合、各フィールドは8byte区切りでメモリを確保する。
	*/
	var x1 Hoge
	fmt.Printf("%#v\n", x1)        // => main.Hoge{A:[]int(nil), B:0}
	fmt.Println(unsafe.Sizeof(x1)) // => 32 ( = 24byte + 8byte )

	x2 := new(Hoge)
	fmt.Printf("%#v\n", x2)        // => &main.Hoge{A:[]int(nil), B:0}
	fmt.Println(unsafe.Sizeof(x2)) // => 8 (64bit環境でポインタのサイズは8byte)

	x3 := &Hoge{}
	fmt.Printf("%#v\n", x3)        // => &main.Hoge{A:[]int(nil), B:0
	fmt.Println(unsafe.Sizeof(x3)) // => 8 (64bit環境でポインタのサイズは8byte)

	/*
		sliceの初期化
		map, slice, channelは内部にデータ構造を持つため、makeで初期化する。
		map, slice, channelのゼロ値はnilなので、以下の出力となる。
	*/
	y1 := new([]int)
	fmt.Printf("%#v\n", y1) // => &[]int(nil)

	y2 := make([]int, 0)
	fmt.Printf("%#v\n", y2) // => []int{}

	/*
		まとめ
		newはゼロ値で初期化してポインタを返す。
		map, slice, channelは、内部にデータ構造を持ち、そのポインタ型である。また、そのフィールドは初期化が必要である。
		そのため、newではなくmakeで初期化を行う。
	*/
}
