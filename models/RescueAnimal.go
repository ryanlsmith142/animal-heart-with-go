package models

type RescueAnimal struct {
	Id string
	Name string
	Type string
}

func (rescueAnimal *RescueAnimal) GetId() string {
	return rescueAnimal.Id
}

func (rescueAnimal *RescueAnimal) GetName() string {
	return rescueAnimal.Name
}

func (rescueAnimal *RescueAnimal) GetType() string {
	return rescueAnimal.Type
}