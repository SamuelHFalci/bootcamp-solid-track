package stripe

import "fmt"

type StripeService struct {

}

func NewPaypalService() *StripeService {
	return &StripeService{}
}

func (p *StripeService) Pay(total float64){
	fmt.Printf("Paying with Paypal...: %2.f", total)
}