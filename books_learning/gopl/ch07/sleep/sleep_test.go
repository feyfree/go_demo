package sleep

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

//!+sleep
var period = flag.Duration("period", 1*time.Second, "sleep period")

func TestSleep(t *testing.T) {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
