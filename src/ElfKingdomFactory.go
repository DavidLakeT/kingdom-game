package main

type ElfKingdomFactory struct {
	KingdomFactory
}

func (kingdomFactory *ElfKingdomFactory) createCastle() Castle {

	return &ElfCastle{}
}

func (kingdomFactory *ElfKingdomFactory) createKing() King {

	return &ElfKing{}
}

func (kingdomFactory *ElfKingdomFactory) createArmy() Army {

	return &ElfArmy{}
}
