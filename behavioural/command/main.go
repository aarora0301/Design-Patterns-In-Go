package main

import "fmt"

type Robot struct {
	x int
	y int
}

func (r *Robot) moveUp() {
	r.y++
	fmt.Println("Robot moved up")
}

func (r *Robot) moveDown() {
	r.y--
	fmt.Println("Robot moved down")
}

func (r *Robot) moveLeft() {
	r.x--
	fmt.Println("Robot moved left")
}

func (r *Robot) moveRight() {
	r.x++
	fmt.Println("Robot moved right")
}

type Player interface {
	moveUp()
	moveDown()
	moveLeft()
	moveRight()
}

type Command interface {
	execute()
}

type MoveUpCommand struct {
	player Player
}

func (m *MoveUpCommand) execute() {
	m.player.moveUp()
}

type MoveDownCommand struct {
	player Player
}

func (m *MoveDownCommand) execute() {
	m.player.moveDown()
}

type MoveLeftCommand struct {
	player Player
}

func (m *MoveLeftCommand) execute() {
	m.player.moveLeft()
}

type MoveRightCommand struct {
	player Player
}

func (m *MoveRightCommand) execute() {
	m.player.moveRight()
}

type button struct {
	command Command
}

func (b *button) press() {
	b.command.execute()
}

func main() {
	robot := &Robot{}
	up := &MoveUpCommand{robot}
	down := &MoveDownCommand{robot}
	left := &MoveLeftCommand{robot}
	right := &MoveRightCommand{robot}

	upButton := &button{up}
	downButton := &button{down}
	leftButton := &button{left}
	rightButton := &button{right}

	upButton.press()
	downButton.press()
	leftButton.press()
	rightButton.press()
}
