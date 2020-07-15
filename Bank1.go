package main

import (
	"fmt"
	"time"
)

type draw struct {
	Isok   chan bool
	Amount int
}

var withdraw = make(chan draw)
var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	isok := make(chan bool)
	withdraw <- draw{isok, amount}
	flag := <-isok
	if(!flag){
		fmt.Println("余额不足!")
	}
	return flag
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw_ := <-withdraw:
			if balance >= withdraw_.Amount {
				balance -= withdraw_.Amount
				withdraw_.Isok <- true
			}else {
				withdraw_.Isok <- false
			}
		}
	}
}

func main(){
	go teller()

	Deposit(10)

	go Withdraw(12)

	time.Sleep(5 * time.Second)
}
