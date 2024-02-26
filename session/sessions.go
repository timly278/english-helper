package session

import (
	"fmt"
	"github/timly278/english-helper/clock"
	"github/timly278/english-helper/service"
	"math/rand"
	"strings"
)

var session = []func(*Config){
	DoStory,
	DoAnswerQuestion,
	DoVocabSentence,
	DoDescribeImage,
	DoAlphabet,
	DoDailyJournal,
}

func (c *Config) RandomSession() {
	session[service.RandomInt(0, len(session))](c)
}

func DoStory(c *Config) {
	fmt.Printf("\n************************************************************\n\n")
	fmt.Println("\t\tWelcome to Story session")
	fmt.Printf("\n************************************************************\n\n")

	data := service.ReadFile(c.Story.Path)
	topic, err := service.GetRandomDash(data, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	hints := strings.Split(topic[0], "    + ")

	fmt.Println(hints[0])
	for _, v := range hints[1:] {
		fmt.Printf("+ %s", v)
	}
	fmt.Printf("\n************************************************************\n")
	AskForReady()
	isNatural := clock.StartTimer(c.Story.Duration)
	if isNatural {
		clock.Sound(c.SoundPath)
	}
}

func DoAnswerQuestion(c *Config) {
	data := service.ReadFile(c.AnswerQuestion.Path)
	fmt.Printf("\n************************************************************\n\n")
	fmt.Printf("You have %d questions today:\n\n", c.AnswerQuestion.NumOfQuestions)

	questions, err := service.GetRandomDash(data, c.AnswerQuestion.NumOfQuestions)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < c.AnswerQuestion.NumOfQuestions; i++ {
		fmt.Printf("%d - %s\n\n", i+1, questions[i])
	}
	fmt.Printf("************************************************************\n")
	AskForReady()
	clock.StartTimer(c.AnswerQuestion.Duration)
	clock.Sound(c.SoundPath)
}

func DoVocabSentence(c *Config) {
	data := service.ReadFile(c.VocabSentence.Path)
	questions, err := service.GetRandomDash(data, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n************************************************************\n\n")
	fmt.Println("\t\tWelcome to vocab sentence session")
	fmt.Printf("\n************************************************************\n")

	AskForReady()
	words := strings.Split(questions[0], ", ")
	for _, w := range words {
		fmt.Printf("%s\n", w)
		clock.StartTimer(c.VocabSentence.Duration)
		fmt.Printf("\r--------------- next -------------\n")
	}
	fmt.Println("Finish!")
	clock.Sound(c.SoundPath)
}

// DoAlphabet shows a random letter and user would make a sentence with it
// within a short specified time
func DoAlphabet(c *Config) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	perm := rand.Perm(len(alphabet))
	fmt.Printf("\n************************************************************\n\n")
	fmt.Println("\tWelcome to Alphabet sentence session")
	fmt.Printf("\n************************************************************\n")
	AskForReady()
	for i, v := range perm {
		fmt.Printf("letter[%d] is: %s\n", i, string(alphabet[v]))
		clock.StartTimer(c.VocabSentence.Duration)
		fmt.Printf("\r--------------- next -------------\n")
	}
	fmt.Println("Finish!")
	clock.Sound(c.SoundPath)
}

func DoDailyJournal(c *Config) {

}

func AskForReady() {
	var key string
	for {
		fmt.Printf("Are you ready? y/n: ")
		fmt.Scan(&key)
		if key == "y" || key == "Y" {
			break
		}
	}
}
