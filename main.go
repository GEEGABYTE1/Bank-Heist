package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	isHeistOn := true
	eludedGuards := rand.Intn(100)
	rand.Seed(time.Now().UnixNano())

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step. ")

	} else {
		isHeistOn = false
	}

	fmt.Printf("Status of Hesit %v", isHeistOn)
	fmt.Println(" ")
}
