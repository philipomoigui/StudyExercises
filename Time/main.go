package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Date(2020, 5, 9, 02, 38, 56, 23837, time.UTC)
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())

	fmt.Printf(hour + ":" + minute + ":" + second + " " + year + "/" + month + "/" + day)
	elapsedTime()
	fmt.Println("___________________")
	futureDate()
	fmt.Println("___________________")
	localTime()

}

func elapsedTime() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	length := end.Sub(start)
	fmt.Println("\nThe execution took exactly", length.Seconds(), "seconds!")
}

func futureDate() {
	current := time.Now()
	fmt.Println(current.Format(time.ANSIC))
	// caluclating future time duration
	sss := time.Duration(21966 * time.Second)

	futureDate := current.Add(sss)
	fmt.Println("6 hours, 6 minutes and 6 seconds from now will be :", futureDate.Format(time.ANSIC))

}

func localTime() {
	current := time.Now()
	
	NYZone, _ := time.LoadLocation("America/Los_Angeles")
	NYTime := current.In(NYZone)
	
	LAZone, _ := time.LoadLocation("America/New_York")
	LATime := current.In(LAZone)
	
	
	fmt.Println("The local current time is: ", current.Format(time.ANSIC))
	fmt.Println("The current time in New York is: ", NYTime.Format(time.ANSIC))
	fmt.Println("The current time in Los Angeles is: ", LATime.Format(time.ANSIC))
}
