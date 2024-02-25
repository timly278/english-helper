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
	// Start the countdown timer
	for duration >= 0 {
		fmt.Printf("\r%d:%02d ", int(duration.Minutes()), int(duration.Seconds())%60)
		select {
		case <-time.After(time.Second):
			duration -= time.Second
		case <-time.After(time.Millisecond * 100):
			// Check for user input asynchronously every 100 milliseconds
			select {
			case key := <-getKey():
				switch key {
				case '\r': // Enter key
					return // Stop the timer
				case ' ': // Space key
					fmt.Print("\rTimer paused. Press Space again to resume.")
					<-getKey() // Wait for another key press
				}
			default:
				// No input, continue the timer
			}
		}
	}

	// Final message when timer ends
	fmt.Printf("\r0:00\n")
}

func getKey() <-chan rune {
	ch := make(chan rune)
	go func() {
		var b [1]byte
		os.Stdin.Read(b[:])
		ch <- rune(b[0])
	}()
	return ch
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
