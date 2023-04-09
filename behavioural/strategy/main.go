package main

import "fmt"

type cache struct {
	storage      map[int]int
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

type evictionAlgo interface {
	evict(*cache)
}

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting using LRU")
}

type fifo struct{}

func (f *fifo) evict(c *cache) {
	fmt.Println("Evicting using FIFO")
}

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting using LFU")

}
