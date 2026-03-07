package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber 	string
	Balance		  	float64
	OwnerName 		string
}

type SavingsAccount struct {
	Account			//embed the account struct here
	InterestRate    float64
}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.Balance += amount
	fmt.Println("balance deposited", amount)
	return nil
}

func (a Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if a.Balance < amount {
		return errors.New("insufficient balance")
	}
	a.Balance -= amount
	fmt.Println("balance withdrawn successfully")
	return nil
}

func (a Account) String() string {
	return fmt.Sprintf("Here is Account Number:%v with a balance of %v", a.AccountNumber, a.Balance)
}

func (s *SavingsAccount) addInterest() error {
	interest := s.Balance * s.InterestRate
	if err := s.Deposit(interest); err != nil {
		return err
	}
	fmt.Println("Interest added successfully")
	return nil
}

func main() {

	acc := Account{
		AccountNumber: "0569327993",
		Balance: 1023.45,
		OwnerName: "daniel isaaac",
	}

	savingAcc := SavingsAccount{
		Account: acc,
		InterestRate: 4.05,
	}

	savingAcc.addInterest()

	fmt.Println(acc)
}