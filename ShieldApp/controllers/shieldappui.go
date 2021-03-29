package controllers

import (
	"bufio"
	"fmt"
	"os"
	"shieldapp/models"
	"strconv"
)

//Map to store reference of mission struct
var AllMission = make(map[string]*models.Mission)

//Map to store reference of avengerdatabase struct
var AvengerList = make(map[string]*models.AvengersDatabase)

//Display UI
func ShieldEmployee(Name string, Role string) {
	fmt.Printf("========Welcome %v===========\n", Name)

	i := 0
	for i == 0 {

		//As of now the app is only design for nick fury as asked
		if Role == "Director" {
			fmt.Println("----------xxxxxxxxx-------------")
			fmt.Println("1. Check the missions")
			fmt.Println("2. Assign mission to Avengers")
			fmt.Println("3. Check mission’s details")
			fmt.Println("4. Check Avenger’s details")
			fmt.Println("5. Update Mission’s status")
			fmt.Println("6. List Avengers")
			fmt.Println("7. Logout")
			fmt.Println("----------xxxxxxxxx-------------")

		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println("Enter your choice :- ")
		action, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Println("----------xxxxxxxxx-------------")
		fmt.Printf("you entered %d\n", action)
		fmt.Println("----------xxxxxxxxx-------------")

		switch action {
		case 1:
			//function to check all the mission details
			CheckMission()
		case 2:
			//function to assign mission to avengers
			AssignMission()
		case 3:
			//function to check details of a mission
			CheckMissionDetails()
		case 4:
			//function to check details of a avenger
			CheckAvengerDetails()
		case 5:
			//function to update details of a mission
			UpdateMission()
		case 6:
			//function to check all the avenger status
			ListAvengers()
		default:
			fmt.Println("Invalid input")
		}

		if action == 7 {
			break
		}

	}
}
