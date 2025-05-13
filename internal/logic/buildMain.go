package logic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/katelinlis/BackendMasters/internal/model"
	"net/http"
)

func (game *Game) BuildMain(context *gin.Context, claims *model.MyCustomClaims, item string) {
	if game == nil {
		context.JSON(http.StatusBadRequest, "")
		return
	}
	player := game.SearchPlayer(claims.User)

	if game.Turn != player.ID && game.TurnStatus != true {
		context.JSON(http.StatusBadRequest, gin.H{"error": errors.New("не ваш ход")})
		return
	}

	if player.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": errors.New("not auth")})
		return
	}

	for i := 0; i < len(player.MainBuilds); i++ {
		build := player.MainBuilds[i]
		if build.Name == item {
			if player.Bank < build.Cost {
				context.JSON(http.StatusBadRequest, gin.H{"error": errors.New("у вас недостаточно денег")})
				return
			}
			player.Lock()
			player.Bank -= build.Cost
			player.MainBuilds[i].Builded = true
			player.Unlock()
		}

	}

	game.Next(player.ID)
}
