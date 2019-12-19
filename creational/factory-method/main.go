package main

import "fmt"

type Transport interface {
	deliver()
}

type Truck struct{}

func (t Truck) deliver() {
	fmt.Println("Delivered by a 🚚")
}

type Ship struct{}

func (s Ship) deliver() {
	fmt.Println("Delivered by a 🚢")
}

type TransportFactory interface {
	createTransport() Transport
}

type RoadLogistic struct{}

func (r RoadLogistic) createTransport() Transport {
	return Truck{}
}

type ShipLogistic struct{}

func (s ShipLogistic) createTransport() Transport {
	return Ship{}
}

type Logistics struct {
	transportFactory TransportFactory
}

func (l Logistics) planDelivery() {
	transport := l.transportFactory.createTransport()
	transport.deliver()
}

func main() {
	road := Logistics{transportFactory: RoadLogistic{}}
	ship := Logistics{transportFactory: ShipLogistic{}}

	road.planDelivery()
	ship.planDelivery()
}
