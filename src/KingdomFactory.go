package main

type KingdomFactory interface {
	createCastle() Castle
	createKing() King
	createArmy() Army
}
