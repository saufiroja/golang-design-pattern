package vendingmachine

import "fmt"

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (h *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item already dispensed")
}

func (h *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item already dispensed")
}

func (h *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (h *HasMoneyState) DispenseItem() error {
	fmt.Printf("item dispensed\n")

	h.VendingMachine.ItemCount -= 1
	if h.VendingMachine.ItemCount == 0 {
		h.VendingMachine.SetState(h.VendingMachine.NoItem)
	} else {
		h.VendingMachine.SetState(h.VendingMachine.HasItem)
	}

	return nil
}
