package main

type OrcKingdomFactory struct {
	KingdomFactory
}

func (kingdomFactory *OrcKingdomFactory) createCastle() Castle {

	return &OrcCastle{}
}

func (kingdomFactory *OrcKingdomFactory) createKing() King {

	return &OrcKing{}
}

func (kingdomFactory *OrcKingdomFactory) createArmy() Army {

	return &OrcArmy{}
}
