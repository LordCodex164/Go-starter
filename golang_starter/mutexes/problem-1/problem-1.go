package main

import (
	"fmt"
	"sync"
	"time"
)

/*
you have a shared counter representing total API requests handled.

10 goroutines

Each increments the counter 1,000 times

Print final value
*/

//problem 2

type BankAccount struct {
	balance int32
	mu sync.RWMutex
}

func NewBankAccount() *BankAccount{
	return &BankAccount{
		balance: 0,
	}
}

func (b *BankAccount) readAccount() int32 {
	b.mu.RLock()
	defer b.mu.RLock()
	return b.balance
} 

//deposit money
func (b *BankAccount) Deposit(value int) int32 {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += int32(value)
	return b.balance
}

func startProcess (p *BankAccount) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Println(">deposit money", p.Deposit(100))
		<- ticker.C
	}
}

//withdraw

func  (b *BankAccount) withDraw(value int32) int32 {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance -= value
	return b.balance	
}

func WithDrawProcess (b *BankAccount) {
	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		balance := b.readAccount()
		if balance > 0 {
			fmt.Println(">withdrawing money", b.withDraw(50))
			break;
		}
		<- ticker.C
	}
}


func main(){
	fmt.Println(">>> starting process")
	account := NewBankAccount()
	go startProcess(account)
	WithDrawProcess(account)
	fmt.Println(">>account balance", account.balance)
}

