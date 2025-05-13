package logic

import (
	"github.com/katelinlis/BackendMasters/internal/model"
)

// Оптимизированное применение эффектов красных карт
func (g *Game) applyRedCardEffectsCached(currentPlayer *model.Player, dice int, redCardMap map[string]model.PlayerCard) {
	for _, item := range redCardMap {

		// Предварительная проверка
		if item.Rules.When != "dice" || item.Rules.Value != dice {
			continue
		}

		// Использовать кэш подсчёта карт (externalCount) вместо пересчёта каждый раз
		var externalCount uint8
		totalCost := item.Rules.ValueAdd * g.calculateExternalCardCount(item, &externalCount)

		currentPlayer.Lock()
		if currentPlayer.Bank < totalCost {
			totalCost = currentPlayer.Bank
		}
		currentPlayer.Bank -= totalCost
		currentPlayer.Unlock()

		// Распределение изменений банка
		g.distributeCachedBankChanges(item, totalCost, externalCount)
	}
}

// Extracted helper: Handle red card map updates
func (g *Game) updateRedCardMap(redCardMap map[string]model.PlayerCard, player *model.Player, card model.Card) {
	if existingCard, exists := redCardMap[card.Name]; exists {
		if playerData, found := existingCard.Player[player.ID]; found {

			playerData.CountOfCard++

		} else {
			existingCard.Player[player.ID] = model.PlayerOfCard{Player: player, CountOfCard: 1}
		}
	} else {
		newPlayerCard := model.PlayerCard{Card: card, Player: map[int64]model.PlayerOfCard{
			player.ID: {Player: player, CountOfCard: 1},
		}}
		redCardMap[card.Name] = newPlayerCard
	}
}
