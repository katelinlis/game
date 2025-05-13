package logic

import (
	"github.com/katelinlis/BackendMasters/internal/model"
	"testing"
	"time"
)

func TestRedCardsBasic(t *testing.T) {
	// Подготовка игроков с красными картами
	player1 := &model.Player{ID: 1, Bank: 10} // Игрок с картой
	player2 := &model.Player{ID: 2, Bank: 20} // Игрок, с которого "снимают" монеты

	players := []*model.Player{player1, player2}

	redCard := model.Card{
		Name:  "RedCardTest",
		Color: "red",
		Rules: model.BuildRules{
			When:     "dice",
			Value:    5, // Карта активируется при броске 5
			ValueAdd: 2, // Забрать "2 монеты" у другого игрока
			Name:     "player_bank",
		},
	}

	// Добавляем красную карту игроку 1
	player1.Builds = []model.Card{redCard}

	// Формируем игровую сессию
	game := &Game{Players: &players, Bank: 100, queueUserBalance: make(chan BalancedUserAdd, 10)}

	game.InitQueueBalance()

	// Симулируем бросок кубика со значением 5 для игрока 2
	diceResult := 5
	game.GameLoop(diceResult, 2)
	time.Sleep(time.Second / 1000)
	// Проверяем, что баланс игрока 2 уменьшился
	if player2.Bank != 18 {
		t.Errorf("Expected player 2 balance to be 18, got %d", player2.Bank)
	}

	// Проверяем, что баланс игрока 1 увеличился
	if player1.Bank != 12 {
		t.Errorf("Expected player 1 balance to be 12, got %d", player1.Bank)
	}
}

func TestRedCardsNoEffect(t *testing.T) {
	// Подготовка игроков с красной картой, но без выполнения условия
	player1 := &model.Player{ID: 1, Bank: 10}
	player2 := &model.Player{ID: 2, Bank: 20}

	players := []*model.Player{player1, player2}

	redCard := model.Card{
		Name:  "RedCardTest",
		Color: "red",
		Rules: model.BuildRules{
			When:     "dice",
			Value:    5, // Активируется только при значении кубика 5
			ValueAdd: 2, // Забрать "2 монеты" у другого игрока
			Name:     "player_bank",
		},
	}

	player1.Builds = []model.Card{redCard}

	game := &Game{Players: &players, Bank: 100}

	// Бросок кубика со значением 4 (карта не активируется)
	diceResult := 4
	game.GameLoop(diceResult, 2)

	// Балансы игроков не должны измениться
	if player2.Bank != 20 {
		t.Errorf("Expected player 2 balance to remain 20, got %d", player2.Bank)
	}
	if player1.Bank != 10 {
		t.Errorf("Expected player 1 balance to remain 10, got %d", player1.Bank)
	}
}

func TestRedCardsMultiplePlayers(t *testing.T) {
	// Подготовка игроков с красной картой
	player1 := &model.Player{ID: 1, Bank: 5}  // Игрок с картой
	player2 := &model.Player{ID: 2, Bank: 10} // Жертва карт
	player3 := &model.Player{ID: 3, Bank: 15} // Без изменений

	players := []*model.Player{player1, player2, player3}

	redCard := model.Card{
		Name:  "RedCardTest",
		Color: "red",
		Rules: model.BuildRules{
			When:     "dice",
			Value:    6, // Активируется при значении кубика 6
			ValueAdd: 3, // Забрать "3 монеты" у другого игрока
			Name:     "player_bank",
		},
	}

	// Добавляем красную карту игроку 1
	player1.Builds = []model.Card{redCard}

	game := &Game{Players: &players, Bank: 100, queueUserBalance: make(chan BalancedUserAdd, 10)}

	game.InitQueueBalance()

	// Симулируем бросок кубика со значением 6
	diceResult := 6

	game.GameLoop(diceResult, 2)
	time.Sleep(time.Second * 2)
	// Проверяем, что баланс игрока 2 уменьшился на 3
	if player2.Bank != 7 {
		t.Errorf("Expected player 2 balance to be 7, got %d", player2.Bank)
	}

	// Проверяем, что баланс игрока 1 увеличился на 3
	if player1.Bank != 8 {
		t.Errorf("Expected player 1 balance to be 8, got %d", player1.Bank)
	}

	// Проверяем, что баланс игрока 3 не изменился
	if player3.Bank != 15 {
		t.Errorf("Expected player 3 balance to remain 15, got %d", player3.Bank)
	}
}

func TestRedCardsBankOverdraftProtection(t *testing.T) {
	for i := 0; i < 20; i++ {
		// Тест на случай, если у "жертвы" карты недостаточно монет
		player1 := &model.Player{}
		player1.PlayerInit(1)
		player1.Bank = 5
		player2 := &model.Player{}
		player2.PlayerInit(2)
		player2.Bank = 1

		players := []*model.Player{player1, player2}

		redCard := model.Card{
			Name:  "RedCardTest",
			Color: "red",
			Rules: model.BuildRules{
				When:     "dice",
				Value:    2, // Активируется при значении кубика 2
				ValueAdd: 5, // Забрать "5 монет" у другого игрока
				Name:     "player_money_transfer",
			},
		}

		player1.Builds = []model.Card{redCard, redCard}
		player2.Builds = []model.Card{redCard}

		game := &Game{Players: &players, Bank: 100, queueUserBalance: make(chan BalancedUserAdd, 10)}

		game.InitQueueBalance()

		// Симулируем бросок кубика со значением 2
		diceResult := 2
		game.GameLoop(diceResult, 2)
		time.Sleep(time.Second / 100)

		//t.Log(game.Bank)
		// Проверяем, что у игрока 2 сняли 1 монету, а баланс не стал отрицательным
		if player2.Bank != 0 {
			t.Errorf("Expected player 2 balance to be 0, got %d", player2.Bank)
		}

		// Проверяем, что игроку 1 добавили только 1 монету
		if player1.Bank != 6 {
			t.Errorf("Expected player 1 balance to be 6, got %d", player1.Bank)
		}
	}
}
