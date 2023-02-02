package main

type OrcKing struct {
	King
}

func (king *OrcKing) getDescription() string {

	var description = "Orc King Description :)"

	return description
}
