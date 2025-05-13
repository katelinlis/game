package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	consts "github.com/katelinlis/BackendMasters/internal/const"
	"github.com/katelinlis/BackendMasters/internal/logic"
	"github.com/katelinlis/BackendMasters/internal/model"
)

func JWTParse(tokenString string) (MainClaims *model.MyCustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKEY), nil
	})

	if err != nil {
		return MainClaims, err
	}

	if claims, ok := token.Claims.(*model.MyCustomClaims); ok {
		return claims, nil
	}
	return MainClaims, errors.New("invalid token")
}

func AuthAndGetGame(context *gin.Context, gameFind map[string]*logic.GameSession) (game *logic.Game, claims *model.MyCustomClaims) {
	tokenString := context.GetHeader("auth")

	claims, err := JWTParse(tokenString)
	if err != nil {
		return nil, nil
	}

	if gameFind != nil {
		gameId, _ := context.Params.Get("id")
		localID, _ := uuid.Parse(gameId)

		if game, exist := gameFind[localID.String()]; exist {
			return game.Game, claims
		}
	}

	return game, claims

}
