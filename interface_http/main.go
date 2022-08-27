package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWritter{}

	// // bs := []byte{}
	// bs := make([]byte, 99999) // create byte slice with 99999 element inside of it
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

func (logWritter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
