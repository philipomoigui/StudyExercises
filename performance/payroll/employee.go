package payroll

import "fmt"

type Payer interface{
	Pay() (string, float64)
}

type Employee struct {
	Id int
	FirstName string
	LastName string
}

func PayDetails(p Payer){
	fullName, pay := p.Pay()
	fmt.Printf("%s got paid with %.2f for the year\n", fullName, pay)
}