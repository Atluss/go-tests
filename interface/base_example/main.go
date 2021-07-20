package main

import "fmt"

// Vehicle интерфейс машины
type Vehicle interface {
	move()
}

// drive функция вызывает функцию соответствующего интерфейса
func drive(vehicle Vehicle) {
	vehicle.move()
}

// Car структура машины
type Car struct {
	Name string
}

func (c Car) move() {
	fmt.Println(c.Name)
}

// Aircraft структура воздушного средства
type Aircraft struct{}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {
	tesla := Car{"Tesla"}
	boing := Aircraft{}
	v := Vehicle(tesla)
	tesla.Name = "VAZ"

	v.move()     // Tesla
	drive(tesla) // Vaz
	drive(boing)
}
