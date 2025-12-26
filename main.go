package main


import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open(inputFilePath)

	if err != nil {
		log.Fatalf("Could not print the file %s: %s", inputFilePath, err)
	}
	defer f.Close()

	fmt.Printf("Reading data from file %s", inputFilePath)
	fmt.Println("====================================================")

	for {
		b := make([]byte, 8) // for first 8 bytes of the file
		n, err := f.Read(b)
		if err != nil {
			// if end of line then break the infinite loop
			if errors.Is(err, io.EOF){
				break
			}
			fmt.Printf("error: %s", err.Error())
			break

		}
		str := string(b[:n]) // only print the first n bytes to avoid garbage characters
		fmt.Printf("read: %s \n", str)
	}

}
