package main

import "fmt"
import "time"

func backgroundFunction(delay time.Duration, done chan bool) {
	time.Sleep(delay)
	fmt.Println("Slept for", delay.Seconds(), "seconds")
	done <- true
}

func main() {
	var spawn = 5
	var control chan bool = make(chan bool) // HL

	for n := 1; n <= spawn; n++ {
		delay := time.Duration(n) * time.Second
		go backgroundFunction(delay, control)
	}

	for done := 0; done < spawn; done++ {
		<-control
	}
}
