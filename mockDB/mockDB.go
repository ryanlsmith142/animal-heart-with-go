package mockDB

import "../models/RescueAnimal"

var RescueAnimals []RescueAnimal

func InitializeRescueAnimalsArray() {
	RescueAnimals = []RescueAnimal {
		RescueAnimal {Id: "1", Name: "Jerry", Type: "Dog"},
		RescueAnimal {Id: "2", Name: "Ryan", Type: "Cat"},
	}
}
