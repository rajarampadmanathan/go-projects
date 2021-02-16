package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myWriter struct{}

func main() {
	fmt.Println("Hello")
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// response := make([]byte, 99999)
	// resp.Body.Read(response)
	// fmt.Println(string(response))
	x, _ := io.Copy(myWriter{}, resp.Body)
	fmt.Println("Bytes written:", x)

}

func (m myWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
