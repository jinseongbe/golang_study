package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance int
	mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
	a.balance -= val
}

func (a *Account) Deposit(val int) {
	a.balance += val
}

func (a *Account) Balance() int {
	balance := a.balance
	return balance
}

var accounts []*Account

func Transfer(sender, receiver, money int) {

	// DeadLock!!!!
	accounts[sender].mutex.Lock() // 쓰레드1에서 예를 락 하고 다음 락을 하려고함
	time.Sleep(1000 * time.Millisecond)
	accounts[receiver].mutex.Lock() // 그러나 쓰레드2에서 예를 락 해버려서 쓰레드1이 이 자원에 접근하지 못함, 결국 둘다 멈춰버림!

	accounts[sender].Widthdraw(money)
	accounts[receiver].Deposit(money)

	accounts[receiver].mutex.Unlock()
	accounts[sender].mutex.Unlock()

	fmt.Println("Transfer", sender, receiver, money)
}

func GetTotalBalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

func RandomTrasfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTrasfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	go func() {
		for {
			Transfer(0, 1, 100)
		}
	}()

	go func() {
		for {
			Transfer(1, 0, 100)
		}
	}()

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}
