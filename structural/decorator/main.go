package main

import "fmt"

type Seat interface {
	getPrice() int
}

type EconomySeat struct {
}

func (e *EconomySeat) getPrice() int {
	return 100
}

type BusinessSeat struct{}

func (b *BusinessSeat) getPrice() int {
	return 200
}

type WindowSeat struct {
	seat Seat
}

func (s *WindowSeat) getPrice() int {
	return s.seat.getPrice() + 50
}

type AisleSeat struct {
	seat Seat
}

func (s *AisleSeat) getPrice() int {
	return s.seat.getPrice() + 25
}

func main() {
	economySeat := &EconomySeat{}

	// select window seat
	windowSeat := &WindowSeat{seat: economySeat}
	fmt.Println("Price", windowSeat.getPrice())
}
