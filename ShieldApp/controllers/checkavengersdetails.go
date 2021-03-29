package controllers

import (
	"bufio"
	"fmt"
	"os"
)

func CheckAvengerDetails() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Avenger Name :- ")
	scanner.Scan()
	avengerName := scanner.Text()

	avengerDetails, found := AvengerList[avengerName]

	if found {
		fmt.Printf("Person Name: %s\n", avengerDetails.PersonName)
		fmt.Printf("Abilities :%s\n", avengerDetails.Abilities)
		fmt.Printf("Mission Assigned :%d\n", len(avengerDetails.MissionAssigned))
		fmt.Printf("Mission Completed :%d\n", len(avengerDetails.MissionCompleted))

	} else {
		fmt.Println("Sorry no Avenger found try again")
	}

}
