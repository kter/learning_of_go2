package main

import (
	"fmt"
)

type Point struct{ X, Y int }

func (p Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

type Point1 struct{ X, Y int }

func (p *Point1) Set(x, y int) {
	p.X = x
	p.Y = y
}

func main() {
	/*
		下記のようにレシーバーが値型だと値が変わらない。
		レシーバーのコピーが走り呼び出し側と内部で実態が違うものになるため
	*/
	p1 := Point{}
	p1.Set(1, 2)
	fmt.Println(p1.X)
	fmt.Println(p1.Y)

	p2 := &Point{}
	p2.Set(1, 2)
	fmt.Println(p2.X)
	fmt.Println(p2.Y)

	/*
		下記のようにレシーバーがポインタ型だ良い
	*/
	p3 := Point1{}
	p3.Set(1, 2)
	fmt.Println(p3.X)
	fmt.Println(p3.Y)

	p4 := &Point1{}
	p4.Set(1, 2)
	fmt.Println(p4.X)
	fmt.Println(p4.Y)
}
