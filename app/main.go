package main

import (
	"github/timly278/english-helper/clock"
	"github/timly278/english-helper/session"
	"fmt"
	"os"
)

func main() {
	clock.MakeSound("./resource/sound1.mp3")
	conf, err := session.LoadConfig(".")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Printf("duration: %f\n", conf.Story.Duration.Seconds())

	if conf.Mode == session.MODE_MANUAL {
		fmt.Println("not implement")
	} else {
		fmt.Println("Your today's lesson is:")
		// conf.RandomSession()
		session.DoStory(conf)
	}
}

//TODO: write code to validate if the configuration value in .yaml file are valid or not.
//      if not, throw error right on the load config stage.

//TODO: consider improve the app to an CLI application using Go-Cobra
