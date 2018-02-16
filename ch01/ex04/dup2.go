// 途中
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// キーが文字列、値がintの配列を生成
	counts := make(map[string]int)
	fileName := os.Args[1:]
	// 標準入力の引数からfilenameを取得
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		// dataをSplitしてlineの数を数える
		for _, line := range strings.Split(string(data), "\n") {
			// ↓の書き方がなにをしてるかいまいちわからない
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileName)
		}
	}
}
