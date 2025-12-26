package main


import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

	currentLineContents := ""

	for {
		buffer := make([]byte, 8) // for first 8 bytes of the file
		n, err := f.Read(buffer)
		if err != nil {
			// if end of line then break the infinite loop
			if currentLineContents != ""{
				fmt.Printf("read: %s\n", currentLineContents)
				currentLineContents = ""
			}
			if errors.Is(err, io.EOF){
				break
			}
			fmt.Printf("error: %s", err.Error())
			break
		}
		str := string(buffer[:n])
		parts := strings.Split(str, "\n")

		for i := 0; i < len(parts)-1; i++{
			fmt.Printf("read: %s%s\n", currentLineContents, parts[i])
			currentLineContents = ""
		}
		currentLineContents += parts[len(parts)-1]
		
	}

}
