package main

import (
	"fmt"
	"os"
)

// range使えばもっときれいに書ける

func main() {
	var s, sep, result string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		result = fmt.Sprintf("%d : %s", i, os.Args[i])
		fmt.Println(result)
	}
}
