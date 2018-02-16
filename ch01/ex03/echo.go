package main

import (
	"fmt"
	"strings"
	"time"
)


func main() {
	var num int = 10
	var s string = "s"
	start := time.Now()
	for i := 1; i < num; i++ {
		s += s
	}
	end := time.Since(start).Seconds()
	fmt.Print(end)

	start2 := time.Now()
	ss := []string{} 
	for i := 1; i < num; i++ {
		strings.Join(ss,"s")
	}
	end2 := time.Since(start2).Seconds()
	fmt.Sprintf("%d秒", end2)
}