package solid

import "fmt"

//First example where we will not follow the SRP principle
type ShoppingCart struct {
	prices []float64
}

func (s *ShoppingCart) AddProduct(price float64) {
	s.prices = append(s.prices, price)
}

func (s *ShoppingCart) CalculateTotal() float64 {
	var total float64
	for _, price := range s.prices {
		total += price
	}
	return total
}
func (s *ShoppingCart) SaveToDatabase() {
	fmt.Println("Saving order to database")
}

func (s *ShoppingCart) SendEmail() {
	fmt.Println("Sending confirmation email")
}

//Second example where we will follow the SRP principle

type Order struct {
	cost []float64
}

func (o *Order) AddCost(price float64) {
	o.cost = append(o.cost, price)
}

func (o *Order) CalculateTotal() float64 {
	var total float64
	for _, price := range o.cost {
		total += price
	}
	return total
}

type OrderRepository struct{}

func (r *OrderRepository) SaveToDatabase(order *Order) {
	fmt.Println("Saving order to database")
}

type EmailService struct{}

func (e *EmailService) SendEmail(order *Order) {
	fmt.Println("Sending confirmation email")
}

func main() {
	cart := ShoppingCart{}

	cart.AddProduct(100)
	cart.AddProduct(200)

	fmt.Println("Total:", cart.CalculateTotal())

	cart.SaveToDatabase()
	cart.SendEmail()

	test := Order{}
	test.AddCost(100)
	test.AddCost(200)

	fmt.Println("Total:", test.CalculateTotal())

	repo := OrderRepository{}
	repo.SaveToDatabase(&test)
	emailService := EmailService{}
	emailService.SendEmail(&test)
}
