package main

import "fmt"

// Define products
type Chair interface {
	sitOn()
}

type CoffeeTable interface {
	getCoffeeCup()
}

type Sofa interface {
	lieOn()
}

type ModernChair struct{}

func (mc ModernChair) sitOn() {
	fmt.Println("You are sitting on a Modern chair.")
}

type VictorianChair struct{}

func (vc VictorianChair) sitOn() {
	fmt.Println("You are sitting on a Victorian chair.")
}

type ModernCoffeeTable struct{}

func (mc ModernCoffeeTable) getCoffeeCup() {
	fmt.Println("Take this Modern coffee cup.")
}

type VictorianCoffeetable struct{}

func (vc VictorianCoffeetable) getCoffeeCup() {
	fmt.Println("Take this Victorian coffee cup.")
}

type ModernSofa struct{}

func (ms ModernSofa) lieOn() {
	fmt.Println("You are lying on a brand new Modern sofa.")
}

type VictorianSofa struct{}

func (vs VictorianSofa) lieOn() {
	fmt.Println("You are lying on a brand new Victorian sofa.")
}

type Factory interface {
	createChair() Chair
	createCoffeeTable() CoffeeTable
	createSofa() Sofa
	getFactoryName() string
}

type ModernFactory struct{}

func (mf ModernFactory) createChair() Chair {
	return ModernChair{}
}

func (mf ModernFactory) createCoffeeTable() CoffeeTable {
	return ModernCoffeeTable{}
}

func (mf ModernFactory) createSofa() Sofa {
	return ModernSofa{}
}

func (mf ModernFactory) getFactoryName() string {
	return "Modern Factory"
}

type VictorianFactory struct{}

func (vf VictorianFactory) createChair() Chair {
	return VictorianChair{}
}

func (vf VictorianFactory) createCoffeeTable() CoffeeTable {
	return VictorianCoffeetable{}
}

func (vf VictorianFactory) createSofa() Sofa {
	return VictorianSofa{}
}

func (vf VictorianFactory) getFactoryName() string {
	return "Victorian Factory"
}

func main() {
	var factory Factory
	factory = ModernFactory{}

	fmt.Println(factory.getFactoryName())
	chair := factory.createChair()
	chair.sitOn()

	coffeTable := factory.createCoffeeTable()
	coffeTable.getCoffeeCup()

	sofa := factory.createSofa()
	sofa.lieOn()

	factory = VictorianFactory{}

	fmt.Println(factory.getFactoryName())
	chair = factory.createChair()
	chair.sitOn()

	coffeTable = factory.createCoffeeTable()
	coffeTable.getCoffeeCup()

	sofa = factory.createSofa()
	sofa.lieOn()
}
