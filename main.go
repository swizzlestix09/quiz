package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Questions struct {
	Question string
	Answer   string
}

func createQuizList(data [][]string) {
	fmt.Printf("%+v\n", data)
	//each line in its own array - 0 is q 1 is ans
	//returning
	return
}
func main() {
	file, err := os.Open("problems.csv")
	//read csv
	if err != nil {
		log.Fatal("error opening csv:", err)
	}
	defer file.Close()
	//store in a list
	fileReader := csv.NewReader(file)
	data, err := fileReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// questionsList := createQuizList(data)
	createQuizList(data)

	// fmt.Printf("%+v\n", questionsList)
}
