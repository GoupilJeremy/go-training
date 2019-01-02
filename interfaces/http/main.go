package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	written, err := io.Copy(lw, resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("number of byte written:", written)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
