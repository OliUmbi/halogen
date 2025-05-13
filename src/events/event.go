package events

import "sync"

type Observer struct {
	Name     string
	Callback func(payload any)
}

type Bus struct {
	observers []*Observer
	lock      sync.Mutex
}

var bus = &Bus{}

func Subscribe(observer *Observer) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	bus.observers = append(bus.observers, observer)
}

func Unsubscribe(observer *Observer) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	var observers = make([]*Observer, 0, len(bus.observers))

	for _, o := range bus.observers {
		if o != observer {
			observers = append(observers, o)
		}
	}
	
	bus.observers = observers
}

func Publish(name string, payload any) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	for _, observer := range bus.observers {
		if observer.Name == name {
			go observer.Callback(payload)
		}
	}
}
