// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawRequest struct {
	amount  int
	confirm chan<- bool
}

var deposits = make(chan int)             // send amount to deposit
var balances = make(chan int)             // receive balance
var withdraw = make(chan withdrawRequest) //

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	// need a separate confirmation for every withdrawal
	confirm := make(chan bool)
	withdraw <- withdrawRequest{amount: amount, confirm: confirm}
	return <-confirm // blocks
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case req := <-withdraw:
			if req.amount <= balance {
				balance -= req.amount
				req.confirm <- true
			} else {
				req.confirm <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
