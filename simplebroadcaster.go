package goutils

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type SimpleBroadcaster struct {
	sync.RWMutex
	chanSize    int
	channels    []chan interface{}
	inChan      chan interface{}
	donePublish chan struct{}
}

func NewSimpleBroadcaster(chanSize int) *SimpleBroadcaster {
	b := &SimpleBroadcaster{
		chanSize:    chanSize,
		channels:    make([]chan interface{}, 0),
		inChan:      make(chan interface{}, chanSize),
		donePublish: make(chan struct{}),
	}

	b.runPublish()

	return b
}

func (b *SimpleBroadcaster) Subscribe() (chan interface{}, int) {
	ch := make(chan interface{}, b.chanSize)
	b.Lock()
	b.channels = append(b.channels, ch)
	b.Unlock()
	return ch, len(b.channels) - 1
}

func (b *SimpleBroadcaster) Publish(data interface{}) {
	b.inChan <- data
}

func (b *SimpleBroadcaster) runPublish() chan interface{} {
	go func(b *SimpleBroadcaster) {
		for {
			select {
			case data := <-b.inChan:
				b.RLock()
				for i := range b.channels {
					b.channels[i] <- data
				}
				b.RUnlock()
			case <-b.donePublish:
				log.Debug("broadcaster shutdown")
				return
			}
		}
	}(b)

	return b.inChan
}

func (b *SimpleBroadcaster) GetPublishChan() chan interface{} {
	return b.inChan
}

func (b *SimpleBroadcaster) Shutdown() {
	b.donePublish <- struct{}{}
}
