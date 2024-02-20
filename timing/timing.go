package timing

import (
	"fmt"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v ms to complete.\n", name, convertTimeDurationToMiliseconds(time.Since(start)))
	}
}

func convertTimeDurationToMiliseconds(duration time.Duration) float64 {
	nanoseconds := float64(duration.Nanoseconds())
	milliseconds := nanoseconds / float64(time.Millisecond)

	return (milliseconds)
}
