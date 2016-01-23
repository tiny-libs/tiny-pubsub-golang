package pubsub

import (
	// "log"
	"runtime"
)

type channel struct {
	namespace string
	subscriptions []chan interface{}
	channels map[string]*channel
}

type pubsub struct {
	channel
}

func newPubsub() *pubsub {
	ps := new(pubsub)

	ps.channels = make(map[string]*channel)
	return ps
}

func (ps *pubsub) on(namespace string) chan interface{} {
	chann, ok := ps.channels[namespace]
	if(!ok) {
		chann = &channel{ namespace : namespace}
		ps.channels[namespace] = chann
	}
	currSub := make(chan interface{})
	chann.subscriptions = append(chann.subscriptions, currSub)

	return currSub
}

func (ps *pubsub) publish(namespace string, args string) *pubsub {
	chann := ps.channels[namespace]

	for _, sub := range chann.subscriptions {
		go func(sub chan interface{}) {
			// log.Println("sub", sub);
			sub <- args
			runtime.Gosched()
		}(sub)
	}
	return ps
}