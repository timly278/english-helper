package clock

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func StartTimer(duration time.Duration) {
	Enter := make(chan bool, 1)
	tick := time.NewTicker(time.Second)
	isRunning := true
	var b []byte = make([]byte, 1)

	go func() {
		for {
			os.Stdin.Read(b)
			switch b[0] {
			case '1':
				Enter <- true
				return

			case '2', ' ':
				if isRunning {
					tick.Stop()
					isRunning = false
				} else {
					tick.Reset(time.Second)
					isRunning = true
				}
			}
		}
	}()
	for duration >= 0 {
		select {
		case isEnter := <-Enter:
			if isEnter {

				return
			}
		case <-tick.C:
			fmt.Printf("\r%d:%02d ", int(duration.Minutes()), int(duration.Seconds())%60)
			duration -= time.Second
		}
	}

	fmt.Printf("\r0:00\n")
}

func Sound(path string) {
	MakeSound(path)
}

func MakeSound(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("cannot open sound file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
