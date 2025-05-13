package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katelinlis/BackendMasters/internal/model"
)

func (game *Game) BuyItem(context *gin.Context, id int64) {

	player := &model.Player{}

	if game.Turn != id {
		context.JSON(http.StatusBadRequest, gin.H{"error": "не ваш ход", "status": ""})
		return
	}

	if game.TurnStatus != true {
		context.JSON(http.StatusBadRequest, gin.H{"error": "сначала надо кинуть кубик", "status": ""})
		return
	}

	for i := 0; i < len(*game.Players); i++ {
		if (*game.Players)[i].ID == id {
			player = (*game.Players)[i]
			break
		}
	}

	item := context.Query("item")

	for i := 0; i < len(*game.DeckVisible); i++ {
		deck := (*game.DeckVisible)[i]
		if deck.Name == item {

			if deck.Cost <= player.Bank {
				*game.DeckVisible = remove(*game.DeckVisible, i)
				player.Builds = append(player.Builds, deck)

				player.Bank -= deck.Cost

				game.FillDeckVisible()
			}

			break
		}
	}

	game.Next(id)

	context.JSON(http.StatusOK, gin.H{"status": "ok", "error": ""})

}

func remove(slice []model.Card, s int) []model.Card {
	return append(slice[:s], slice[s+1:]...)
}
