package logic

import (
	"log"
	"sync"
	"time"

	"github.com/katelinlis/BackendMasters/internal/model"
)

type GameSession struct {
	Game    *Game
	mutex   sync.RWMutex
	LastUse time.Time
}

type Game struct {
	Players          *[]*model.Player // 4 Players
	DeckVisible      *[]model.Card    // 6 cards
	Deck             *model.Cards
	gameBank         sync.Mutex
	Bank             int16
	Turn             int64
	TurnStatus       bool
	queueUserBalance chan BalancedUserAdd
	QueueNotiff      chan model.Nottif
	Lobby            *model.Lobby
}

type BalancedUserAdd struct {
	userID     int64
	balanceAdd uint8
}

func GameInit(lobby *model.Lobby) *Game {
	game := &Game{}
	game.queueUserBalance = make(chan BalancedUserAdd)
	game.QueueNotiff = make(chan model.Nottif)
	game.Deck = initDeck()

	game.Bank = int16(205)

	game.Players = &[]*model.Player{}

	for _, player := range lobby.PlayerList {
		p := &model.Player{}
		p.PlayerInit(player.ID)

		*game.Players = append(*game.Players, p)
	}

	game.Lobby = lobby
	game.InitQueueBalance()

	return game
}

func (game *Game) SearchPlayer(uuid int64) (player *model.Player) {
	if game.Players != nil {
		for i := 0; i < len(*game.Players); i++ {
			if (*game.Players)[i].ID == uuid {
				player = (*game.Players)[i]
			}
		}
	}

	return player
}

func (game *Game) sendForAllPlayers(notif model.Nottif) {

	game.Lobby.BroadcastSendMessage(notif)

}

func (game *Game) sendForPersonal(notif model.Nottif, user int64) {
	player := game.Lobby.PlayerList.GetPlayer(user)
	if player == nil {
		return
	}
	select {
	case player.QNotiff <- notif:
	default:
		log.Printf("QNotiff full or unread for player %d, message dropped", user)
	}
}

func (g *Game) bankOperation(player *model.Player, value uint8) bool {
	defer g.gameBank.Unlock()

	g.gameBank.Lock()
	if g.Bank >= int16(value) {
		player.Lock()
		defer player.Unlock()

		g.Bank -= int16(value)
		player.Bank += value

		return true
	}

	return false
}

func (g *Game) FillDeckVisible() {

	length := 0
	if g.DeckVisible != nil {
		length = len(*g.DeckVisible)
	}

	remaining := len(*g.Deck)
	toTake := 6 - length
	if toTake > remaining {
		toTake = remaining // Берём только доступное количество
	}

	deck, newDeck := (*g.Deck)[:toTake], (*g.Deck)[toTake:]
	*g.Deck = newDeck

	if g.DeckVisible == nil {
		g.DeckVisible = &[]model.Card{}
	}
	*g.DeckVisible = append(*g.DeckVisible, deck...)

}

func (g *Game) Next(currentUser int64) {

	if g.Players == nil || len(*g.Players) == 0 {
		return
	}

	g.TurnStatus = false

	// Найти индекс текущего игрока
	currentIndex := -1
	for i, player := range *g.Players {
		if player.ID == currentUser {
			currentIndex = i
			break
		}
	}

	if currentIndex == -1 {
		return
	}

	// Определить следующего игрока по циклической очереди

	nextIndex := (currentIndex + 1) % len(*g.Players)
	g.Turn = (*g.Players)[nextIndex].ID

	playerLobby := g.Lobby.PlayerList.GetPlayer(g.Turn)

	g.sendForAllPlayers(model.Nottif{Name: playerLobby.Name, Message: "Turn", Who: playerLobby.ID})

	for _, player := range *g.Players {
		g.sendForPersonal(model.Nottif{
			Message: "newBalance",
			Number:  int64(player.Bank),
		}, player.ID)
	}

}

func (g *Game) InitQueueBalance() {
	go func() { // Горутина для работы с queueUserBalance
		for val := range g.queueUserBalance { // Чтение из канала
			player := g.SearchPlayer(val.userID) // Поиск игрока по userID
			if player != nil {                   // Если игрок найден, обновляем его банк
				player.Lock()
				player.Bank += val.balanceAdd
				player.Unlock()
			}
		}
	}()
}
