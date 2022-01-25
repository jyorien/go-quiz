package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
	QnAs := createQnAs(data)
	answerSheet := createAnswerSheet(QnAs)
	var givenAnswers []string
	for index, QnA := range QnAs {
		str := fmt.Sprintf("%s: ", QnA.Question)
		fmt.Print(str)
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		givenAnswers = append(givenAnswers, input)
		QnAs[index] = QnA
	}

	correctCount := 0
	for index, answer := range givenAnswers {
		correct := answer == answerSheet[index]
		if correct {
			correctCount += 1
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
