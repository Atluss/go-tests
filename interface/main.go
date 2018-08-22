package main

import "fmt"

//Vehicle интефейс машины
type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

//Car структура машины
type Car struct{}

//Aircraft структура воздущного средства
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {

	tesla := Car{}
	boing := Aircraft{}
	drive(tesla)
	drive(boing)
}
