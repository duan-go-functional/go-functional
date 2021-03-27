package main

import (
	"github.com/duan-go-functional/fmt"
	f "github.com/duan-go-functional/go_functional"
)

func main() {
	tmp := f.Block(
		fmt.Println(nil, "hello world!"),
		fmt.Println(nil, "hello world!"),
		fmt.Println(nil, "hello world!"),
	)
	tmp()
}
