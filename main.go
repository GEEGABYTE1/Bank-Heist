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

	openedVault := rand.Intn(100)

	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn && openedVault <= 70 {
		isHeistOn = false
		fmt.Println("The vault can't be opened")
	}

	fmt.Printf("Status of Heist: %v", isHeistOn)
	fmt.Println(" ")
}
