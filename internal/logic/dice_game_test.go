package logic

import (
	"fmt"
	"github.com/katelinlis/BackendMasters/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initMockLobby(players int, builds []string) *Game {
	lobby := model.NewLobby(
		"title",
		"username",
		1)
	players--
	for i := 1; i <= players; i++ {
		lobby.AddToLobby("user2", int64(i+1))
	}

	game := GameInit(lobby)

	for _, item := range *game.Players {
		buildMap := make(map[string]bool)
		for _, buildName := range builds {
			buildMap[buildName] = true
		}

		for index, build := range (*item).MainBuilds {
			if buildMap[build.Name] {
				build.Builded = true
				(*item).MainBuilds[index] = build
			}
		}

	}

	game.FillDeckVisible()

	return game
}

func TestDice(t *testing.T) {
	t.Run("No players exist", func(t *testing.T) {
		// Arrange
		game := &Game{}

		// Act
		result := game.Dice(1, false)

		// Assert
		assert.Equal(t, 0, result, "Dice function should return 0 when no players exist.")
	})

	t.Run("Less than 2 players", func(t *testing.T) {
		// Arrange

		game := initMockLobby(1, []string{model.Terminal})

		// Act
		result := game.Dice(1, false)

		// Assert
		assert.Equal(t, 0, result, "Dice function should return 0 when there are less than 2 players.")
	})

	t.Run("Is turn Zero", func(t *testing.T) {
		// Arrange

		game := initMockLobby(2, []string{model.Terminal})

		// Act
		result := game.Dice(1, false)

		// Assert
		assert.NotEqual(t, 0, result, "not 0 if turn is zero")
	})

	t.Run("Invalid user turn", func(t *testing.T) {
		// Arrange
		game := initMockLobby(1, []string{model.Terminal})
		game.Turn = 2

		// Act
		result := game.Dice(1, false)

		// Assert
		assert.Equal(t, 0, result, "Dice function should return 0 when it's not the user's turn.")
	})

	t.Run("Valid single dice roll", func(t *testing.T) {
		// Arrange
		game := initMockLobby(2, []string{model.Terminal})
		game.Turn = 1

		// Act
		result := game.Dice(1, false)

		// Assert
		assert.NotEqual(t, 0, result, "Dice function should return a non-zero value for a valid single dice roll.")
	})

	t.Run("Valid rolls", func(t *testing.T) {

		const rolls = 1000000
		counts := make(map[int]int)

		for i := 0; i < rolls; i++ {
			num := randRange(1, 6)
			counts[num]++
		}

		// Вывод статистики вероятностей
		for i := 1; i <= 6; i++ {
			fmt.Printf("Значение %d выпало %.2f%% времени\n", i, float64(counts[i])/float64(rolls)*100)
		}

	})

	t.Run("Valid two dice roll", func(t *testing.T) {
		// Arrange
		game := initMockLobby(2, []string{model.Terminal})
		game.Turn = 1
		result := 0

		for i := 0; i < 50; i++ {
			// Act
			result2 := game.Dice(game.Turn, true)
			game.Next(game.Turn)
			t.Log(result2)
			if result2 > result {
				result = result2

			}

		}

		// Assert
		t.Log(result)
		assert.GreaterOrEqual(t, 12, result, "Dice function should return greater or equal of 12")
	})

	t.Run("Not user's turn", func(t *testing.T) {
		// Arrange
		game := initMockLobby(1, []string{model.Terminal})
		game.Turn = 2

		// Act
		result := game.Dice(1, false)

		// Assert
		assert.Equal(t, 0, result, "Dice function should return 0 if it's not the user's turn.")
	})

	t.Run("Turn has already been played", func(t *testing.T) {
		// Arrange
		game := initMockLobby(1, []string{model.Terminal})
		game.Turn = 1
		game.TurnStatus = true

		// Act
		result := game.Dice(1, false)

		// Assert
		assert.Equal(t, 0, result, "Dice function should return 0 if the turn has already been played.")
	})
}
