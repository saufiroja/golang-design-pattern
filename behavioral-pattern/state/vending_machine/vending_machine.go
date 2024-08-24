package vendingmachine

import "fmt"

type VendingMachineState interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}

type VendingMachine struct {
	HasItem       VendingMachineState
	ItemRequested VendingMachineState
	HasMoney      VendingMachineState
	NoItem        VendingMachineState
	CurrentState  VendingMachineState
	ItemCount     int
	ItemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	vendingMachine := &VendingMachine{
		ItemCount: itemCount,
		ItemPrice: itemPrice,
	}

	hasItemState := &HasItemState{VendingMachine: vendingMachine}

	itemRequestedState := &ItemRequestedState{VendingMachine: vendingMachine}

	hasMoneyState := &HasMoneyState{VendingMachine: vendingMachine}

	noItemState := &NoItemState{VendingMachine: vendingMachine}

	vendingMachine.SetState(hasItemState)
	vendingMachine.HasItem = hasItemState
	vendingMachine.ItemRequested = itemRequestedState
	vendingMachine.HasMoney = hasMoneyState
	vendingMachine.NoItem = noItemState

	return vendingMachine
}

func (v *VendingMachine) AddItem(count int) error {
	return v.CurrentState.AddItem(count)
}

func (v *VendingMachine) RequestItem() error {
	return v.CurrentState.RequestItem()
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.CurrentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.CurrentState.DispenseItem()
}

func (v *VendingMachine) SetState(state VendingMachineState) {
	v.CurrentState = state
}

func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.ItemCount += count
}
