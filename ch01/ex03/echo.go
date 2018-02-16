// 何故か表示されない
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	num := 1000
	s := "s"
	// インクリメントでの時間計測
	start := time.Now()
	for i := 1; i < num; i++ {
		s += "s"
	}
	end := time.Now()
	// Subで引き算した結果を表示
	fmt.Println(fmt.Sprintf("%.5f 秒", end.Sub(start).Seconds()))

	// Joinの時間計測
	start2 := time.Now()
	ss := []string{}
	for i := 1; i < num; i++ {
		// 文字列を追加
		strings.Join(ss, "s")
	}
	end2 := time.Now()
	// Subで引き算した結果を表示
	fmt.Println(fmt.Sprintf("%.5f 秒", end2.Sub(start2).Seconds()))
}
