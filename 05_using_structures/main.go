package main

import "fmt"

const uMax float64 = 65535
const kmToMiles float64 = 1.6093

//Car is..
type Car struct {
	gasPedal      uint16
	breakpedal    uint16
	steeringWheel int16
	topSpeedKMH   float64
}

// Value reciver
func (c Car) kmph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / uMax)
}

// Pointer reciver
func (c *Car) newTopSpeed(speed float64) {
	c.topSpeedKMH = speed
}

func (c *Car) valueReciver(breakPedal uint16) {
	c.breakpedal = breakPedal
}
func setTopSpeed(c Car, speed float64) Car {
	c.topSpeedKMH = speed
	return c
}

func main() {
	theCar := Car{
		gasPedal:      223,
		breakpedal:    0,
		steeringWheel: 1265,
		topSpeedKMH:   250,
	}
	fmt.Println(theCar.gasPedal)
	fmt.Println(theCar.kmph())
	theCar = setTopSpeed(theCar, 350)
	fmt.Println(theCar.kmph())
	theCar.valueReciver(3)
	fmt.Printf("Breaks %d\n", theCar.breakpedal)

}
