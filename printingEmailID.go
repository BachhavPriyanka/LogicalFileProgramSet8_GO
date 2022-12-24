package main

import (
	"bufio"
	"fmt"
	"os"
)

func printEmail(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	Scanner1 := bufio.NewScanner(file)

	Scanner1.Split(bufio.ScanWords)
	for Scanner1.Scan() {

		fmt.Println(Scanner1.Text())
		
	}

	

}
func printLineByLine(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Printing Line by Line :", scanner.Text())
	}

	defer file.Close()
}
func main() {
	fileName := "LearnerData.txt"

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	printLineByLine(fileName)
	printEmail(fileName)
}
