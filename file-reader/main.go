package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// this will print the second argument that we are providing to our 'go run main.go myfile.txt'
	// fmt.Println(os.Args[1])

	// os.Open docs (https://pkg.go.dev/os@go1.17#Open)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}

// secondary if you wanted to print out the main.go file as well, first build the file 'go build main.go' then run the executable created in the directory with a call to the original file './main main.go'
