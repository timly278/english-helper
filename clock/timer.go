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

func StartTimer(duration time.Duration) bool {

	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				_ = t
				fmt.Printf("\r%d:%d ", int(duration.Minutes()), int(duration.Seconds())%60)
				duration = duration - time.Second
			}
		}
	}()

	time.Sleep(duration)
	ticker.Stop()
	done <- true
	return true
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
