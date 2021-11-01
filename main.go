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

	if isHeistOn {

		leftSafely := rand.Intn(5)

		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("The Heist has failed.")
		case 1:
			isHeistOn = false
			fmt.Println("The vault couldn't be opened due to lack of time.")
		case 2:
			isHeistOn = false
			fmt.Println("The combination was incorrect.")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 1000 + rand.Intn(1000000)
		fmt.Printf("Amount Stolen from bank: %d", amtStolen)
	}

	fmt.Printf("Status of Heist: %v", isHeistOn)
	fmt.Println(" ")
}
