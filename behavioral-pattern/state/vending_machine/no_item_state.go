package vendingmachine

import "fmt"

type NoItemState struct {
	VendingMachine *VendingMachine
}

func (n *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) AddItem(count int) error {
	n.VendingMachine.IncrementItemCount(count)
	n.VendingMachine.SetState(n.VendingMachine.HasItem)
	return nil
}

func (n *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}
