package controllers

import (
	"fmt"
	"strings"
)

func CheckMission() {

	if len(AllMission) == 0 {
		fmt.Println("No Mission has been assigned to an Avenger")
	} else {

		fmt.Println("Mission Name -----|-----Status-----|--------- Avengers---------")

		for _, eachMission := range AllMission {

			AllAssignedAvenger := strings.Join(eachMission.AvengersOnMisson, ", ")

			fmt.Println(eachMission.MissionName + "-----|-----" + eachMission.MissionStatus + "-----|---------" + AllAssignedAvenger)

		}

		fmt.Println("=========End of missions=================")
	}

}
