package logic

import (
	"sync"
	"testing"

	"github.com/katelinlis/BackendMasters/internal/model"
)

func TestBankOperation(t *testing.T) {
	// Инициализируем банк и игрока
	game := Game{
		Bank:     100,
		gameBank: sync.Mutex{},
	}
	player := model.Player{}
	player.PlayerInit(1)

	// Тест: Успешное выполнение операции
	t.Run("Operation succeeds with sufficient bank funds", func(t *testing.T) {
		success := game.bankOperation(&player, 50)
		if !success {
			t.Errorf("Expected operation to succeed, but it failed")
		}
		if game.Bank != 50 {
			t.Errorf("Expected bank to have 50 money, but got %d", game.Bank)
		}
		if player.Bank != 50 {
			t.Errorf("Expected player to have 50 money, but got %d", player.Bank)
		}
	})

	// Тест: Отказ операции при недостатке средств в банке
	t.Run("Operation fails with insufficient bank funds", func(t *testing.T) {
		success := game.bankOperation(&player, 60)
		if success {
			t.Errorf("Expected operation to fail due to insufficient bank funds, but it succeeded")
		}
		if game.Bank != 50 {
			t.Errorf("Expected bank to remain at 50 money, but got %d", game.Bank)
		}
		if player.Bank != 50 {
			t.Errorf("Expected player to remain at 50 money, but got %d", player.Bank)
		}
	})

	// Тест: Операция на полную сумму доступных средств
	t.Run("Operation succeeds with exact bank amount", func(t *testing.T) {
		success := game.bankOperation(&player, 50)
		if !success {
			t.Errorf("Expected operation to succeed with exact bank amount, but it failed")
		}
		if game.Bank != 0 {
			t.Errorf("Expected bank to be empty, but got %d", game.Bank)
		}
		if player.Bank != 100 {
			t.Errorf("Expected player to have 100 money, but got %d", player.Bank)
		}
	})
}
