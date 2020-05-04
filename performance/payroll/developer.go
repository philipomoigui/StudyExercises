package payroll

import (
	"fmt"
	"strings"
	"errors"
) 

type Developer struct{
	Individual Employee 
	HourlyRate float64
	HoursWorkedInYear float64
	Review map[string]interface{}
}

func (d Developer)fullName() string{
	fullName := d.Individual.FirstName + " " + d.Individual.LastName
	return fullName
}

func (d Developer)Pay() (string, float64){
	fullName := d.Individual.FirstName+ " " + d.Individual.LastName
	return fullName, d.HourlyRate * d.HoursWorkedInYear
}

func reviewToInt(str string)(int, error){
	switch strings.Title(str) {
		case "Excellent":
			return 5, nil
		case "Good":
			return 4, nil
		case "Fair":
			return 3, nil
		case "Poor":
			return 2, nil
		case "Unsatisfactory":
			return 1, nil
		default:
			return 0, errors.New("Invalid rating: " + str)
		}
	}
	
	func overAllRating (i interface{})(int, error){
		switch v := i.(type){
		case int:
			return v, nil
		case string: 
		rating, err := reviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
		default:
		return 0, errors.New("Unknown type")
		}
		
	}
	
	func (d Developer)ReviewRating() error{
		total := 0
		for _, v := range d.Review {
			rating, err := overAllRating(v)
			if err != nil {
				return err
			}
			total += rating
		}
		averageRating := float64(total) / float64(len(d.Review))
		fmt.Printf("%s got a review rating of %.2f\n", d.fullName(), averageRating)
		return nil
	}