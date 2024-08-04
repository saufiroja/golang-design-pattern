package main

import (
	"fmt"
)

func main() {
	game := newGame()

	// add terrorist
	game.addTerorist(TerroristDressType)
	game.addTerorist(TerroristDressType)
	game.addTerorist(TerroristDressType)
	game.addTerorist(TerroristDressType)
	game.addTerorist(TerroristDressType)

	// add counter terrorist
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}

// Flyweight factory
const (
	// Terrorist dress type
	TerroristDressType = "tDress"
	// Counter-Terrorist dress type
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress // -> dress interface
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

// flyweight interface
type Dress interface {
	getColor() string
}

// concrete flyweight object
type TerroristDress struct {
	color string
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{
		color: "red",
	}
}

func (t *TerroristDress) getColor() string {
	return t.color
}

// concrete flyweight object
type CounterTerroristDress struct {
	color string
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{
		color: "green",
	}
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

// context

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := dressFactorySingleInstance.getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

// client code
type game struct {
	terrorist        []*Player
	counterTerrorist []*Player
}

func newGame() *game {
	return &game{
		terrorist:        make([]*Player, 1),
		counterTerrorist: make([]*Player, 1),
	}
}

func (g *game) addTerorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorist = append(g.terrorist, player)
	return
}

func (g *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorist = append(g.counterTerrorist, player)
	return
}
