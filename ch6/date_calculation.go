package main

import (
	"fmt"
	"time"
)

func main() {
	//time and date
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n ", t)
	// time now
	timeNow := time.Now()
	Tomorrow := timeNow.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %d", Tomorrow.Day())
	// example
	// fmt.Printf("the time is %s\n ", timeNow)
	// fmt.Println("the month is", t.Month())
	// fmt.Println("the day is", t.Day())
	// fmt.Println("today is", t.Weekday())
	// tommorrow := t.AddDate(0, 0, 1)
	// fmt.Printf("Tommorrow is %v,%v,%v,%v", tommorrow.Weekday(), tommorrow.Month(), tommorrow.Day(), tommorrow.Year())

}