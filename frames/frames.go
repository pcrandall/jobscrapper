package frames

import (
	"fmt"
	"time"
)

type FrameType struct {
	GetFrame  func(int) string
	GetLength func() int
}

// Create a function that returns the next frame, based on length
func DefaultGetFrame(frames []string) func(int) string {
	return func(i int) string {
		return frames[i%(len(frames)-1)]
	}
}

// Create a function that returns frame length
func DefaultGetLength(frames []string) func() int {
	return func() int {
		return len(frames)
	}
}

// Given frames, create a FrameType with those frames
func DefaultFrameType(frames []string) FrameType {
	return FrameType{
		GetFrame:  DefaultGetFrame(frames),
		GetLength: DefaultGetLength(frames),
	}
}

var FrameMap = map[string]FrameType{
	"forrest": Forrest,
	"parrot":  Parrot,
	"clock":   Clock,
	"nyan":    Nyan,
	"rick":    Rick,
}

func Start(c chan bool) {

	frames, ok := FrameMap["rick"]

	if !ok {
		fmt.Printf("frames= %+v\n", frames)
	}

	i := 0

	for {
		// fmt.Printf("frames= %+v\n", frames)
		select {

		case <-c: // Stop frames
			return

		default:

			if i >= frames.GetLength() {
				i = 0
			}

			time.Sleep(time.Millisecond * 70) // Refresh rate

			// Clear screen
			clearScreen := "\033[2J\033[H"
			newLine := "\n"

			frame := clearScreen + frames.GetFrame(i) + newLine

			fmt.Println(frame) // Write frames
			i++
		}
	}
}
