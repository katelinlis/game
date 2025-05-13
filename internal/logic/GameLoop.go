package logic

import (
	"github.com/katelinlis/BackendMasters/internal/model"
	"log"
)

const (
	colorBlue  = "blue"
	colorGreen = "green"
	colorRed   = "red"
)

func (g *Game) GameLoop(dice int, user int64) (playerBalance uint8) {
	// Кэш: отслеживание карт и сохранённых данных
	CachedCard := make(map[string]model.PlayerCard)

	// Найти текущего игрока
	var currentPlayer *model.Player

	for i, player := range *g.Players {
		// Проверка на текущего игрока
		isCurrentUser := player.ID == user

		// Обработка построек игрока и манипуляции с красными картами
		portBuilt := g.hasBuiltMainItem(player, model.PORT)
		g.processPlayerBuildsCached(player, dice, isCurrentUser, CachedCard, portBuilt)

		// Проверка правила "low" для основного игрока
		if isCurrentUser {
			g.processLowBankMainBuilds(player)
			currentPlayer = player
			playerBalance = player.Bank
		}

		// Обновление игрока в общем массиве — важно для корректного состояния
		(*g.Players)[i] = player
	}

	// Применение эффектов красных карт (если такой игрок найден)
	if currentPlayer != nil {
		g.applyRedCardEffectsCached(currentPlayer, dice, CachedCard)
	}

	return playerBalance
}

// Helper function to process low-bank rules for player's main builds
func (g *Game) processLowBankMainBuilds(player *model.Player) {
	for _, card := range player.MainBuilds {
		if card.Rules.When == "low" && player.Bank == 0 {
			g.bankOperation(player, 1)
		}
	}
}

// Проверка: Построено ли здание (`mainBuild`)
func (g *Game) hasBuiltMainItem(player *model.Player, itemName string) bool {
	for _, item := range player.MainBuilds {
		if item.Name == itemName && item.Builded {
			return true
		}
	}
	return false
}

// Кэшированная обработка построек игрока
func (g *Game) processPlayerBuildsCached(player *model.Player, dice int, isCurrentUser bool, redCardMap map[string]model.PlayerCard, portBuilt bool) {
	for _, card := range player.Builds {
		g.playersColor(card, player, dice, colorBlue)

		// Ускорение: Игнорировать лишние карты без "порта"
		if card.Name == model.CONSTSushibar && !portBuilt {
			continue
		}

		// Обработка красных карт
		if card.Color == colorRed && !isCurrentUser {
			g.updateRedCardMap(redCardMap, player, card)
		}

		// Обработка текущего пользователя
		if isCurrentUser {
			g.playersColor(card, player, dice, colorGreen)
		}
	}
}

// Кэш: Вычисление количества карт у игроков
func (g *Game) calculateExternalCardCount(item model.PlayerCard, externalCount *uint8) uint8 {
	for _, player := range item.Player {
		*externalCount += player.CountOfCard
	}
	return *externalCount
}

// Распределение изменений с кэшированием
func (g *Game) distributeCachedBankChanges(item model.PlayerCard, totalCost, externalCount uint8) {
	minimalCost := totalCost / externalCount          // Сумма, которую можно заплатить каждому
	gameBankShare := int16(totalCost % externalCount) // Остаток, который раньше шел банку

	// Добавление изменений для пользователей
	for _, player := range item.Player {
		playerCut := minimalCost * player.CountOfCard // Основная часть для игрока (пропорциональная картам)

		// Если есть остаток от деления (gameBankShare), перераспределяем его
		if gameBankShare > 0 {
			playerCut += 1  // Добавляем одну монету из остатка
			gameBankShare-- // Уменьшаем общий остаток
		}

		// Обновляем баланс игрока через очередь
		select {
		case g.queueUserBalance <- BalancedUserAdd{
			userID:     player.ID,
			balanceAdd: playerCut,
		}:
		default:
			log.Printf("queueUserBalance full or unread for player %d, message dropped", len(g.queueUserBalance))
		}
	}

	// Блокировка НЕ НУЖНА для gameBank, так как банк игры больше не участвует
}
func (g *Game) playersColor(card model.Card, player *model.Player, dice int, color string) {

	if card.Color == color && card.Rules.When == "dice" {
		if (card.Rules.Value != 0 && card.Rules.Value == dice) ||
			inRange(dice, card.Rules.ValueRange) {

			if card.Rules.Name == "player_bank" {
				g.bankOperation(player, card.Rules.ValueAdd)
			}
		}

	}
}

func inRange(a int, rng [2]int) bool {
	return a >= rng[0] && a <= rng[1]
}
