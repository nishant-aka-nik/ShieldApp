package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckMissionDetails() {

	scanner := bufio.NewScanner(os.Stdin)

	if len(AllMission) == 0 {
		fmt.Println("No Mission has been assigned to an Avenger")
	} else {

		fmt.Println("Enter Mission Name :- ")
		scanner.Scan()
		missionName := scanner.Text()

		missionDetails, found := AllMission[missionName]

		if found {
			fmt.Printf("Mission Name: %s\n", missionDetails.MissionName)
			fmt.Printf("Mission Details: %s\n", missionDetails.MissionDetails)
			fmt.Printf("Mission Status: %s\n", missionDetails.MissionStatus)
			AllAssignedAvenger := strings.Join(missionDetails.AvengersOnMisson, ", ")

			fmt.Printf("Mission Avengers: %s\n", AllAssignedAvenger)
		} else {
			fmt.Println("Sorry no mission found try again")
		}

	}
}
