package main

import (
	"fmt"
	"os"
	"strconv"
)

const defaultSize = 15

func main() {

	size := defaultSize

	args := os.Args[1:]
	if len(args) >= 1 {
		arg := args[0]
		s, err := strconv.ParseInt(arg, 10, 0)
		if err != nil {
			panic(err)
		}

		size = int(s)
	}

	generator := NewASCII()
	pass, err := generator.Generate(size)
	if err != nil {
		panic(err)
	}

	fmt.Println(pass)
}
