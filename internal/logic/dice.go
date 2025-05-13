package logic

import (
	"crypto/rand"
	"math/big"

	"github.com/katelinlis/BackendMasters/internal/model"
)

func randRange(min, max int) int {
	// Вычисляем случайное значение с использованием crypto/rand
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err) // Обработка ошибки, если что-то пошло не так
	}
	return int(nBig.Int64()) + min

}

func (g *Game) Dice(user int64, two bool) int {

	if g.Players == nil {
		return 0
	}
	if len(*g.Players) < 2 {
		return 0
	}

	if g.Turn == 0 {
		g.FillDeckVisible()

		g.Turn = user
	}

	if g.Turn != user {
		return 0
	}

	if g.Turn == user && g.TurnStatus == true {
		return 0
	}

	player := g.SearchPlayer(user)
	builds := player.MainBuilds
	pLobby := g.Lobby.PlayerList.GetPlayer(user)

	terminalBuilded := false

	for i := 0; i < len(builds); i++ {
		if builds[i].Name == model.Terminal && builds[i].Builded {
			terminalBuilded = true
		}

	}

	dice := randRange(1, 6)
	dice2 := 0

	if two && terminalBuilded {
		dice2 = randRange(1, 6)
	}

	go g.sendForAllPlayers(model.Nottif{
		Number:  int64(dice),
		Number2: int64(dice2),
		Message: "dice",
		Who:     user,
		Name:    pLobby.Name,
	})
	balance := g.GameLoop(dice+dice2, user)
	g.TurnStatus = true

	balanceHigh := false

	for i := 0; i < len(*g.DeckVisible); i++ {
		deck := (*g.DeckVisible)[i]
		if deck.Cost <= balance {
			balanceHigh = true
		}
	}

	if !balanceHigh {
		g.Next(user) // If Bank Player < build costs
	}

	return dice + dice2
}
