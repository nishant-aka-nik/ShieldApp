package controllers

import (
	"fmt"
	"strings"
)

func ListAvengers() {
	fmt.Println("Avenger Name | Status | Assigned Mission")

	for _, avenger := range AvengerList {

		AllAssignedMission := strings.Join(avenger.MissionAssigned, ", ")
		fmt.Println(avenger.PersonName + " | " + avenger.Status + " | " + AllAssignedMission)
	}
}
