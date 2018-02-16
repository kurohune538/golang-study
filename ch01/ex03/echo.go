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
	start := time.Now()
	for i := 1; i < num; i++ {
		s += "s"
	}
	end := time.Now()
	fmt.Println(fmt.Sprintf("%.5f 秒", end.Sub(start).Seconds()))

	start2 := time.Now()
	ss := []string{}
	for i := 1; i < num; i++ {
		strings.Join(ss, "s")
	}
	end2 := time.Now()
	fmt.Println(fmt.Sprintf("%.5f 秒", end2.Sub(start2).Seconds()))
}
