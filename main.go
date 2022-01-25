package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// open csv file
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// close file at the end of program
	defer file.Close()

	// read csv values
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// create list of QnAs
	QnAs := createQnAs(data)
	// get answer sheet from QnA list
	answerSheet := createAnswerSheet(QnAs)
	correctCount := 0

	// timer goroutine
	timeLimit := flag.Int("limit", 4, "time limit for quiz in seconds")
	flag.Parse()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//  get user input
	for index, QnA := range QnAs {
		fmt.Printf("%s: ", QnA.Question)
		answerChannel := make(chan string)
		go func() {
			var input string
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			answerChannel <- input
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %v out of %v", correctCount, len(answerSheet))
			return
		case ans := <-answerChannel:
			if ans == answerSheet[index] {
				correctCount += 1
			}
		}

	}
	correctString := fmt.Sprintf("You scored %v out of %v", correctCount, len(answerSheet))
	fmt.Println(correctString)
}

type QnA struct {
	Question string
	Answer   string
}

// return qna list from csv file
func createQnAs(data [][]string) []QnA {
	QnAList := []QnA{}

	for _, line := range data {
		singleQnA := QnA{
			line[0],
			line[1],
		}
		QnAList = append(QnAList, singleQnA)
	}
	return QnAList
}

func createAnswerSheet(data []QnA) []string {
	var answers []string
	for _, qna := range data {
		answers = append(answers, qna.Answer)
	}
	return answers
}
