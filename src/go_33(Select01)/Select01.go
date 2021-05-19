package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	carCh1 := make(chan Car)
	carCh2 := make(chan Car)
	carCh3 := make(chan Car)

	planeCh1 := make(chan Plane)
	planeCh2 := make(chan Plane)
	planeCh3 := make(chan Plane)

	go StratCarWork(carCh1)
	go StratPlaneWork(planeCh1)
	go MakeTire(carCh1, planeCh1, carCh2, planeCh2)
	go MakeEngine(carCh2, planeCh2, carCh3, planeCh3)

	for {
		select {
		case carResult := <-carCh3:
			fmt.Println(carResult)
		case planeResult := <-planeCh3:
			fmt.Println(planeResult)
		}
	}
}

type Car struct {
	val string
}

type Plane struct {
	val string
}

func StratCarWork(ch1 chan Car) {
	i := 0
	for {
		j := rand.Intn(100)
		if j == 0 {
			j = 1
		}
		time.Sleep(time.Duration(j) * time.Millisecond)
		ch1 <- Car{val: "Car" + strconv.Itoa(i) + ": "}
		i++
	}
}

func StratPlaneWork(ch1 chan Plane) {
	i := 0
	for {
		j := rand.Intn(100)
		if j == 0 {
			j = 1
		}
		time.Sleep(time.Duration(j) * time.Millisecond)
		ch1 <- Plane{val: "Plane" + strconv.Itoa(i) + ": "}
		i++
	}
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Tire_P, "
			outPlaneChan <- plane
		}
	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlaneChan <- plane
		}
	}
}
