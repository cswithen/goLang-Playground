package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// setting a custom type
type logWriter struct{}

func main() {
	// http docs (https://pkg.go.dev/net/http@go1.17) get request (https://pkg.go.dev/net/http@go1.17#Get)
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//defining a new instance of the custom type
	lw := logWriter{}

	// METHOD 01 : BASIC
	// give me an empty byteslice with the possibility of 99,999 spaces
	// bs := make([]byte, 99999)

	// use the read function that the Body portion has access to, and read data to the byteslice (bs)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// METHOD 02 : ADVANCED (https://pkg.go.dev/io@go1.17#Copy)
	// io.Copy(os.Stdout, resp.Body)

	// using a custom type to manage our Write
	io.Copy(lw, resp.Body)
}

//setting a function for our custom type
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
