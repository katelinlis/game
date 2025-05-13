package internal

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katelinlis/BackendMasters/internal/api"
	"github.com/katelinlis/BackendMasters/internal/logic"
	"github.com/katelinlis/BackendMasters/internal/model"
	"github.com/katelinlis/BackendMasters/internal/utils"

	"net/http"

	"github.com/katelinlis/BackendMasters/internal/store"
)

type Server struct {
	router *gin.Engine
	store  store.Store
	Game   map[string]*logic.GameSession
	lobby  *model.LobbyList
}

func Init() Server {

	server := Server{
		router: gin.Default(),
		Game:   make(map[string]*logic.GameSession),
		lobby:  &model.LobbyList{},
	}

	server.router.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.1.49:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "auth", "Content-Type", "Connection"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if origin == "http://localhost:3000" ||
				origin == "http://192.168.1.49:3000" ||
				origin == "http://localhost" ||
				origin == "https://game.katelinlis.com" {
				return true
			}

			println(origin)

			return origin == "chrome-extension://ophmdkgfcjapomjdpfobjfbihojchbko"
		},
		MaxAge: 12 * time.Hour,
	}))

	apiRouter := server.router.Group("/api")

	api.Lobby(apiRouter.Group("lobby"), server.lobby)

	go server.lobby.DeletesLobby()

	apiRouter.POST("init/:id", func(context *gin.Context) {
		id, _ := context.Params.Get("id")
		localID, _ := uuid.Parse(id)

		_, user := utils.AuthAndGetGame(context, nil)

		for _, item := range (*server.lobby).List {
			if item.ID == localID {
				if item.PlayerList[0].ID == user.User {

					if len(item.PlayerList) > 1 {
						readyCount := 0
						for _, player := range item.PlayerList {
							if player.Ready {
								readyCount++
							}

						}

						if readyCount != len(item.PlayerList) {
							context.JSON(http.StatusUnprocessableEntity, "не все готовы")
							return
						}

					}

					if _, exist := server.Game[localID.String()]; !exist {

						server.Game[localID.String()] = &logic.GameSession{
							Game:    logic.GameInit(item),
							LastUse: time.Now(),
						}

						item.BroadcastSendMessage(model.Nottif{
							Message: "Start game",
						})

						item.GameStarted = true
						context.JSON(http.StatusOK, "запушен")
						return

					} else {
						item.BroadcastSendMessage(model.Nottif{
							Message: "Start game",
						})
						context.JSON(http.StatusOK, "запушен")
					}
				}
			}
		}
		context.JSON(http.StatusUnprocessableEntity, "не правомерный доступ")

	})

	api.WebSocket(server.router, server.lobby, server.Game)

	apiRouter.GET("players/:id", func(context *gin.Context) {
		game, _ := utils.AuthAndGetGame(context, server.Game)

		if game == nil {
			context.JSON(http.StatusBadRequest, "")
		}
		context.JSON(http.StatusOK, game.Players)

	})

	apiRouter.GET("deck/:id", func(context *gin.Context) {
		game, _ := utils.AuthAndGetGame(context, server.Game)

		if game == nil {
			context.JSON(http.StatusBadRequest, "")
			return
		}

		context.JSON(http.StatusOK, game.DeckVisible)

	})

	apiRouter.POST("buyShop/:id", func(context *gin.Context) {
		game, user := utils.AuthAndGetGame(context, server.Game)

		if game == nil {
			context.JSON(http.StatusBadRequest, "")
		}

		game.BuyItem(context, user.User)

	})

	apiRouter.POST("build/:id", func(context *gin.Context) {
		item := context.Query("item")

		game, claims := utils.AuthAndGetGame(context, server.Game)

		game.BuildMain(context, claims, item)

	})

	apiRouter.GET("getUser/:id", func(context *gin.Context) {
		game, id := utils.AuthAndGetGame(context, server.Game)

		if game == nil {
			context.JSON(http.StatusBadRequest, "")
			return
		}
		player := game.SearchPlayer(id.User)
		context.JSON(http.StatusOK, &player)
	})

	apiRouter.POST("dice/:id", func(context *gin.Context) {
		game, id := utils.AuthAndGetGame(context, server.Game)

		two := context.Request.URL.Query().Has("two")

		if game == nil {
			context.JSON(http.StatusBadRequest, "")
			return
		}

		dices := game.Dice(id.User, two)
		if dices == 0 {
			context.JSON(http.StatusBadRequest, dices)
			return
		}

		context.JSON(http.StatusOK, dices)
	})

	return server
}

func (s Server) Start() {

	s.router.Run(":8080")
}
