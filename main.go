package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	count := 0
	for _, s := range os.Args[1:] {
		count += len(strings.Fields(s))
	}
	fmt.Println(count)
}
