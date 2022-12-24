package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonAppleRecipie := `{
	"Thinly Sliced Peeled Apples" : "6 Cups",
	"sugar" : "3/4 cup",
	"flour" : "2 tablespoons",
	"cinamon" : "1/4 teaspoon",
	"nutmeg" : "1/8 tablesppon",
	"lemon juice" : "1 tablespoon"}`

	store := map[string]string{}

	err := json.Unmarshal([]byte(jsonAppleRecipie), &store)
	fmt.Println(store)

	file, err := os.Create("recipie.csv")

	defer file.Close()
	if err != nil {
		panic(err)
	}
	for key, value := range store {
		_, err = file.WriteString(fmt.Sprintf("%s : %s\n", key, value))
		if err != nil {
			panic(err)
			return
		}
	}
}
