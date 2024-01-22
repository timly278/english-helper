package session

import (
	"github/timly278/english-helper/service"
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
	data := service.ReadFile(c.Story.Path)
	topic := service.GetRandomDash(string(data))
	service.PrintOut(topic)
}

func DoAnswerQuestion(c *Config) {

}

func DoVocabSentence(c *Config) {

}

func DoDescribeImage(c *Config) {

}

// DoAlphabet shows a random letter and user would make a sentence with it
// within a short specified time
func DoAlphabet(c *Config) {

}

func DoDailyJournal(c *Config) {

}
