package main

import (
	old_fmt "fmt"

	"github.com/duan-go-functional/f_fmt"
	"github.com/duan-go-functional/fmt"
	f "github.com/duan-go-functional/go_functional"
)

func main() {
	tmp := f.Block(
		"hello world!",
		f_fmt.Println(nil),
		f_fmt.Println(nil),
		f_fmt.Println(nil),
		f_fmt.Println(nil),
		// f_fmt.Println(nil),
	)

	tmp()
	old_fmt.Println("-----------------")
	old_fmt.Println("res tmp:", tmp())
	old_fmt.Println("-----------------")

	f.Block(
		fmt.Println(nil, "new fmt:", tmp()),
	)()
}
