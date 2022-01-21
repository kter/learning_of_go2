package main

import (
	"fmt"
)

// [2]intへのエイリアスIntPair
type IntPair [2]int

func (ip IntPair) First() int {
	return ip[0]
}

func (ip IntPair) Last() int {
	return ip[1]
}

type Strings []string

// 文字列のスライスを区切り文字を連結するメソッド
func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}

func main() {

	ip := IntPair{1, 2}
	// 1
	fmt.Println(ip.First())
	// 2
	fmt.Println(ip.Last())
	// A,B,C
	fmt.Println(Strings{"A", "B", "C"}.Join(","))
}
