package api

import (
	"log"
	"time"

	"github.com/katelinlis/BackendMasters/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/katelinlis/BackendMasters/internal/logic"
	"github.com/katelinlis/BackendMasters/internal/model"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/socket.io/v2/socket"
)

func WebSocket(c *gin.Engine, item *model.LobbyList, server map[string]*logic.GameSession) {

	s := socket.DefaultServerOptions()
	s.SetServeClient(true)
	s.SetPingInterval(20 * time.Second) // Интервал отправки Ping
	s.SetPingTimeout(60 * time.Second)  // Тайм-аут до разрыва соединения
	s.SetCors(&types.Cors{              // Кросс-доменные заголовки
		Origin:      "*",
		Credentials: true,
	})

	// Создаем сервер Socket.IO
	io := socket.NewServer(nil, nil)

	// Основной namespace для лобби
	io.On("connection", func(clients ...interface{}) {
		client := clients[0].(*socket.Socket)
		done := make(chan struct{})
		lobby := &model.Lobby{}
		player := &model.PlayerListLobby{}

		// 1. Логика подключения клиента
		log.Println("New client connected:", client.Id())
		client.On("join-lobby", func(args ...interface{}) {

			// Аргумент `lobbyID` от клиента
			lobbyID, _ := uuid.Parse(args[0].(string))
			token, _ := args[1].(string)

			jwt, err := utils.JWTParse(token)
			if err != nil {
				done <- struct{}{}
				client.Emit("lobby-not-found", "")
				return
			}

			lobby = item.GetLobby(lobbyID)
			if lobby == nil {
				done <- struct{}{}
				client.Emit("lobby-not-found", "")
				return
			}

			player = lobby.PlayerList.GetPlayer(jwt.User)

			if player == nil && player.ID == 0 {
				done <- struct{}{}
				client.Emit("player not found", "")
				return
			}

			lobby.BroadcastSendMessage(model.Nottif{Name: player.Name, Message: "connectUser", Who: player.ID})

			player.Online = time.Now()
			player.WSid = string(client.Id())

			// Горутин для отправки уведомлений игроку через его QNotiff
			go func() {
				defer func() {
					select {
					case <-done:
						// Уже закрыто
					default:
						lobby.Lock()
						lobby.BroadcastSendMessage(model.Nottif{Name: player.Name, Message: "disconnectUser", Who: player.ID})

						defer lobby.Unlock()

						close(done) // Закрываем только один раз
					}
				}()

				for {
					select {
					case notiff := <-player.QNotiff: // Получаем новое уведомление
						client.Emit("notification", notiff)
					case <-done:
						log.Printf("Player %d disconnected from lobby %s", player.ID, lobbyID)
						return
					}
				}
			}()

			// Отправляем локальное состояние лобби обратно клиенту
			client.Join(socket.Room(lobbyID.String())) // Клиент присоединяется в комнату по ID
			client.Emit("lobby-state", lobby.PlayerList)
		})

		client.On("ping", func(args ...interface{}) {
			if player != nil && player.ID != 0 {
				// Обновление времени последней активности
				player.Online = time.Now()
				log.Printf("Player %d sent ping", player.ID)
			}

		})

		// 3. Отправка уведомлений при отключении
		client.On("disconnect", func(args ...interface{}) {
			done <- struct{}{}
			lobby.PlayerList = lobby.PlayerList.Delete(player.ID)
			log.Println("Client disconnected:", client.Id())
		})

	})

	c.POST("/socket.io/*f", gin.WrapH(io.ServeHandler(s)))
	c.GET("/socket.io/*f", gin.WrapH(io.ServeHandler(s)))
}
