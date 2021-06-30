package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var defaultCSV = "problems.csv"

var csvFlag = flag.String("csv", defaultCSV, "a csv file in the format of 'question,answer' (default \"problems.csv\")")
var limitFlag = flag.Int("limit", 30, "the time limit for the quiz in seconds (default 30")

func init() {
	flag.Parse()
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read file " + filePath + "\n")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV file ", err)
	}
	return records
}

func parseCSVLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

// we pass in problem struct into play, and we don't care about the initial format
// maybe it came from csv or json, or other format
type problem struct {
	question string
	answer   string
}

func end_game(result int) {
	fmt.Printf("\nRan out of time, result: %v\n", result)
	os.Exit(0)
}

func play(problems *[]problem) int {
	var result int
	problems_len := len(*problems)
	timer := time.NewTicker(time.Duration(*limitFlag) * time.Second)

	answerCh := make(chan string)

	for i, problem := range *problems {
		fmt.Printf("Problem #%v/%v, %v: ", i+1, problems_len, problem.question)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			end_game(result)
		case answer := <-answerCh:
			if answer == problem.answer {
				result += 1
			}
		}
	}

	for i, problem := range *problems {
		fmt.Printf("Problem #%v/%v, %v: ", i+1, problems_len, problem.question)
		var answer string
		fmt.Scanln(&answer)
		if answer == problem.answer {
			result += 1
		}
	}
	return result
}

// packages to use:
// flags, csv, os, time
func main() {
	problems := parseCSVLines(readCsvFile(*csvFlag))
	result := play(&problems)
	fmt.Printf("Result: %v/%v\n", result, len(problems))
}
