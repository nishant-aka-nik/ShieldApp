package main

import (
	"bufio"
	"fmt"
	"os"
	"shieldapp/controllers"
	"shieldapp/models"
)

func init() {
	controllers.InitAvengerDatabase()
}

func main() {

	//Adding login functionality for future enhancements if some other avengers also given access to app
	var credential []models.Credentials
	furycredential := models.Credentials{
		Name:     "Nick Fury",
		Username: "Nick", //username to login to app
		Password: "Fury", //password to login to app
		Role:     "Director",
	}

	credential = append(credential, furycredential)

	scanner := bufio.NewScanner(os.Stdin)

	i := 0
	for i == 0 {
		fmt.Println("=======------S.H.I.E.L.D ------=========")
		fmt.Println("Enter your username :- ")
		scanner.Scan()
		username := scanner.Text()
		fmt.Println("Enter your password :- ")
		scanner.Scan()
		password := scanner.Text()



		if credential[0].Username == username && credential[0].Password == password {
			controllers.ShieldEmployee(credential[0].Name, credential[0].Role)
			break
		}
		fmt.Println("Sorry wrong login credentials try again")
		fmt.Println("=======================================")

	}

}

