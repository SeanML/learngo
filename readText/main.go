package main

import (
	"fmt"
	"io"
	"os"
)

// Method 1
// func main() {
// 	b, err := ioutil.ReadFile(os.Args[1])

// 	if err != nil {
// 		fmt.Println("Error encountered:", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(string(b))
// }

// Method 2
func main() {
	b, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, b)
}
