package main

import (
	"fmt"
	"log"
)

// state Actions that can be performed on the vending machine
type state interface {
	addItem(int) error
	selectItem() error
	dispenseItem() error
	insertMoney(money int) error
}

// VendingMachine is the main struct that holds the states
type vendingMachine struct {
	hasItem       state
	noItem        state
	itemRequested state
	hasMoney      state

	currentState state
	itemCount    int
	itemPrice    int
}

func (v *vendingMachine) selectItem() error {
	return v.currentState.selectItem()
}
func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) addItem(count int) error {
	n.vendingMachine.incrementItemCount(count)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

func (n *noItemState) insertMoney(money int) error {
	fmt.Println("No item in the machine")
	return nil
}

func (n *noItemState) dispenseItem() error {
	fmt.Println("No item in the machine")
	return nil
}

func (n *noItemState) selectItem() error {
	fmt.Println("No item in the machine")
	return nil
}

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (h *hasItemState) addItem(count int) error {
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *hasItemState) insertMoney(money int) error {
	fmt.Println("Please select an item first")
	return nil
}

func (h *hasItemState) dispenseItem() error {
	fmt.Println("Please select an item first")
	return nil
}

func (i *hasItemState) selectItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) addItem(count int) error {
	fmt.Println("Item dispense in progress")
	return nil
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	fmt.Println("Insert Money first")
	return nil
}

func (i *itemRequestedState) selectItem() error {
	fmt.Println("Item already selected")
	return nil
}

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (h *hasMoneyState) addItem(count int) error {
	fmt.Println("Item dispense in progress")
	return nil
}
func (h *hasMoneyState) insertMoney(money int) error {
	fmt.Println("OOSSS! You have already inserted money")
	return nil
}

func (i *hasMoneyState) dispenseItem() error {
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

func (i *hasMoneyState) selectItem() error {
	fmt.Println("Item already selected")
	return nil
}

func NewVendingMachine(itemcount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemcount,
		itemPrice: itemPrice,
	}

	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.noItem = noItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	return v
}

func main() {
	vendingMachine := NewVendingMachine(10, 10)
	err := vendingMachine.selectItem()
	if err != nil {
		log.Fatal(err)
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatal(err)
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatal(err)
	}

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatal(err)
	}
}
