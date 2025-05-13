package model

import "math/rand/v2"

type Card struct {
	Cost        uint8      `json:"cost"`
	Color       string     `json:"color"`
	Mark        string     `json:"mark"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Rules       BuildRules `json:"rules"`
}

type Cards []Card

func (c *Cards) AddCard(card Card, count int) {
	for i := 0; i < count; i++ {
		*c = append(*c, card)
	}
}
func (c *Cards) Shuffle(count int) {
	for i := 0; i < count; i++ {
		rand.Shuffle(len(*c), func(i, j int) {
			(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
		})
	}

}

type BuildRules struct {
	When       string
	Name       string
	Value      int    `json:"value"`
	ValueRange [2]int `json:"value_range"`
	ValueAdd   uint8
	Be         string
}

const PORT = "Порт"
const Terminal = "Вокзал"

func InitBuildsCards() []MainBuild {
	cards := []MainBuild{ // value - 1 build
		{
			Card{
				Cost:        0,
				Name:        "Ратуша",
				Description: "Если у вас нет монет, возьмите 1 монету из банка.",
				Rules: BuildRules{
					When:     "low",
					Name:     "player_bank",
					Value:    1,
					ValueAdd: 1,
				},
			},
			true,
		},
		{
			Card{
				Cost:        2,
				Name:        PORT,
				Description: "Если на кубиках выпало <10> или больше, можете добавить <2> к результату броска.",
				Rules: BuildRules{
					When:     "low",
					Name:     "player_bank",
					Value:    0,
					ValueAdd: 1,
				},
			},
			false,
		},
		{
			Card{
				Cost:        4,
				Name:        Terminal,
				Description: "Можете бросать 2 кубика вместо 1",
			},
			false,
		},
		{
			Card{
				Cost:        10,
				Name:        "Торговый центр",
				Description: "Если на кубиках выпало <10> или больше, можете добавить <2> к результату броска.",
			},
			false,
		},
		{
			Card{
				Cost:        22,
				Name:        "Радиовышка",
				Description: "Один раз вы можете перебросить кубики.",
			},
			false,
		},
		{
			Card{
				Cost:        30,
				Name:        "Аэропорт",
				Description: "Если вы ничего не построили в этот ход, возьмите 10 монет из банка",
			},
			false,
		},
	}

	return cards
}

func InitDeckPlayerCards() []Card {
	return []Card{
		Baker(0),
		Field(0),
	}
}

func Field(cost uint8) Card {
	return Card{
		Cost:  cost,
		Color: "blue",
		Name:  "Пшеничное поле",
		Rules: BuildRules{
			When:     "dice",
			Value:    1,
			ValueAdd: 1,
			Name:     "player_bank",
		},
	}
}

func Baker(cost uint8) Card {
	return Card{
		Mark:  CONSTFoods,
		Cost:  cost,
		Color: "green",
		Name:  "Пекарня",
		Rules: BuildRules{
			When:       "dice",
			ValueRange: [2]int{2, 3},
			ValueAdd:   1,
			Name:       "player_bank",
		},
	}
}

func Flowers() Card {
	return Card{
		Cost:  2,
		Color: "blue",
		Name:  "Цветник",
		Rules: BuildRules{
			When:     "dice",
			Value:    4,
			ValueAdd: 1,
			Name:     "player_bank",
		},
	}
}

func Shop() Card {
	return Card{
		Cost:  2,
		Mark:  CONSTFoods,
		Color: "green",
		Name:  "Магазин",
		Rules: BuildRules{
			When:     "dice",
			Value:    4,
			ValueAdd: 3,
			Name:     "player_bank",
		},
	}
}

func Forest() Card {
	return Card{
		Cost:  3,
		Color: "blue",
		Name:  "Лес",
		Rules: BuildRules{
			When:     "dice",
			Value:    5,
			ValueAdd: 1,
			Name:     "player_bank",
		},
	}
}

func Farm() Card {
	return Card{
		Cost:  1,
		Color: "blue",
		Name:  "Ферма",
		Rules: BuildRules{
			When:     "dice",
			Value:    2,
			ValueAdd: 1,
			Name:     "player_bank",
		},
	}
}

func Mine() Card {
	return Card{
		Cost:  6,
		Color: "blue",
		Name:  "Шахта",
		Rules: BuildRules{
			When:     "dice",
			Value:    9,
			ValueAdd: 6,
			Name:     "player_bank",
		},
	}
}

func AppleGarden() Card {
	return Card{
		Cost:  3,
		Color: "blue",
		Name:  "Яблоневый сад",
		Rules: BuildRules{
			When:     "dice",
			Value:    10,
			ValueAdd: 3,
			Name:     "player_bank",
		},
	}
}

const CONSTSushibar = "Суси-бар"
const CONSTSellers = "seller"
const CONSTFoods = "foods"

func Sushibar() Card {
	return Card{
		Mark:  CONSTSellers,
		Cost:  2,
		Color: "red",
		Name:  CONSTSushibar,
		Rules: BuildRules{
			When:     "dice",
			Value:    1,
			ValueAdd: 3,
			Name:     "player_money_transfer",
		},
	}
}

func Cafe() Card {
	return Card{
		Mark:  CONSTSellers,
		Cost:  2,
		Color: "red",
		Name:  "Кафе",
		Rules: BuildRules{
			When:     "dice",
			Value:    3,
			ValueAdd: 1,
			Name:     "player_money_transfer",
		},
	}
}

func Restourant() Card {
	return Card{
		Mark:  CONSTSellers,
		Cost:  3,
		Color: "red",
		Name:  "Ресторан",
		Rules: BuildRules{
			When:       "dice",
			ValueRange: [2]int{9, 10},
			ValueAdd:   2,
			Name:       "player_money_transfer",
		},
	}
}
func Fastfood() Card {
	return Card{
		Mark:  CONSTSellers,
		Cost:  1,
		Color: "red",
		Name:  "Закусочная",
		Rules: BuildRules{
			When:     "dice",
			Value:    8,
			ValueAdd: 1,
			Name:     "player_money_transfer",
		},
	}
}

func Pizza() Card {
	return Card{
		Mark:  CONSTSellers,
		Cost:  1,
		Color: "red",
		Name:  "Пиццерия",
		Rules: BuildRules{
			When:     "dice",
			Value:    7,
			ValueAdd: 1,
			Name:     "player_money_transfer",
		},
	}
}

func Flour_SHOP() Card {
	return Card{
		Cost:        1,
		Color:       "green",
		Name:        "Лес",
		Description: "Возьмите 1 монету за каждый ваш цветник",
		Rules: BuildRules{
			When:     "dice",
			Value:    6,
			ValueAdd: 1,
			Name:     "player_bank",
		},
	}
}
