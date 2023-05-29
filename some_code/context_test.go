package some_code

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	messages := make(chan int, 10)

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// consumer
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}(ctx)

	defer close(messages)
	defer cancel()

	select {
	case <-ctx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}
}

func TestWithCancel(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		cancelFunc()
	}()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Done")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
