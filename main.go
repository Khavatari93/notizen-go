package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Notes struct {
	ID   int
	Note string
}

func test() {
	f, err := os.Open("database.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}

func createFile() {

	if _, err := os.Stat("database.csv"); err == nil {
		return
	} else {
		fmt.Println("Creating note database...")
		file, err := os.Create("database.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		file2, err := os.OpenFile("database.csv", os.O_RDWR, 0755)
		if err != nil {
			log.Fatal(err)
		}

		table := [][]string{
			{"ID", "Note"},
			{"1", "Your first Note!"},
		}

		csvWriter := csv.NewWriter(file2)
		csvWriter.WriteAll(table)

		if err := csvWriter.Error(); err != nil {
			log.Fatalln("error writing csv:", err)
		}

		return
	}

}

func main() {
	createFile()
	test()
}
