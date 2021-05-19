// 컨베이어벨트처럼 사용하는 방식!
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch1 := make(chan Car)
	ch2 := make(chan Car)
	ch3 := make(chan Car)

	go StratWork(ch1)
	go MakeTire(ch1, ch2)
	go MakeEngine(ch2, ch3)

	for {
		result := <-ch3
		fmt.Println(result)
	}
}

type Car struct {
	val string
}

func StratWork(ch1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		ch1 <- Car{val: "Car" + strconv.Itoa(i) + ": "}
		i++
	}
}

func MakeTire(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Tire, "

		outChan <- car
	}
}

func MakeEngine(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Engine, "

		outChan <- car
	}
}
