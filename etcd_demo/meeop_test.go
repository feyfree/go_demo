package etcd_demo

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestReadFile(t *testing.T) {
	dir, _ := os.Getwd()

	filepath := dir + "/aaa.txt"
	fmt.Println(filepath)

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second * 5)
			file, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			writer := bufio.NewWriter(file)
			writer.WriteString("a")
			writer.Flush()
			addFile(filepath, "aaa")
			file.Close()
		}
	}()

	go func() {
		listen("aaa")
	}()

	select {}
}
