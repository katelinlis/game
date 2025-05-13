package logic

import (
	"testing"
)

func TestGameNextExtended(t *testing.T) {
	// Тест 1: Пустой список игроков
	t.Run("Empty Players List", func(t *testing.T) {
		game := initMockLobby(0, []string{})
		game.Players = nil

		game.Next(1)
		if game.Turn != 0 {
			t.Errorf("Expected Turn to remain 0 for empty players list, got %d", game.Turn)
		}
	})

	// Тест 2: Один игрок
	t.Run("Single Player", func(t *testing.T) {
		game := initMockLobby(1, []string{})
		game.Next(1) // Единственный игрок, ход остаётся за ним.
		if game.Turn != 1 {
			t.Errorf("Expected Turn to remain 1 for single player, got %d", game.Turn)
		}
	})

	// Тест 3: Два игрока
	t.Run("Two Players", func(t *testing.T) {
		game := initMockLobby(2, []string{})
		// Ход переходит от первого ко второму
		game.Next(1)
		if game.Turn != 2 {
			t.Errorf("Expected Turn to be 2, got %d", game.Turn)
		}
		// Ход переходит от второго к первому
		game.Next(2)
		if game.Turn != 1 {
			t.Errorf("Expected Turn to be 1, got %d", game.Turn)
		}
	})

	// Тест 4: Три игрока
	t.Run("Three Players", func(t *testing.T) {
		game := initMockLobby(3, []string{})
		// Ход переходит от первого ко второму
		game.Next(1)
		if game.Turn != 2 {
			t.Errorf("Expected Turn to be 2, got %d", game.Turn)
		}
		// Ход переходит от второго к третьему
		game.Next(2)
		if game.Turn != 3 {
			t.Errorf("Expected Turn to be 3, got %d", game.Turn)
		}
		// Ход переходит от третьего к первому (цикличность)
		game.Next(3)
		if game.Turn != 1 {
			t.Errorf("Expected Turn to be 1, got %d", game.Turn)
		}
	})

	// Тест 5: Неверный текущий пользователь
	t.Run("Invalid Current User", func(t *testing.T) {
		game := initMockLobby(3, []string{})
		// Передаём ID, которого нет в списке игроков
		game.Next(999)
		if game.Turn != 0 {
			t.Errorf("Expected Turn to remain 0 when invalid user provided, got %d", game.Turn)
		}
	})

	// Тест 6: Несортированный список игроков
	t.Run("Unordered Players", func(t *testing.T) {
		game := initMockLobby(3, []string{})
		// Ход переходит от первого игрока ко второму
		game.Next(3)
		if game.Turn != 1 {
			t.Errorf("Expected Turn to be 1, got %d", game.Turn)
		}
		// Ход переходит от второго к третьему
		game.Next(2)
		if game.Turn != 3 {
			t.Errorf("Expected Turn to be 3, got %d", game.Turn)
		}
		// Ход переходит от третьего к первому (цикличность)
		game.Next(1)
		if game.Turn != 2 {
			t.Errorf("Expected Turn to be 2, got %d", game.Turn)
		}
	})
}
