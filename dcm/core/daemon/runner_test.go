package daemon

import (
	"fmt"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	t.Run("test daemon run", func(t *testing.T) {
		msg := make(chan string)
		testLimit := 0
		limitCount := 4
		for {
			go Runner(msg)
			PrintMemUsage()
			time.Sleep(time.Second * 2)

			msgReceiver := <-msg
			fmt.Println(msgReceiver)

			if testLimit == limitCount {
				break
			}

			testLimit++
		}
	})
}
