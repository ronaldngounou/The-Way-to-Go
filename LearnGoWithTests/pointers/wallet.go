package main

import "fmt"

type Bitcoin int 

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// In Go, when you call a function or a method the arguments are copied.
// When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.

//You can see that the addresses of the two balances are different.
// So when we change the value of the balance inside the code,
// we are working on a copy of what came from the test. Therefore the balance in the test is unchanged.
// let us point to some values and then let us change them. So rather than taking a copy of the whole Wallet, 
// we instead take a pointer to that wallet so that we can change the original values within it.
// The difference is the receiver type is *Wallet rather than Wallet which 
// you can read as "a pointer to a wallet".

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return (w).balance // can also write (*w).balance car they are automatically dereferenced
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount 
}

