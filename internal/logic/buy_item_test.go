package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/katelinlis/BackendMasters/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuyItem(t *testing.T) {
	gin.SetMode(gin.TestMode) // Установка тестового режима для Gin

	t.Run("Player tries to buy item when it's not their turn", func(t *testing.T) {
		// Arrange
		game := initMockLobby(2, []string{})
		game.Turn = 2 // Устанавливаем ход за другого игрока

		// Устанавливаем HTTP-запрос с тестовым контекстом
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/buy-item?item=Card1", nil)

		// Act
		game.BuyItem(c, 1)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "не ваш ход")
	})

	t.Run("Player tries to buy item without rolling dice", func(t *testing.T) {
		// Arrange
		game := initMockLobby(2, []string{})
		game.Turn = 1
		game.TurnStatus = false // Кости не были брошены

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/buy-item?item=Card1", nil)

		// Act
		game.BuyItem(c, 1)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "сначала надо кинуть кубик")
	})

	t.Run("Player buys item successfully", func(t *testing.T) {
		// Arrange
		game := initMockLobby(2, []string{})
		game.Turn = 1
		game.TurnStatus = true // Кости брошены

		player := &(*game.Players)[0]
		(*player).Bank = 10 // Игрок имеет достаточно денег

		game.DeckVisible = &[]model.Card{
			{
				Name: "Card1",
				Cost: 5,
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/buy-item?item=Card1", nil)

		// Act
		game.BuyItem(c, 1)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"status":"ok"`)
		assert.Equal(t, 5, int((*player).Bank))            // Убедитесь, что банк уменьшен
		assert.Len(t, (*player).Builds, 3)                 // Проверяем, что покупка добавлена
		assert.Equal(t, "Card1", (*player).Builds[2].Name) // Проверяем имя здания
	})

	t.Run("Player tries to buy item they can't afford", func(t *testing.T) {
		// Arrange
		game := initMockLobby(2, []string{})
		game.Turn = 1
		game.TurnStatus = true

		player := &(*game.Players)[0]
		(*player).Bank = 2 // У игрока недостаточно денег

		game.DeckVisible = &[]model.Card{
			{
				Name: "Card2",
				Cost: 5,
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/buy-item?item=Card2", nil)

		// Act
		game.BuyItem(c, 1)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)                // Запрос обрабатывается успешно
		assert.Equal(t, 2, int((*player).Bank))               // Банк остаётся неизменным
		assert.Len(t, (*player).Builds, 2)                    // Здания не добавлены
		assert.Len(t, *game.DeckVisible, 1)                   // Карта остаётся видимой
		assert.Equal(t, "Card2", (*game.DeckVisible)[0].Name) // Проверяем имя карты
	})

	t.Run("Valid request but item not in deck", func(t *testing.T) {
		// Arrange
		game := initMockLobby(2, []string{})
		game.Turn = 1
		game.TurnStatus = true

		player := &(*game.Players)[0]
		(*player).Bank = 10 // Игрок имеет достаточно денег

		game.DeckVisible = &[]model.Card{
			{
				Name: "Card3",
				Cost: 5,
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/buy-item?item=NonExistentCard", nil)

		// Act
		game.BuyItem(c, 1)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)                // Запрос обработан
		assert.Equal(t, 10, int((*player).Bank))              // Банк не меняется
		assert.Len(t, (*player).Builds, 2)                    // Здания не добавлены
		assert.Len(t, *game.DeckVisible, 1)                   // Дек остаётся без изменений
		assert.Equal(t, "Card3", (*game.DeckVisible)[0].Name) // Проверяем имя оставшейся карты
	})
}
