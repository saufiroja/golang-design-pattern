package vendingmachine

import "fmt"

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.ItemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", i.VendingMachine.ItemPrice)
	}

	fmt.Printf("Money inserted: %d\n", money)
	i.VendingMachine.SetState(i.VendingMachine.HasMoney)
	return nil
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}
