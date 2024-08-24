package vendingmachine

import "fmt"

type HasItemState struct {
	VendingMachine *VendingMachine
}

func (h *HasItemState) RequestItem() error {
	if h.VendingMachine.ItemCount == 0 {
		h.VendingMachine.SetState(h.VendingMachine.NoItem)
		return fmt.Errorf("no item present")
	}

	fmt.Printf("Item requested\n")
	h.VendingMachine.SetState(h.VendingMachine.ItemRequested)
	return nil
}

func (h *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	h.VendingMachine.IncrementItemCount(count)
	return nil
}

func (h *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (h *HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}
