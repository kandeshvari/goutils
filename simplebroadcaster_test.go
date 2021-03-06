package goutils

import (
	"sync"
	"testing"
)

var (
	publishStrs = []string{"one", "two", "three", "four", "five", "donePublish"}
)

func EqualString(t *testing.T, wait, got string) {
	if got != wait {
		t.Fatalf("unexpected string! wait: %s, got: %s", wait, got)
	}
}

func TestNewSimpleBroadcaster(t *testing.T) {
	wg := sync.WaitGroup{}
	b := NewSimpleBroadcaster(10)

	for range [3]struct{}{} {
		c, num := b.Subscribe()
		wg.Add(1)
		go func(ch chan interface{}, n int) {
			defer wg.Done()
			count := 0

			for {
				select {
				case data := <-ch:
					EqualString(t, publishStrs[count], data.(string))
					count++

					if data.(string) == "donePublish" {
						return
					}
				}
			}
		}(c, num)
	}

	for _, s := range publishStrs {
		b.Publish(s)
	}

	wg.Wait()
}
