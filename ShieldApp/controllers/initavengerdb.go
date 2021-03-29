package controllers

import (
	"shieldapp/models"
	"strings"
)

func InitAvengerDatabase() {

	avengers := []models.AvengersDatabase{
		{
			PersonName:    "Iron Man",
			Abilities:     "Highly Intelligent Suit of Armor",
			Status:        "Available",
			AvengerDevice: "Email",
		},
		{
			PersonName:    "Captain America",
			Abilities:     "Highly Intelligent american Shield",
			Status:        "Available",
			AvengerDevice: "SMS",
		},
		{
			PersonName:    "Black Widow",
			Abilities:     "Highly Intelligent women power",
			Status:        "Available",
			AvengerDevice: "SMS",
		},
		{
			PersonName:    "Hulk",
			Abilities:     "green muscles power",
			Status:        "Available",
			AvengerDevice: "Email",
		},
		{
			PersonName:    "Thor",
			Abilities:     "Mjolnir",
			Status:        "Available",
			AvengerDevice: "Pager",
		},
		{
			PersonName:    "Hawk Eye",
			Abilities:     "Arrows",
			Status:        "Available",
			AvengerDevice: "SMS",
		},
	}

	i := 0
	for _, v := range avengers {
		AvengerList[strings.ToLower(v.PersonName)] = &avengers[i]
		i++
	}

}
