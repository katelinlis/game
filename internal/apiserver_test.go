package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/katelinlis/BackendMasters/internal/logic"
	"github.com/katelinlis/BackendMasters/internal/model"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"
)

func TestCreateLobby(t *testing.T) {
	lobby := model.LobbyList{}
	lobbyNew := model.NewLobby("title", "username", 1)

	lobby.AddNewLobby(lobbyNew)
	find := lobby.GetLobby(lobbyNew.ID)

	find.AddToLobby("user2", 2)

	t.Log(lobbyNew.ID)
	//assert.Equal(t, lobbyNew, find)
}

func TestInitGame(t *testing.T) {
	lobby := model.LobbyList{}
	lobbyNew := model.NewLobby("title", "username", 1)

	lobby.AddNewLobby(lobbyNew)
	find := lobby.GetLobby(lobbyNew.ID)

	find.AddToLobby("user2", 2)

	go func() {
		for {
			for i := 0; i < len(find.PlayerList); i++ {
				find.PlayerList.GetPlayer(int64(i) + 1).Online = time.Now()
			}
			time.Sleep(time.Second * 10)
		}
	}()

	game := logic.GameInit(find)

	for i := 0; i < 100; i++ {
		turn := game.Turn
		if turn == 0 {
			turn = int64(rand.Intn(2) + 1)
		}
		game.Dice(turn, false)
		t.Log(game.Bank)

		if game.TurnStatus == true {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			for _, item := range *game.DeckVisible {
				if item.Cost <= game.SearchPlayer(game.Turn).Bank {
					//time.Sleep(time.Second / 2)
					encodedName := url.QueryEscape(item.Name)

					c.Request = httptest.NewRequest(http.MethodPost, "/buy-item?item="+encodedName, nil)
					break
				}
			}

			// Act
			game.BuyItem(c, turn)
			t.Log(len(game.SearchPlayer(game.Turn).Builds), game.Turn)
			assert.Equal(t, http.StatusOK, w.Code)
		} else {
			game.Next(turn)
		}
	}

	t.Log(game)
	t.Log(lobbyNew.ID)
	//assert.Equal(t, lobbyNew, find)
}

func BenchmarkGameSessions(b *testing.B) {
	// Количество сессий
	const sessionsCount = 10

	// Создаем общий список лобби
	lobby := model.LobbyList{}
	var wg sync.WaitGroup

	// Инициализируем N игровых сессий
	for s := 0; s < sessionsCount; s++ {
		wg.Add(1)
		b.Logf(
			"Starting session %d. Lobby size: %d",
			s+1,
			len(lobby.List),
		)

		go func(sessionID int) {
			defer wg.Done()

			// Создаем новое лобби
			lobbyNew := model.NewLobby("title", "username", 1)
			lobby.AddNewLobby(lobbyNew)

			// Добавляем еще одного игрока
			find := lobby.GetLobby(lobbyNew.ID)
			find.AddToLobby("user2", 2)

			// Постоянно обновляем статус игрока "онлайн"
			go func() {
				for {
					for i := 0; i < len(find.PlayerList); i++ {
						find.PlayerList.GetPlayer(int64(i) + 1).Online = time.Now()
					}
					time.Sleep(time.Second * 10)
				}
			}()

			// Инициализация игры
			game := logic.GameInit(find)

			// Запускаем игровую сессию
			for i := 0; i < 100; i++ {
				turn := game.Turn
				if turn == 0 {
					// Генерация случайного игрока (1 или 2)
					turn = int64(rand.Intn(2) + 1)
				}
				game.Dice(turn, false)

				// Логика совершения хода игрока
				if game.TurnStatus == true {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)

					// Игрок пытается купить доступный предмет
					for _, item := range *game.DeckVisible {
						if item.Cost <= game.SearchPlayer(game.Turn).Bank {
							encodedName := url.QueryEscape(item.Name)
							c.Request = httptest.NewRequest("POST", "/buy-item?item="+encodedName, nil)
							break
						}
					}

					// Выполняем покупку
					game.BuyItem(c, turn)
					assert.Equal(b, http.StatusOK, w.Code)
				} else {
					// Переход хода
					game.Next(turn)
				}
			}

			// Вывод результатов сессии
			b.Logf("Session %d finished. Game bank: %d", sessionID, game.Bank)
		}(s)
	}

	// Ждем завершения всех сессий
	wg.Wait()
}
