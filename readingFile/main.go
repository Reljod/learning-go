package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from the future Go expert, Reljod!"
	file, err := os.Create("./outputFile.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Wrote a file with %v characters\n", length)

	defer file.Close()
	defer readFile("./outputFile.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from data: ", string(data))
}
