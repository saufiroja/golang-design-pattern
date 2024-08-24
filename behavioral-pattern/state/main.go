package main

import (
	"fmt"
	vendingmachine "fundamental/golang-design-pattern/behavioral-pattern/state/vending_machine"
)

func main() {
	itemCount := 1
	itemPrice := 10
	vendingMachine := vendingmachine.NewVendingMachine(itemCount, itemPrice)

	err := vendingMachine.RequestItem()
	if err != nil {
		println(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		println(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		println(err.Error())
	}

	fmt.Println("done")

	err = vendingMachine.AddItem(2)
	if err != nil {
		println(err.Error())
	}

	fmt.Println("done")

	err = vendingMachine.RequestItem()
	if err != nil {
		println(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		println(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		println(err.Error())
	}
}
