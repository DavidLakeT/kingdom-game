package main

import "fmt"

func main() {

	fmt.Println("(------------- ELF KINGDOM INFO -------------)")
	var elfKingdom = createKingdom("elf")
	fmt.Printf("KING DESCRIPTION: %s\n", elfKingdom.king.getDescription())
	fmt.Printf("CASTLE DESCRIPTION: %s\n", elfKingdom.castle.getDescription())
	fmt.Printf("ARMY DESCRIPTION: %s\n", elfKingdom.army.getDescription())

	fmt.Println()

	fmt.Println("(------------- ELF KINGDOM INFO -------------)")
	var orcKingdom = createKingdom("orc")
	fmt.Printf("KING DESCRIPTION: %s\n", orcKingdom.king.getDescription())
	fmt.Printf("CASTLE DESCRIPTION: %s\n", orcKingdom.castle.getDescription())
	fmt.Printf("ARMY DESCRIPTION: %s\n", orcKingdom.army.getDescription())
}

func createKingdom(kingdomtype string) Kingdom {

	var kingdomFactory = MakeFactory(kingdomtype)

	return Kingdom{
		king:   kingdomFactory.createKing(),
		castle: kingdomFactory.createCastle(),
		army:   kingdomFactory.createArmy(),
	}
}
