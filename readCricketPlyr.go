package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Player struct {
	Name, Period             string
	MatchesPlayed, RunScored int
	AvgScore                 float64
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func readCSV(fileName string) {
	fHandler, err := os.Open(fileName)
	check(err)
	defer fHandler.Close()

	reader := csv.NewReader(fHandler)

	records, err := reader.ReadAll()
	check(err)

	//fmt.Println(records) // displaying complete records in csv file in 2D array

	for _, record := range records {
		fmt.Println("\n", record)
	}
	//Printing without Header Row
	fmt.Println("--DATA WITHOUT HEADER ROW--")
	var data []Player
	for i, line := range records {
		if i > 0 {
			var r Player
			for j, field := range line {
				if j == 0 {
					r.Name = field
				} else if j == 1 {
					r.Period = field
				} else if j == 2 {
					fieldString, _ := strconv.Atoi(field)
					r.MatchesPlayed = fieldString
				} else if j == 3 {
					fieldString, _ := strconv.Atoi(field)
					r.RunScored = fieldString
				} else if j == 4 {
					fieldString, _ := strconv.ParseFloat(field, 64)
					r.AvgScore = fieldString
				}
			}
			data = append(data, r)
		}
	}
	// fmt.Println(data)
	for _, record := range data {
		fmt.Println("\n", record)
	}
}
func main() {
	readCSV("Cricket_Players_Stats.csv")
}
