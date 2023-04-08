package main

import (
	"fmt"
	"sync"
)

type iPool interface {
	getID() string
}

type pool struct {
	idle     []iPool
	active   []iPool
	capacity int
	mulock   *sync.Mutex
}

func initPool(poolObjects []iPool) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Pool objects cannot be empty")
	}

	active := make([]iPool, 0)
	return &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   &sync.Mutex{},
	}, nil
}

func (p *pool) loan() (iPool, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No idle objects")
	}

	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan object with id: %s", obj.getID())
	fmt.Println()
	return obj, nil
}

func (p *pool) release(obj iPool) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if err := p.remove(obj); err != nil {
		return err
	}
	p.idle = append(p.idle, obj)
	fmt.Printf("Release object with id: %s", obj.getID())
	fmt.Println()
	return nil
}

func (p *pool) remove(obj iPool) error {
	currentActive := p.active
	for i, v := range currentActive {
		if v.getID() == obj.getID() {
			p.active = append(currentActive[:i], currentActive[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Object not found in active pool")
}

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}

func main() {
	connections := make([]iPool, 0)
	for i := 0; i < 10; i++ {
		connections = append(connections, &connection{id: fmt.Sprintf("connection-%d", i)})
	}

	pool, err := initPool(connections)
	if err != nil {
		panic(err)
	}
	conn1, err := pool.loan()
	if err != nil {
		panic(err)
	}
	conn2, err := pool.loan()
	if err != nil {
		panic(err)
	}
	pool.release(conn1)
	pool.release(conn2)
}
