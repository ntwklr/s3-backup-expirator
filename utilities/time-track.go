package utilities

import (
	"fmt"
	"time"
)

// TimeTrack return the elapsed time since start
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)

	fmt.Printf("%s took %v\n", name, elapsed)
}
