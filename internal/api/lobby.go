package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	consts "github.com/katelinlis/BackendMasters/internal/const"
	"github.com/katelinlis/BackendMasters/internal/model"
	"github.com/katelinlis/BackendMasters/internal/utils"
	"net/http"
	"reflect"
)

func Copy(source interface{}, destin interface{}) {
	x := reflect.ValueOf(source)
	if x.Kind() == reflect.Ptr {
		starX := x.Elem()
		y := reflect.New(starX.Type())
		starY := y.Elem()
		starY.Set(starX)
		reflect.ValueOf(destin).Elem().Set(y.Elem())
	} else {
		destin = x.Interface()
	}
}

func Lobby(group *gin.RouterGroup, lobby *model.LobbyList) {
	group.GET("", func(context *gin.Context) {
		var lobbyList []model.LobbyPublic

		for _, item := range (*lobby).List {

			if !item.GameStarted {
				newLobby := model.LobbyPublic{
					ID:         item.ID,
					Name:       item.Name,
					PlayerList: item.PlayerList,
					Created:    item.Created,
				}
				lobbyList = append(lobbyList, newLobby)
			}

		}

		context.JSON(http.StatusOK, lobbyList)
	})
	group.POST("", func(context *gin.Context) {
		type CreateLobby struct {
			Title    string `json:"title"`
			Username string `json:"username"`
		}

		change := CreateLobby{}

		if err := json.NewDecoder(context.Request.Body).Decode(&change); err != nil {
			context.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}

		_, user := utils.AuthAndGetGame(context, nil)
		if user == nil {
			context.JSON(http.StatusForbidden, "")
			return
		}

		lobbyNew := model.NewLobby(change.Title, user.Username, user.User)

		lobby.AddNewLobby(lobbyNew)

		context.JSON(http.StatusOK, lobbyNew)
	})

	group.GET("/:id", func(context *gin.Context) {
		id := context.Param("id")
		localID, _ := uuid.Parse(id)
		_, user := utils.AuthAndGetGame(context, nil)
		if user == nil {
			context.JSON(http.StatusForbidden, "")
			return
		}

		item := lobby.GetLobby(localID)

		if item == nil {
			context.JSON(http.StatusNotFound, "")
		}

		for _, player := range item.PlayerList {

			if player.ID == user.User {
				context.JSON(http.StatusOK, item)
				return
			}
		}
		context.JSON(http.StatusNotFound, "")

	})

	group.POST("/auth", func(ctx *gin.Context) {

		user, ok := bot.ValidateWebappRequest(ctx.Request.URL.Query(), consts.TG_API)

		if !ok {
			ctx.JSON(http.StatusForbidden, "incorrect request")
			return
		}

		t := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{

				"user":     user.ID,
				"username": user.Username,
			})
		s, err := t.SignedString([]byte(consts.JwtKEY))
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		ctx.JSON(http.StatusOK, s)
	})

	group.POST("/:id/ready", func(context *gin.Context) {
		id := context.Param("id")
		localID, _ := uuid.Parse(id)
		_, user := utils.AuthAndGetGame(context, nil)
		if user == nil {
			context.JSON(http.StatusForbidden, "")
			return
		}

		item := lobby.GetLobby(localID)

		if item == nil {
			context.JSON(http.StatusNotFound, "")
		}
		defer item.Unlock()
		item.Lock()
		for _, player := range item.PlayerList {
			if player.ID == user.User {

				player.Ready = !player.Ready

				message := "not ready"
				if player.Ready {
					message = "ready"
				}

				item.BroadcastSendMessage(model.Nottif{
					Message: message,
					Who:     player.ID,
				})

				context.JSON(http.StatusOK, item)
				return
			}

		}

		context.JSON(http.StatusNotFound, "")

	})

	group.PUT("/:id", func(context *gin.Context) {
		id := context.Param("id")

		localID, err := uuid.Parse(id)
		if err != nil {
			context.JSON(http.StatusUnprocessableEntity, "")
			return
		}

		_, user := utils.AuthAndGetGame(context, nil)
		if user == nil {
			context.JSON(http.StatusForbidden, "")
			return
		}

		type JoinLobby struct {
			Username string `json:"username"`
		}

		change := JoinLobby{}

		if err := json.NewDecoder(context.Request.Body).Decode(&change); err != nil {
			context.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}

		item := lobby.GetLobby(localID)

		if item == nil {
			context.JSON(http.StatusNotFound, "")
		}

		if len(item.PlayerList) > 4 {
			context.JSON(http.StatusUnprocessableEntity, "лобби полное")
			return
		}
		item.AddToLobby(change.Username, user.User)
		context.JSON(http.StatusOK, item)

	})
}
