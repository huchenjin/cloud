package part1

import "testing"

func TestString(t *testing.T) {
	AlertArray()
}

func TestChan(t *testing.T) {
	var ch = make(chan int, 10)
	go Producer(ch)
	Consumer(ch)
}
