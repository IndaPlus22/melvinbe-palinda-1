package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	currentTime := time.Now()
	formattedTime := time.Time.Format(currentTime, "15.04.05:")

	fmt.Println("The time is", formattedTime, text)

	time.Sleep(delay)
	Remind(text, delay)
}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	Remind("Time to sleep", 60*time.Second)
}
