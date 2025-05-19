package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	f, err := os.Open("hum.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Loop infinitely
	loop := beep.Loop(-1, streamer)

	speaker.Play(beep.Seq(loop))

	// Read stdin to wait for preceding program to complete
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

}
