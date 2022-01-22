package main

import (
	"fmt"
	"sync"
	"time"
)

var st struct{ A, B, C int }

// ミューテックスを保持するパッケージ変数
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// LOCK
	mutex.Lock()

	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)

	// mutexを使わないとst.A, st.B, st.Cの値が揃わない
	fmt.Println(st.A, st.B, st.C)

	// UNLOCK
	mutex.Unlock()
}

func main() {
	// generate mutex
	mutex = new(sync.Mutex)

	// generate sync.WaitGroup
	wg := new(sync.WaitGroup)
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("AAAAAAAAAAAAAAA")
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
			wg.Done()
		}()
	}
	// これがないとゴルーチンの処理が始まる前に処理が終わってしまう
	wg.Wait()
}
