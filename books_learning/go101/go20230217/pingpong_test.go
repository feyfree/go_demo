package go20230217

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball := <-table // 接球
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue { // 溢出结束
			os.Exit(0)
		}
		lastValue = ball
		table <- ball // 回球
		time.Sleep(time.Second)
	}
}

func TestPingpong(t *testing.T) {
	table := make(chan Ball)
	go func() {
		table <- 1 // （裁判）发球
	}()
	go Play("A:", table)
	Play("B:", table)
}
