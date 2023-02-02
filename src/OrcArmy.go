package main

type OrcArmy struct {
	Army
}

func (army *OrcArmy) getDescription() string {

	var description = "Orc Army Description :)"

	return description
}
