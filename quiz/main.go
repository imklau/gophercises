package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName, timeLimit := readArguments()

	records := readCsvFile(&fileName)

	startQuiz(records, &timeLimit)
}

func readArguments() (string, int) {
	fileName := flag.String("csv", "problems.csv", "a path to csv file")
	timeLimit := flag.Int("timer", 10, "the number of seconds to run the quiz")

	flag.Parse()

	return *fileName, *timeLimit
}

func readCsvFile(file *string) [][]string {
	csvFile, err := os.Open(*file)

	if err != nil {
		log.Fatal("Unable to open the file: "+*file, err)
	}

	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse the CSV file.", err)
	}

	return records
}

func startQuiz(records [][]string, timeLimit *int) {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0
	userAnswer := ""

	for _, value := range records {
		answerChan := make(chan string)

		item := problem{
			question: value[0],
			answer:   strings.TrimSpace(value[1]),
		}

		fmt.Print(item.question, " ")

		go func() {
			fmt.Scanf("%s\n", &userAnswer)
			answerChan <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTime is over!\nYour score is: %d of %d.\n", score, len(records))
			return
		case answer := <-answerChan:
			if answer == item.answer {
				score += 1
			}
		}
	}

	fmt.Printf("Your score is: %d of %d.\n", score, len(records))
}
