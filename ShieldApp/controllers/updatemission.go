package controllers

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func UpdateMission() {
	scanner := bufio.NewScanner(os.Stdin)

	if len(AllMission) == 0 {
		fmt.Println("No Mission has been assigned to an Avenger")

	} else {
		fmt.Println("Enter Mission Name :- ")
		scanner.Scan()
		missionName := scanner.Text()

		missionCompleteDetails, found := AllMission[missionName]

		if found {

			if missionCompleteDetails.MissionStatus == "Completed" {
				fmt.Println("Sorry this mission is already completed")
			} else {

				fmt.Println("Enter new status :- ")
				scanner.Scan()
				newMissionStatus := scanner.Text()

				if newMissionStatus != "Completed" {
					fmt.Println("Sorry invalid mission status")
					return
				}

				missionCompleteDetails.MissionStatus = newMissionStatus

				for _, v := range missionCompleteDetails.AvengersOnMisson {
					avenger := AvengerList[v]
					indexOfMission := sort.SearchStrings(avenger.MissionAssigned, missionCompleteDetails.MissionName)
					avenger.MissionAssigned[indexOfMission] = avenger.MissionAssigned[len(avenger.MissionAssigned)-1]
					avenger.MissionAssigned[len(avenger.MissionAssigned)-1] = ""
					avenger.MissionAssigned = avenger.MissionAssigned[:len(avenger.MissionAssigned)-1]

					//
					avenger.MissionCompleted = append(avenger.MissionCompleted, missionCompleteDetails.MissionName)

					//
					if len(avenger.MissionAssigned) == 0 {
						avenger.Status = "Available"
					}
				}

				for _, v := range missionCompleteDetails.AvengersOnMisson {
					avengerDevice, avengerName := CheckAvengerDevice(v)
					fmt.Printf("%s notification sent to %s\n", avengerDevice, avengerName)
				}

			}

		} else {
			fmt.Println("Sorry no mission found try again")
		}

	}

}
