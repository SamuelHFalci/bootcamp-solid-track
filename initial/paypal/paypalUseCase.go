package paypal

import "fmt"

type PaypalService struct {

}

func NewPaypalService() *PaypalService {
	return &PaypalService{}
}

func (p *PaypalService) Pay(total float64){
	fmt.Printf("Paying with Paypal...: %2.f", total)
}