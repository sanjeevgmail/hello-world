package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
	scale       int
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
	c.scale++
}

func (c *Counter) Mult() {
	c.scale = c.scale * 2
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v scale: %v", c.total, c.lastUpdated, c.scale)
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	c.Increment()
	c.Mult()
	fmt.Println(c.String())
}
