package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// question struct stores a single question and its corresponding answer.
type question struct {
	q, a string
}

type score int

// check handles a potential error.
// It stops execution of the program ("panics") if an error has happened.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// questions reads in questions and corresponding answers from a CSV file into a slice of question structs.
func questions() []question {
	f, err := os.Open("quiz-questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []question
	for _, row := range table {
		questions = append(questions, question{q: row[0], a: row[1]})
	}
	return questions
}

// ask asks a question and returns an updated score depending on the answer.
func ask(s score, question question, c chan score) {
	//for {
	//q := <-question
	fmt.Println(question.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	if strings.Compare(text, question.a) == 0 {
		fmt.Println("Correct!")
		s++
		c <- s
	} else {
		fmt.Println("Incorrect :-(")
		c <- s
	}
	//c <- s
	//}
}

// This also works
//func ask(s score, question chan question, c chan score) {
//	for {
//		q := <-question
//		fmt.Println(q.q)
//		scanner := bufio.NewScanner(os.Stdin)
//		fmt.Print("Enter answer: ")
//		scanner.Scan()
//		text := scanner.Text()
//		if strings.Compare(text, q.a) == 0 {
//			fmt.Println("Correct!")
//			s++
//			c <- s
//		} else {
//			fmt.Println("Incorrect :-(")
//			c <- s
//		}
//	}
//}

// The new main function
//func main() {
//	s := score(0)
//	qs := questions()
//	qc := make(chan question)
//	c1 := make(chan score)
//	go ask(s, qc, c1)
//	for _, q := range qs {
//		qc <- q
//		select {
//		case <-time.After(5 * time.Second):
//			fmt.Println("")
//			fmt.Println("Final score", s)
//			return
//		case s1 := <-c1:
//			s = s1
//		}
//	}
//	fmt.Println("Final score", s)
//}

func main() {
	s := score(0)
	qs := questions()
	c1 := make(chan score)
	for _, q := range qs {
		go ask(s, q, c1)
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("")
			fmt.Println("Final score", s)
			return
		case s1 := <-c1:
			s = s1
		}
	}
	fmt.Println("Final score", s)
}
