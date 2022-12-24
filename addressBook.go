package main

import (
	"encoding/csv"
	"os"
)

type Friend struct {
	Name            string
	MobNumber       string
	AlternateMobNum string
	Address         string
	City            string
}

func main() {
	records := [][]string{
		{"Name", "MobileNumber", "AlternateMobNum", "Address", "City"},
		{"Priyanka", "9850503365", "8669335732", "Ojhar", "Nahik"},
		{"Gaurav", "888776587", "022543200", "nizamuddin", "Delhi"},
		{"Harry", "9089898989", "7978787878", "Bunglow no.34", "Bangalore"},
		{"Raj", "9009009009", "88888888", "Chennai Near XYZ ", "Hitech city"},
		{"Jenny", "8282828282", "022543200", "Thaane", "Mumbai"},
		{"Diksha", "8282828282", "022543200", "Lonawala", "Pune"},
		{"Malika", "999977777", "998837342", "ojhar", "Nasik"},
		{"Gaurav", "888776587", "022543200", "nizamuddin", "Delhi"},
	}
	file, err := os.Create("friends.csv")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	writeToFile := csv.NewWriter(file)
	err = writeToFile.WriteAll(records)
	if err != nil {
		panic(err)
	}

}
