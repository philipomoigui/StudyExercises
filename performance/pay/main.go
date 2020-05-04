package main

import (
	"fmt"
	"os"
	pp "performance/payroll"
)

var employeeReview = make(map[string]interface{})

func init() {
	fmt.Println("Welcome to the Employee Pay and Performance Review")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func init() {
	employeeReview["WorkQuality"] = 5
  	employeeReview["TeamWork"] = 2
  	employeeReview["Communication"] = "poor"
  	employeeReview["Problem-solving"] = 4
  	employeeReview["Dependability"] = "unsatisfactory"
}

func main() {
	d := pp.Developer{Individual: pp.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := pp.Manager{Individual: pp.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	pp.PayDetails(d)
	pp.PayDetails(m)
	 
	err := d.ReviewRating()
 	if err != nil {
   	 fmt.Println(err)
    	os.Exit(1)
}

}