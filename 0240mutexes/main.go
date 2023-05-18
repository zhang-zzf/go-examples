package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mutex    sync.Mutex
	counters map[string]int
}

func (c *Container) AddCount(key string, delta int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counters[key] += delta
}

func NewContainer() *Container {
	return &Container{
		counters: map[string]int{},
	}
}

func main() {
	container := NewContainer()
	container.AddCount("a", 2)
	fmt.Println("Container:", container.counters)
	var wg sync.WaitGroup
	f := func(key string, delta int) {
		defer wg.Done()
		container.AddCount(key, delta)
	}
	wg.Add(3)
	go f("a", 2)
	go f("a", 2)
	go f("a", 2)
	wg.Wait()
	fmt.Println("Container:", container.counters)
}
