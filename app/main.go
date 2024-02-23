package main

import (
	"fmt"
	"github/timly278/english-helper/session"
	"os"
)

func main() {
	conf, err := session.LoadConfig(".")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	if conf.Mode == session.MODE_MANUAL {
		fmt.Println("not implement")
	} else {
		fmt.Println("Your today's lesson is:")
		// conf.RandomSession()
		session.DoAnswerQuestion(conf)
	}
}

//TODO: write code to validate if the configuration value in .yaml file are valid or not.
//      if not, throw error right on the load config stage.

//TODO: consider improve the app to an CLI application using Go-Cobra
