package perhitung

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(namaGame string) {
	gamer.Games = append(gamer.Games, namaGame)
}
