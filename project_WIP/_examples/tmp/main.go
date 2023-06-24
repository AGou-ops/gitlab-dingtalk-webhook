package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("output.log")
	if err != nil {
		fmt.Println("failed to create output.log", err)
		return
	}
	defer file.Close()

	data := []byte("hello world")
	file.Write(data)

	writer := io.MultiWriter(file)
	writer.Write([]byte("hello world"))

	file.Seek(0, 0)
	buffer := make([]byte, 9999999)
	file.Read(buffer)
	fmt.Println(string(buffer))
}
