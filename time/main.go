package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("Dates and times", n)

	t := time.Date(2021, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("The time was", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
}
