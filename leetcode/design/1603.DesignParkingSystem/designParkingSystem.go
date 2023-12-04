package main

import (
	"fmt"
)

type ParkingSystem struct {
	capacity []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[]int{0, big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.capacity[carType] == 0 {
		return false
	}
	this.capacity[carType]--
	return true
}

func (this *ParkingSystem) AddCar1(carType int) bool {
	if this.capacity[carType] > 0 {
		this.capacity[carType]--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func main() {
	parking := Constructor(1, 1, 0)

	fmt.Println(parking.AddCar(1))
	fmt.Println(parking.AddCar(2))
	fmt.Println(parking.AddCar(3))
}
