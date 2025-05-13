package logic

import (
	"github.com/katelinlis/BackendMasters/internal/model"
)

func initDeck() *model.Cards {
	Deck := &model.Cards{}
	Deck.AddCard(model.Flowers(), 5)
	Deck.AddCard(model.Field(1), 6)
	Deck.AddCard(model.Baker(1), 6)
	Deck.AddCard(model.Shop(), 5)
	Deck.AddCard(model.Forest(), 6)
	Deck.AddCard(model.Farm(), 5)
	Deck.AddCard(model.Sushibar(), 5)
	Deck.AddCard(model.Cafe(), 5)
	Deck.AddCard(model.Fastfood(), 5)
	Deck.AddCard(model.Restourant(), 5)
	Deck.AddCard(model.Pizza(), 5)
	Deck.AddCard(model.Mine(), 6)
	Deck.AddCard(model.AppleGarden(), 6)
	//Deck.AddCard(model.Flour_SHOP(), 4) -- TODO need add func to calc of Цветник

	Deck.Shuffle(10)
	return Deck
}
