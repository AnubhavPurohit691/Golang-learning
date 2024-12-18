package main

import "fmt"

// should be open for extention close for modification
type paymenter interface {
	pay(amount int)
}
type payment struct {
	gateway paymenter
}

func (p payment) Payment(amount int) {
	paye := razor{}
	paye.pay(amount)
}

type razor struct{}

func (r razor) pay(amount int) {
	fmt.Println("payment using razor pay", amount)
}

func main() {
	razorGw := razor{}
	razorpay := payment{
		gateway: razorGw,
	}
	razorpay.Payment(4400)

}
