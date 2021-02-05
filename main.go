package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var src = strings.NewReader("Insert text here\n")
	var dst = os.Stdout            //os.Stdin -> read, os.Stdout -> write
	var n, err = io.Copy(dst, src) //copy from src(Writer) to dst(Reader)

	if err != nil {
		panic(err)
	}

	fmt.Println("Copied elements:", n)

}
