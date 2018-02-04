package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type q struct {
	question string
	answer   string
}

func main() {

	questions := map[int]q{
		1: q{question: "q0?", answer: `answer0`},
		2: q{question: "q1?", answer: "answer1"},
		3: q{question: "q2?", answer: "answer2"},
	}

	for {
		rand.Seed(time.Now().UTC().UnixNano())
		randNumber := rand.Intn(len(questions)) + 1

		reader := bufio.NewReader(os.Stdin)
		fmt.Println(questions[randNumber].question)

		_, _ = reader.ReadString('\n')
		fmt.Println(questions[randNumber].answer)
	}
}

