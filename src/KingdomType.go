package main

type KingdomType interface{}

type Kingdom struct {
	king   King
	castle Castle
	army   Army
}

const (
	ELF string = "elf"
	ORC string = "orc"
)

func MakeFactory(kingdomType string) KingdomFactory {

	if kingdomType == ELF {

		return &ElfKingdomFactory{}
	}

	if kingdomType == ORC {

		return &OrcKingdomFactory{}
	}

	return nil
}
