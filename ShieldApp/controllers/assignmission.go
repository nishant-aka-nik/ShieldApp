package controllers

import (
	"bufio"
	"fmt"
	"os"
	"shieldapp/models"
	"strings"
)


func AssignMission() {

	//
	var mission models.Mission
	scanner := bufio.NewScanner(os.Stdin)

	//
	fmt.Println("Enter Avengers")
	scanner.Scan()
	assignedAvengers := strings.Split(scanner.Text(), ",")

	for k, v := range assignedAvengers {
		assignedAvengers[k] = strings.ToLower(strings.TrimSpace(v))
	}

	if len(assignedAvengers) == 0 {
		fmt.Println("Please enter the name of avengers to be assigned")
		return
	} else {
		for _, v := range assignedAvengers {

			if avengerName, check := CheckAvengerisAvailable(v); !check {
				fmt.Printf("Sorry %s is already occupied in 2 missions\n", avengerName)
				return
			}

			mission.AvengersOnMisson = append(mission.AvengersOnMisson, v)

			AvengerList[v].Status = "On Mission"

		}
	}

	//
	fmt.Println("Enter Mission Name")
	scanner.Scan()
	mission.MissionName = scanner.Text()

	for _, v := range assignedAvengers {
		AvengerList[v].MissionAssigned = append(AvengerList[v].MissionAssigned, mission.MissionName)
	}

	//
	fmt.Println("Enter Mission Details")
	scanner.Scan()
	mission.MissionDetails = scanner.Text()

	//
	mission.MissionStatus = "Assigned"

	//
	AllMission[mission.MissionName] = &mission

	//
	fmt.Println("Mission has been assigned.")

	//Successfull message display based on the avenger's device
	for _, v := range assignedAvengers {
		avengerDevice, avengerName := CheckAvengerDevice(v)
		fmt.Printf("%s notification sent to %s\n", avengerDevice, avengerName)
	}

}

func CheckAvengerDevice(avengerName string) (string, string) {
	return AvengerList[avengerName].AvengerDevice, AvengerList[avengerName].PersonName
}

func CheckAvengerisAvailable(avengerName string) (string, bool) {
	return avengerName, len(AvengerList[avengerName].MissionAssigned) < 2
}
